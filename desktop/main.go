//go:build windows

// Faaliyet Zaman Çizelgesi — kurulumsuz, tek dosyalık Windows uygulaması.
// Arayüz Windows 10/11'de hazır gelen Edge WebView2 motoruyla gösterilir;
// tarayıcı açılmaz, internet gerekmez, yönetici yetkisi gerekmez.
package main

import (
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"unicode/utf16"
	"unsafe"

	webview2 "github.com/jchv/go-webview2"
	"golang.org/x/sys/windows"
)

//go:embed app.html
var appHTML []byte

//go:embed shim.js
var shimJS string

const uygAdi = "Faaliyet Zaman Çizelgesi"
const projeFiltre = "Faaliyet çizelgesi (*.fzc)|*.fzc|Eski kayıtlar (*.json;*.html)|*.json;*.html;*.htm|Tüm dosyalar (*.*)|*.*"

var (
	comdlg32          = windows.NewLazySystemDLL("comdlg32.dll")
	pGetSaveFileName  = comdlg32.NewProc("GetSaveFileNameW")
	pGetOpenFileName  = comdlg32.NewProc("GetOpenFileNameW")
	user32            = windows.NewLazySystemDLL("user32.dll")
	pMessageBox       = user32.NewProc("MessageBoxW")
	pCreateMenu       = user32.NewProc("CreateMenu")
	pCreatePopupMenu  = user32.NewProc("CreatePopupMenu")
	pAppendMenu       = user32.NewProc("AppendMenuW")
	pSetMenu          = user32.NewProc("SetMenu")
	pGetMenu          = user32.NewProc("GetMenu")
	pDestroyMenu      = user32.NewProc("DestroyMenu")
	pDrawMenuBar      = user32.NewProc("DrawMenuBar")
	pSetWindowLongPtr = user32.NewProc("SetWindowLongPtrW")
	pCallWindowProc   = user32.NewProc("CallWindowProcW")
	pPostMessage      = user32.NewProc("PostMessageW")
	pShowWindow       = user32.NewProc("ShowWindow")
	pLoadImage        = user32.NewProc("LoadImageW")
	pSendMessage      = user32.NewProc("SendMessageW")
	kernel32          = windows.NewLazySystemDLL("kernel32.dll")
	pGetModuleHandle  = kernel32.NewProc("GetModuleHandleW")
)

const (
	wmClose       = 0x0010
	wmCommand     = 0x0111
	wmSetIcon     = 0x0080
	mfString      = 0x0
	mfPopup       = 0x10
	mfSeparator   = 0x800
	mfGrayed      = 0x1
	gwlpWndProc   = ^uintptr(3) // -4
	swMaximize    = 3
	mbOk          = 0x0
	mbOkCancel    = 0x1
	mbYesNoCancel = 0x3
	mbYesNo       = 0x4
	mbIconError   = 0x10
	mbIconQuest   = 0x20
	mbIconWarn    = 0x30
	mbIconInfo    = 0x40
	mbSetFg       = 0x10000
	idOk          = 1
	idCancel      = 2
	idYes         = 6
	idNo          = 7
)

/* ---------------- dosya pencereleri ---------------- */

type openFileName struct {
	lStructSize       uint32
	hwndOwner         uintptr
	hInstance         uintptr
	lpstrFilter       *uint16
	lpstrCustomFilter *uint16
	nMaxCustFilter    uint32
	nFilterIndex      uint32
	lpstrFile         *uint16
	nMaxFile          uint32
	lpstrFileTitle    *uint16
	nMaxFileTitle     uint32
	lpstrInitialDir   *uint16
	lpstrTitle        *uint16
	Flags             uint32
	nFileOffset       uint16
	nFileExtension    uint16
	lpstrDefExt       *uint16
	lCustData         uintptr
	lpfnHook          uintptr
	lpTemplateName    *uint16
	pvReserved        uintptr
	dwReserved        uint32
	FlagsEx           uint32
}

