# Geliştirme ve Derleme

## Mimari

Uygulamanın tamamı tek bir dosyadadır: **`app/faaliyet-cizelgesi.html`**. CDN veya internet bağlantısı kullanmaz. Saf HTML, CSS, JavaScript ve SVG ile yazılmıştır. Excel desteği için [SheetJS](https://sheetjs.com) (Apache-2.0) dosyanın içine gömülüdür.

Dosya içindeki ana bölümler (`/* ===== ... ===== */` başlıklarıyla ayrılmıştır):

| Bölüm | İçerik |
|---|---|
| durum / yardımcılar | Uygulama durumu, tarih biçimleri, renk atama |
| çalışma takvimi | Hafta sonu, Türkiye resmi tatilleri (`TR_SABIT`, `TR_DINI`), özel tatiller; iş günü sayma/ekleme, `yerlestir` (süreyi koruyarak taşıma), `bagSonrasi` (bağdan sonraki en erken başlangıç) |
| faaliyet modeli | Grup hesabı (`gruplariGuncelle`), iş kırılımı sırası (`wbsSirasi`) |
| kalıcı kayıt | v3 biçimi, tarayıcı deposu, HTML'e gömülü veri, dosyaya otomatik yazma |
| geri al / yinele | Her çizimde faaliyet + proje özeti; değiştiyse önceki hal yığına (100 adım) |
| dik açılı ok rotalama | Kutuları engel sayan seyrek ızgara üzerinde, dönüş cezalı en kısa yol (Dijkstra + öncelik kuyruğu). "Sert" hatlar (akış şemasındaki aşama çizgileri) boyunca gidilemez, dik kesmek cezalıdır |
| öncül / ardıl, kritik yol | En erken başlangıç, topolojik sıra, zincirleme plan (`planla`: ileri itme; `koruPlani`: gün farkını koruyarak kaydırma), geriye doğru geç başlangıç/bitiş hesabı ve bolluk |
| süzme / çakışma / çizelge modeli | Süzgeçler, çakışma ölçütleri, `cizelgeModel` (satır ve şeritler; ekran ve PNG ortak) |
| çizim | Çizelge (DOM), çubuk taşıma / süre değiştirme, sürükleyerek bağ, liste, yük, rapor, akış şeması (SVG) |
| ayrıntı / form / ayarlar | Ayrıntı penceresi, sağ çekmece formu, zincir onayı, temel plan, proje ayarları |
| Excel / PNG | İçe aktarma (başlık/sütun tahmini, grup adları, özel alanlar), dışa aktarma, şablon; `cizelgeCanvas` (PNG ve rapor görseli) |
| masaüstü (exe) sürümü | `window.masaustu` köprüsü varsa dosya tabanlı çalışma |

`window.__fzc` test için iç durumu ve birkaç yardımcıyı dışarı açar (salt okunur kullanım amaçlı).

### Proje dosyası (`.fzc`, v3)

```json
{
  "v": 3,
  "kayit": "2026-10-03T09:00:00Z",
  "ayar": { "groupBy": "wbs", "colorBy": "sorumlu", "clashBy": "sk", "ppd": 24, "tab": "cizelge",
            "showLinks": true, "showKritik": true, "showTemel": true, "kapaliGruplar": [] },
  "proje": {
    "haftaSonu": [0, 6], "trTatil": true,
    "tatiller": [ { "tarih": "2026-10-30", "ad": "Köprü günü" } ],
    "kapasite": 40,
    "alanlar": [ { "ad": "Standart", "tur": "metin", "secenekler": [] } ],
    "temel": { "kayit": "2026-09-25T10:00:00Z", "plan": { "o3": { "bas": "2026-10-12", "bit": "2026-10-14" } } }
  },
  "faaliyetler": [
    { "id": "o3", "no": 6, "tur": "is", "ad": "Titreşim testi – X ekseni",
      "baslangic": "2026-10-12", "bitis": "2026-10-14", "takvim": "is",
      "sorumlu": "A. Yılmaz", "proje": "Sistem A", "kaynak": "Titreşim tablası", "saat": 32,
      "aciklama": "", "kapali": false, "yuzde": 0, "ust": "g2",
      "gercekBaslangic": "2026-10-12", "etiketler": ["titreşim"], "ozel": { "Standart": "MIL-STD-810H" },
      "oncul": [ { "id": "m1", "gecikme": 0 } ] }
  ]
}
```

- `tur`: `is`, `kilometre` (başlangıç = bitiş) ya da `grup` (tarihleri alt kayıtlardan hesaplanır, bağ taşımaz).
- `takvim`: `is` (iş günü, varsayılan) ya da `takvim` (takvim günü). Gecikme ardılın birimiyle sayılır.
- `ust`: üst grubun `id`'si. `oncul` bağları kalıcı `id` ile kurulur; `no` kullanıcıya gösterilen sıra numarasıdır.
- v2 dosyaları (v1.1) olduğu gibi okunur: eksik alanlar varsayılan değer alır, `proje` yoksa varsayılan takvim kullanılır.
- Tarayıcı sürümünün "Dosyaya kaydet" çıktısında aynı JSON, HTML içindeki `<script id="gomuluVeri" type="application/json">` etiketinde durur.

### Takvim kuralları

- Ardılın en erken başlangıcı: öncül bitişinden sonraki gün; iş günü birimli ardılda ilk iş gününe oturur, gecikme iş günü olarak eklenir.
- Kilometre taşı gün sonu kabul edilir: öncülü aynı gün bitebilir, ardılı ertesi (iş) günü başlar.
- Dini bayram tarihleri `TR_DINI` tablosundadır (2025–2027). Yeni yıllar eklenirken Diyanet takvimi esas alınmalıdır; arife yarım günleri dikkate alınmaz.

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
3. Etiket olarak sürüm numarası verin (ör. `v1.2.0`).
4. `dist/FaaliyetCizelgesi.exe` ve `app/faaliyet-cizelgesi.html` dosyalarını sürükleyip **Publish release** deyin.
