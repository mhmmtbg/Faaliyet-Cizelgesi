# Geliştirme ve Derleme

## Mimari

Uygulamanın tamamı tek bir dosyadadır: **`app/faaliyet-cizelgesi.html`**. CDN veya internet bağlantısı kullanmaz. Saf HTML, CSS, JavaScript ve SVG ile yazılmıştır. Excel desteği için [SheetJS](https://sheetjs.com) (Apache-2.0) dosyanın içine gömülüdür.

Dosya içindeki ana bölümler (`/* ===== ... ===== */` başlıklarıyla ayrılmıştır):

| Bölüm | İçerik |
|---|---|
| durum / yardımcılar | Uygulama durumu, tarih biçimleri, renk atama |
| kalıcı kayıt | Tarayıcı deposu, HTML'e gömülü veri, dosyaya otomatik yazma |
| dik açılı ok rotalama | Kutuları engel sayan seyrek ızgara üzerinde, dönüş cezalı en kısa yol (Dijkstra + öncelik kuyruğu) |
| öncül / ardıl bağımlılıkları | En erken başlangıç, topolojik sıra, zincirleme tarih planı, döngü kontrolü |
| çizim | Zaman çizelgesi, lejant, liste, adam*saat, akış şeması (SVG) |
| sürükleyerek bağ kurma | Çizelge ve akış şeması için ortak sürükle-bırak |
| Excel | İçe aktarma (başlık/sütun tahmini), dışa aktarma, şablon |
| PNG | Çizelge (canvas) ve akış şeması (SVG → canvas) görselleri |
| masaüstü (exe) sürümü | `window.masaustu` köprüsü varsa dosya tabanlı çalışma |

### Proje dosyası (`.fzc`)

```json
{
  "v": 2,
  "kayit": "2026-10-02T09:00:00Z",
  "ayar": { "groupBy": "sorumlu", "colorBy": "sorumlu", "view": "auto", "ppd": 22, "showLinks": true, "tab": "liste" },
  "faaliyetler": [
    { "id": "o3", "no": 3, "ad": "Titreşim testi – X ekseni",
      "baslangic": "2026-10-12", "bitis": "2026-10-14",
      "sorumlu": "A. Yılmaz", "proje": "Sistem A", "saat": 32,
      "aciklama": "MIL-STD-810H Method 514.8", "kapali": false,
      "oncul": [ { "id": "o2", "gecikme": 0 } ] }
  ]
}
```

- `oncul` bağları kalıcı `id` ile kurulur; `no` kullanıcıya gösterilen sıra numarasıdır.
- Tarayıcı sürümünün "Dosyaya kaydet" çıktısında aynı JSON, HTML içindeki `<script id="gomuluVeri" type="application/json">` etiketinde durur.

## Masaüstü (exe) sürümü

`desktop/` klasörü, aynı HTML'i bir Windows penceresinde gösteren küçük bir Go programıdır:

- **[go-webview2](https://github.com/jchv/go-webview2):** Saf Go, CGO gerektirmez. Windows'taki Edge WebView2 motorunu kullanır. Exe ~7 MB'tır.
- HTML, `go:embed` ile exe'nin içine gömülür ve yalnızca `127.0.0.1` üzerinden, rastgele bir jetonlu adreste sunulur.
- `desktop/shim.js` sayfaya `window.masaustu` arayüzünü sağlar. Bu arayüz Go tarafına bağlanan `native*` fonksiyonlarını çağırır:
  - Windows **Aç / Farklı kaydet** pencereleri
  - Dosya okuma ve yazma (yarım yazmaya karşı geçici dosya + yeniden adlandırma)
  - Kurtarma kaydı ve son açılanlar (`%LOCALAPPDATA%\FaaliyetCizelgesi`)
  - Pencere başlığı ve kaydedilmemiş değişiklik bayrağı
- Gerçek Win32 menü çubuğu kullanılır. Pencere yordamı sarılır; menü komutları ve "kaydedilmemiş değişiklikle kapatma" buradan yönetilir.
- Excel/PNG indirmeleri (`<a download>`) yakalanır ve Windows kaydetme penceresine yönlendirilir.
- `alert` / `confirm`, yerel sunucuya senkron istekle uygulama adlı Windows mesaj kutusu olarak gösterilir.
- Tarih kutularının `gg.aa.yyyy` görünmesi için WebView2 `--lang=tr` ile başlatılır.
- HTML `window.masaustu`'yu bulamazsa normal tarayıcı moduna döner. Yani `app/faaliyet-cizelgesi.html` iki sürüm için de tek kaynaktır.

### Derleme

Gereksinim: [Go 1.22+](https://go.dev/dl/) ve ilk derlemede bağımlılıkları indirmek için internet.

**Windows:**
```bat
desktop\build.bat
```

**Linux / macOS (çapraz derleme):**
```sh
./desktop/build.sh
```

Çıktı: `dist/FaaliyetCizelgesi.exe`. Betik, `app/faaliyet-cizelgesi.html` dosyasını derlemeden önce `desktop/app.html` olarak kopyalar.

`desktop/rsrc_windows_amd64.syso` derlemeye otomatik dahil olur. İçinde şunlar vardır:
- uygulama ikonu (`assets/icon.png`'den üretilmiştir),
- sürüm bilgisi,
- yönetici izni istemeyen (`asInvoker`) ve DPI farkındalıklı bildirim (manifest).

## Yayınlama (GitHub Releases)

Exe dosyası depoya eklenmez (`.gitignore`), **Releases** üzerinden dağıtılır:

1. `desktop/build.bat` ile exe'yi derleyin.
2. GitHub'da depo sayfası → **Releases** → **Draft a new release**.
3. Etiket olarak sürüm numarası verin (ör. `v1.1.0`).
4. `dist/FaaliyetCizelgesi.exe` ve `app/faaliyet-cizelgesi.html` dosyalarını sürükleyip **Publish release** deyin.