const (
	ofnOverwritePrompt = 0x2
	ofnNoChangeDir     = 0x8
	ofnPathMustExist   = 0x800
	ofnFileMustExist   = 0x1000
	ofnExplorer        = 0x80000
)

func filterUTF16(f string) *uint16 {
	if f == "" {
		f = "Tüm dosyalar (*.*)|*.*"
	}
	u := utf16.Encode([]rune(strings.ReplaceAll(f, "|", "\x00") + "\x00\x00"))
	return &u[0]
}

func strPtr(s string) *uint16 {
	if s == "" {
		return nil
	}
	p, _ := windows.UTF16PtrFromString(s)
	return p
}

func defExtOf(filter string) string {
	parts := strings.Split(filter, "|")
	if len(parts) >= 2 {
		e := strings.TrimPrefix(strings.Split(parts[1], ";")[0], "*.")
		if e != "*" {
			return e
		}
	}
	return ""
}

func fileDialog(owner uintptr, save bool, title, defName, filter string) string {
	buf := make([]uint16, 8192)
	dir := sonKlasor()
	if defName != "" {
		if filepath.IsAbs(defName) {
			dir = filepath.Dir(defName)
			defName = filepath.Base(defName)
		}
		copy(buf, utf16.Encode([]rune(defName)))
	}
	ofn := openFileName{
		hwndOwner:       owner,
		lpstrFilter:     filterUTF16(filter),
		nFilterIndex:    1,
		lpstrFile:       &buf[0],
		nMaxFile:        uint32(len(buf)),
		lpstrInitialDir: strPtr(dir),
		lpstrTitle:      strPtr(title),
		lpstrDefExt:     strPtr(defExtOf(filter)),
		Flags:           ofnExplorer | ofnNoChangeDir | ofnPathMustExist,
	}
	ofn.lStructSize = uint32(unsafe.Sizeof(ofn))
	var r uintptr
	if save {
		ofn.Flags |= ofnOverwritePrompt
		r, _, _ = pGetSaveFileName.Call(uintptr(unsafe.Pointer(&ofn)))
	} else {
		ofn.Flags |= ofnFileMustExist
		r, _, _ = pGetOpenFileName.Call(uintptr(unsafe.Pointer(&ofn)))
	}
	runtime.KeepAlive(buf)
	if r == 0 {
		return ""
	}
	p := windows.UTF16ToString(buf)
	st := loadState()
	st.LastDir = filepath.Dir(p)
	saveState(st)
	return p
}

func msgBox(owner uintptr, text, caption string, flags uintptr) int {
	if caption == "" {
		caption = uygAdi
	}
	r, _, _ := pMessageBox.Call(owner, uintptr(unsafe.Pointer(strPtr(text))), uintptr(unsafe.Pointer(strPtr(caption))), flags|mbSetFg)
	return int(r)
}

/* ---------------- kalıcı durum (%LOCALAPPDATA%\FaaliyetCizelgesi) ---------------- */

type appState struct {
	LastDir string   `json:"lastDir"`
	Recent  []string `json:"recent"`
}

var stateMu sync.Mutex

func appDir() string {
	base := os.Getenv("LOCALAPPDATA")
	if base == "" {
		base = os.TempDir()
	}
	d := filepath.Join(base, "FaaliyetCizelgesi")
	os.MkdirAll(d, 0o755)
	return d
}

func loadState() appState {
	stateMu.Lock()
	defer stateMu.Unlock()
	var s appState
	if b, err := os.ReadFile(filepath.Join(appDir(), "ayarlar.json")); err == nil {
		json.Unmarshal(b, &s)
	}
	return s
}

func saveState(s appState) {
	stateMu.Lock()
	defer stateMu.Unlock()
	b, _ := json.MarshalIndent(s, "", " ")
	writeAtomic(filepath.Join(appDir(), "ayarlar.json"), b)
}

func sonKlasor() string {
	if d := loadState().LastDir; d != "" {
		if fi, err := os.Stat(d); err == nil && fi.IsDir() {
			return d
		}
	}
	if h, err := os.UserHomeDir(); err == nil {
		return filepath.Join(h, "Documents")
	}
	return ""
}

func sonAcilanEkle(p string) {
	st := loadState()
	l := []string{p}
	for _, x := range st.Recent {
		if !strings.EqualFold(x, p) {
			l = append(l, x)
		}
	}
	if len(l) > 8 {
		l = l[:8]
	}
	st.Recent = l
	st.LastDir = filepath.Dir(p)
	saveState(st)
}

func writeAtomic(path string, data []byte) error {
	tmp := path + ".tmp~"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return os.WriteFile(path, data, 0o644)
	}
	return nil
}

/* ---------------- menü çubuğu ---------------- */

var komutlar = map[int]string{}
var hwndAna uintptr
var wv webview2.WebView

func menuEkle(m uintptr, id int, metin, komut string) {
	komutlar[id] = komut
	pAppendMenu.Call(m, mfString, uintptr(id), uintptr(unsafe.Pointer(strPtr(metin))))
}
func ayrac(m uintptr) { pAppendMenu.Call(m, mfSeparator, 0, 0) }
func altMenu(ana uintptr, metin string) uintptr {
	m, _, _ := pCreatePopupMenu.Call()
	pAppendMenu.Call(ana, mfPopup, m, uintptr(unsafe.Pointer(strPtr(metin))))
	return m
}

func menuKur() {
	komutlar = map[int]string{}
	bar, _, _ := pCreateMenu.Call()

	d := altMenu(bar, "&Dosya")
	menuEkle(d, 101, "&Yeni\tCtrl+N", "yeni")
	menuEkle(d, 102, "&Aç…\tCtrl+O", "ac")
	son := altMenu(d, "&Son açılanlar")
	var mevcut []string
	for _, p := range loadState().Recent {
		if _, err := os.Stat(p); err == nil {
			mevcut = append(mevcut, p)
		}
	}
	if len(mevcut) == 0 {
		pAppendMenu.Call(son, mfString|mfGrayed, 0, uintptr(unsafe.Pointer(strPtr("(boş)"))))
	} else {
		for i, p := range mevcut {
			menuEkle(son, 900+i, "&"+string(rune('1'+i))+"  "+filepath.Base(p)+"   —   "+filepath.Dir(p), "ac:"+p)
		}
		ayrac(son)
		menuEkle(son, 950, "Listeyi temizle", "__sonTemizle")
	}
	ayrac(d)
	menuEkle(d, 103, "&Kaydet\tCtrl+S", "kaydet")
	menuEkle(d, 104, "&Farklı kaydet…\tCtrl+Shift+S", "farkli")
	ayrac(d)
	menuEkle(d, 105, "&Excel'den aktar…\tCtrl+I", "excelAktar")
	da := altMenu(d, "&Dışa aktar")
	menuEkle(da, 110, "&Excel çalışma kitabı (.xlsx)…\tCtrl+E", "excelKaydet")
	menuEkle(da, 111, "Çizelge &görseli (.png)…", "pngCizelge")
	menuEkle(da, 112, "&Akış şeması görseli (.png)…", "pngAkis")
	ayrac(da)
	menuEkle(da, 113, "Paylaşılabilir &HTML (verilerle birlikte)…", "html")
	menuEkle(d, 106, "Boş Excel &şablonu…", "sablon")
	ayrac(d)
	menuEkle(d, 107, "Ya&zdır…\tCtrl+P", "yazdir")
	menuEkle(d, 108, "Dosya &konumunu aç", "konum")
	ayrac(d)
	menuEkle(d, 109, "Çı&kış\tAlt+F4", "__cikis")

	dz := altMenu(bar, "Dü&zen")
	menuEkle(dz, 201, "Bağımlılık ihlallerini &hizala", "hizala")

	g := altMenu(bar, "&Görünüm")
	menuEkle(g, 301, "&Faaliyet listesi\tCtrl+1", "sekme:liste")
	menuEkle(g, 302, "&Adam*saat\tCtrl+2", "sekme:efor")
	menuEkle(g, 303, "Akış ş&eması\tCtrl+3", "sekme:akis")

	y := altMenu(bar, "&Yardım")
	menuEkle(y, 401, "&Veri klasörünü aç", "__veriKlasoru")
	ayrac(y)
	menuEkle(y, 402, "&Hakkında", "__hakkinda")

	eski, _, _ := pGetMenu.Call(hwndAna)
	pSetMenu.Call(hwndAna, bar)
	if eski != 0 {
		pDestroyMenu.Call(eski)
	}
	pDrawMenuBar.Call(hwndAna)
}

func jsKomut(k string) {
	b, _ := json.Marshal(k)
	wv.Eval("window.__menu && window.__menu(" + string(b) + ")")
}

/* ---------------- pencere yordamı: menü komutları ve kapatma ---------------- */

var eskiProc uintptr
var kirli, kapanabilir bool

func yeniProc(hwnd, msg, wp, lp uintptr) uintptr {
	switch msg {
	case wmCommand:
		if wp>>16 == 0 { // menü
			if k, ok := komutlar[int(wp&0xffff)]; ok {
				switch k {
				case "__cikis":
					pPostMessage.Call(hwnd, wmClose, 0, 0)
				case "__sonTemizle":
					st := loadState()
					st.Recent = nil
					saveState(st)
					menuKur()
				case "__veriKlasoru":
					exec.Command("explorer", appDir()).Start()
				case "__hakkinda":
					msgBox(hwnd, uygAdi+"\nSürüm 1.1\n\nKurulum ve yönetici yetkisi gerektirmez, internet kullanmaz.\n"+
						"Projeler .fzc dosyalarına kaydedilir.\n\nAyarlar ve kurtarma kaydı:\n"+appDir(), "Hakkında", mbIconInfo)
				default:
					jsKomut(k)
				}
				return 0
			}
		}
	case wmClose:
		if kirli && !kapanabilir {
			jsKomut("kapatma-istegi")
			return 0
		}
	}
	r, _, _ := pCallWindowProc.Call(eskiProc, hwnd, msg, wp, lp)
	return r
}

/* ---------------- arayüzle iletişim ---------------- */

type dosyaSonucu struct {
	Yol    string `json:"yol,omitempty"`
	Ad     string `json:"ad,omitempty"`
	Icerik string `json:"icerik,omitempty"`
	B64    string `json:"b64,omitempty"`
	Hata   string `json:"hata,omitempty"`
}

func oku(p string, ikili bool) dosyaSonucu {
	b, err := os.ReadFile(p)
	if err != nil {
		return dosyaSonucu{Hata: err.Error()}
	}
	r := dosyaSonucu{Yol: p, Ad: filepath.Base(p)}
	if ikili {
		r.B64 = base64.StdEncoding.EncodeToString(b)
	} else {
		r.Icerik = string(b)
	}
	return r
}

func geriCagir(id int, v interface{}) {
	js, _ := json.Marshal(v)
	b, _ := json.Marshal(id)
	wv.Eval("window.__ncb(" + string(b) + "," + string(js) + ")")
}

func baglantilar(w webview2.WebView) {
	w.Bind("nativeBaslangic", func() interface{} {
		for _, a := range os.Args[1:] {
			if strings.HasPrefix(a, "-") {
				continue
			}
			p, _ := filepath.Abs(a)
			low := strings.ToLower(p)
			if strings.HasSuffix(low, ".fzc") || strings.HasSuffix(low, ".json") || strings.HasSuffix(low, ".html") || strings.HasSuffix(low, ".htm") {
				if r := oku(p, false); r.Hata == "" {
					sonAcilanEkle(p)
					w.Dispatch(menuKur)
					return r
				}
			}
		}
		return nil
	})
	/* Dosya ve mesaj pencereleri olay işleyicisi içinde değil, ana döngüde açılır; sonuç window.__ncb ile döner */
	w.Bind("nativeDiyalog", func(id int, tur, baslik, ad, filtre string) {
		w.Dispatch(func() {
			var res interface{}
			switch tur {
			case "ac", "acIkili":
				p := fileDialog(hwndAna, false, baslik, "", filtre)
				if p == "" {
					res = nil
				} else {
					r := oku(p, tur == "acIkili")
					if r.Hata != "" {
						msgBox(hwndAna, "Dosya açılamadı:\n"+p+"\n\n"+r.Hata, "", mbIconError)
						res = nil
					} else {
						if tur == "ac" {
							sonAcilanEkle(p)
							menuKur()
						}
						res = r
					}
				}
			case "kaydetYeri":
				p := fileDialog(hwndAna, true, baslik, ad, filtre)
				if p == "" {
					res = nil
				} else {
					res = p
				}
			}
			geriCagir(id, res)
		})
	})
	w.Bind("nativeSoru", func(id int, mesaj, detay string, dugme int, tur string) {
		w.Dispatch(func() {
			ikon := uintptr(mbIconQuest)
			switch tur {
			case "warning":
				ikon = mbIconWarn
			case "error":
				ikon = mbIconError
			case "info":
				ikon = mbIconInfo
			}
			metin := mesaj
			if detay != "" {
				metin += "\n\n" + detay
			}
			var sonuc int
			switch dugme {
			case 3: // Kaydet / Kaydetme / İptal  →  Evet / Hayır / İptal
				r := msgBox(hwndAna, metin+"\n\nEvet: kaydet   ·   Hayır: kaydetmeden devam et", "", mbYesNoCancel|ikon)
				sonuc = map[int]int{idYes: 0, idNo: 1, idCancel: 2}[r]
			case 2:
				r := msgBox(hwndAna, metin, "", mbOkCancel|ikon)
				if r == idOk {
					sonuc = 0
				} else {
					sonuc = 1
				}
			default:
				msgBox(hwndAna, metin, "", mbOk|ikon)
				sonuc = 0
			}
			geriCagir(id, sonuc)
		})
	})
	w.Bind("nativeOku", func(p string) interface{} {
		r := oku(p, false)
		if r.Hata != "" {
			w.Dispatch(func() { msgBox(hwndAna, "Dosya açılamadı:\n"+p+"\n\n"+r.Hata, "", mbIconError) })
			return nil
		}
		sonAcilanEkle(p)
		w.Dispatch(menuKur)
		return r
	})
	w.Bind("nativeYaz", func(p, veri string, ikili, proje bool) map[string]interface{} {
		var b []byte
		var err error
		if ikili {
			b, err = base64.StdEncoding.DecodeString(veri)
		} else {
			b = []byte(veri)
		}
		if err == nil {
			err = writeAtomic(p, b)
		}
		if err != nil {
			hata := err.Error()
			w.Dispatch(func() {
				msgBox(hwndAna, "Kaydedilemedi:\n"+p+"\n\n"+hata+
					"\n\nKlasöre yazma izniniz olduğundan ve dosyanın başka bir programda açık olmadığından emin olun.", "", mbIconError)
			})
			return map[string]interface{}{"ok": false}
		}
		if proje {
			sonAcilanEkle(p)
			w.Dispatch(menuKur)
		} else {
			st := loadState()
			st.LastDir = filepath.Dir(p)
			saveState(st)
		}
		return map[string]interface{}{"ok": true, "ad": filepath.Base(p)}
	})
	w.Bind("nativeKurtarmaYaz", func(veri string) {
		writeAtomic(filepath.Join(appDir(), "kurtarma.json"), []byte(veri))
	})
	w.Bind("nativeKurtarmaOku", func() string {
		b, err := os.ReadFile(filepath.Join(appDir(), "kurtarma.json"))
		if err != nil {
			return ""
		}
		return string(b)
	})
	w.Bind("nativeKurtarmaSil", func() { os.Remove(filepath.Join(appDir(), "kurtarma.json")) })
	w.Bind("nativeKirli", func(d bool, baslik string) {
		kirli = d
		w.Dispatch(func() { w.SetTitle(baslik) })
	})
	w.Bind("nativeKapat", func() {
		kapanabilir = true
		pPostMessage.Call(hwndAna, wmClose, 0, 0)
	})
	w.Bind("nativeKonum", func(p string) {
		if _, err := os.Stat(p); err == nil {
			exec.Command("explorer", "/select,", p).Start()
		}
	})
}

func main() {
	runtime.LockOSThread()

	/* tarih kutuları gg.aa.yyyy ve Türkçe arayüz metinleri */
	os.Setenv("WEBVIEW2_ADDITIONAL_BROWSER_ARGUMENTS", "--lang=tr --disable-features=msSmartScreenProtection")

	/* Uygulama yalnızca bu bilgisayardan erişilebilen yerel bir adresten sunulur */
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		msgBox(0, "Yerel sunucu başlatılamadı: "+err.Error(), "", mbIconError)
		return
	}
	tb := make([]byte, 12)
	rand.Read(tb)
	jeton := hex.EncodeToString(tb)
	mux := http.NewServeMux()
	mux.HandleFunc("/"+jeton+"/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		w.Write(appHTML)
	})
	/* alert / confirm: senkron istek, uygulama adlı Windows mesaj kutusu */
	mux.HandleFunc("/"+jeton+"/__mesaj", func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(io.LimitReader(r.Body, 1<<20))
		if r.URL.Query().Get("k") == "confirm" {
			if msgBox(hwndAna, string(b), "", mbOkCancel|mbIconQuest) == idOk {
				w.Write([]byte("1"))
			} else {
				w.Write([]byte("0"))
			}
			return
		}
		msgBox(hwndAna, string(b), "", mbOk|mbIconInfo)
		w.Write([]byte("1"))
	})
	go http.Serve(ln, mux)

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     false,
		AutoFocus: true,
		DataPath:  filepath.Join(appDir(), "WebView2"),
		WindowOptions: webview2.WindowOptions{
			Title:  uygAdi,
			Width:  1440,
			Height: 900,
			Center: true,
			IconId: 1,
		},
	})
	if w == nil {
		msgBox(0, "Microsoft Edge WebView2 bileşeni bu bilgisayarda bulunamadı.\n\n"+
			"Windows 10 ve 11'de normalde hazır gelir. Bulunamıyorsa uygulamanın HTML sürümünü "+
			"(faaliyet-cizelgesi.html) Edge veya Chrome ile açarak aynı şekilde kullanabilirsiniz.", "", mbIconError)
		return
	}
	defer w.Destroy()
	wv = w
	hwndAna = uintptr(w.Window())
	w.SetSize(1000, 640, webview2.HintMin)

	/* görev çubuğu ve başlık ikonu */
	if hInst, _, _ := pGetModuleHandle.Call(0); hInst != 0 {
		if hb, _, _ := pLoadImage.Call(hInst, 1, 1, 32, 32, 0); hb != 0 {
			pSendMessage.Call(hwndAna, wmSetIcon, 1, hb)
		}
		if hk, _, _ := pLoadImage.Call(hInst, 1, 1, 16, 16, 0); hk != 0 {
			pSendMessage.Call(hwndAna, wmSetIcon, 0, hk)
		}
	}

	/* pencere yordamını sar: menü komutları + kaydedilmemiş değişiklikte kapatma uyarısı */
	eskiProc, _, _ = pSetWindowLongPtr.Call(hwndAna, gwlpWndProc, syscall.NewCallback(yeniProc))
	menuKur()
	pShowWindow.Call(hwndAna, swMaximize)

	baglantilar(w)
	w.Init(shimJS)
	w.Navigate("http://" + ln.Addr().String() + "/" + jeton + "/")
	w.Run()
}
