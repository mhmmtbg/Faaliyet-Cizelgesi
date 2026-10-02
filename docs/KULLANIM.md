# Kullanım

## Faaliyet ekleme ve düzenleme

Soldaki formda **faaliyet adı, başlangıç, bitiş, sorumlu, proje, adam*saat** ve açıklama girilir. Bitiş boş bırakılırsa faaliyet tek günlük sayılır. Sorumlu ve proje alanları daha önce girilen adları önerir; aynı kişinin farklı yazımları ayrı renk almasın diye önerileri kullanın.

Her faaliyete otomatik bir numara verilir (`#1`, `#2`…). Numaralar listede, çizelgede, akış şemasında ve Excel'de görünür; öncül bağlarken bu numaralar kullanılır.

Listedeki satır düğmeleri:

| Düğme | İşlev |
|---|---|
| ✎ | Formu doldurur, düzenlemeye açar |
| ✓ / ↺ | Faaliyeti kapatır / yeniden açar. Kapatılan iş çizelgeden ve çakışma hesabından düşer, kayıt kalır. |
| ✕ | Kaydı tamamen siler |

Faaliyet adına (listede) ya da çubuğa (çizelgede) tıklamak **ayrıntı penceresini** açar.

## Zaman çizelgesi

- **Görünüm:** Tüm faaliyetler, Çeyrek (‹ › ile dönem değiştirme) veya Tarih aralığı. Pencerenin dışına taşan işler kenarda kesik görünür.
- **Satırlar / Renk:** Sorumluya veya projeye göre.
- **Çakışma kontrolü:** Aynı sorumlu, aynı proje ya da ikisinden biri.
- **Filtreler ve arama** tüm sekmeleri etkiler.
- **Ölçek:** `−` `+`, kaydırıcı ve **Sığdır**.
- **Bağlantılar → okları göster** kutusu öncül–ardıl oklarını açıp kapatır.

## Öncül / ardıl bağımlılıkları

Bağ türü *bitişten başlangıca*dır: öncül bitince ardıl başlar. İsteğe bağlı **gecikme** (gün) eklenebilir; negatif gecikme örtüşmeye izin verir.

**Bağ kurmanın üç yolu:**

1. **Sürükle-bırak:** Çizelgede bir çubuğu ya da akış şemasında bir kutuyu tutup başka birinin üzerine bırakın. Tuttuğunuz öncül, bıraktığınız ardıl olur. Geçerli hedef yeşil, geçersiz hedef (döngü ya da zaten var olan bağ) kırmızı çerçeve alır. `Esc` iptal eder. Bildirimdeki **Geri al** bağı kaldırır.
2. **Form:** "Öncül faaliyetler" bölümünden seçip gecikme yazın ve **Ekle**'ye basın.
3. **Excel:** "Öncül" sütununa aynı dosyadaki numaraları yazın: `3`, `3; 5`, gecikmeli `3+2`.

**Tarih değişince:** Bir işin tarihi değişirse zincirdeki ardılların yeni tarihleri hesaplanır ve bir onay penceresinde gösterilir (şu anki / yeni tarih / kayma).

- Ardıllar yalnızca öncüllerinden önce başlamaya kalkarlarsa ileri itilir. Aradaki boşluk tamponu "yer", boşluğu olan ardıllar yerinde kalır.
- "Öncülüne yaslı duran ardılları öne de çek" kutusu, öncül öne alındığında bitişik ardılları da öne getirir.
- Süreler korunur. Kapatılmış faaliyetler taşınmaz. Gecikme takvim günüdür (hafta sonu atlanmaz).

**İhlaller:** Bir işi öncülünden önceye elle koyarsanız iş "⛓ bağımlılık ihlali" olarak işaretlenir ve oku kırmızı kesikli çizilir. Ayrıntı penceresindeki **Bağımlılığa göre hizala** ya da lejanttaki **Hepsini hizala** düzeltir.

## Akış şeması

Her proje ayrı bir bant olarak çizilir. Faaliyetler tarih sırasına göre aşamalara dizilir; tarihleri örtüşenler aynı aşamada alt alta durur.

- **Gri oklar:** aşama sırası
- **Yeşil oklar:** öncül → ardıl bağları
- **Kırmızı kesikli ok:** bağımlılık ihlali
- **Kırmızı çerçeve ve ⚠:** çakışma
- **Kesikli çerçeve:** kapatılmış iş

Üstteki **Proje** seçicisiyle tek bir projenin şeması gösterilir. Kutuya tıklamak ayrıntı penceresini açar; sürükleyip bırakmak bağ kurar.

## Adam*saat

Kişi bazında toplam adam*saat, adam*gün (8 saat) karşılığı, faaliyet sayısı ve projelere dağılım. "Kapatılanları da göster" kutusu kapsamı belirler.

## Excel

- **İçe aktarma:** Dosya seçilince başlık satırı ve sütunlar otomatik eşleştirilir; yanlış eşleşme açılır listelerden düzeltilir. Önizleme okunan, atlanan (tarihsiz) satırları ve öncül bağlarını gösterir. Tarihler Excel tarih hücresi, `12.03.2026`, `2026-03-12`, `12 Mart 2026` biçimlerinde okunur. "Durum" sütununda *kapandı / tamamlandı / bitti* yazan satırlar kapatılmış gelir.
- **Dışa aktarma:** Faaliyetler (no, öncül, durum, çakışma, bağımlılık sütunlarıyla), adam*saat özeti ve çakışmalar ayrı sayfalarda.
- **Şablon:** Doğru sütun başlıklarıyla boş çalışma kitabı.

## Kaydetme ve dosyalar

### Masaüstü (exe)

- **Dosya** menüsü: Yeni (`Ctrl+N`), Aç (`Ctrl+O`), Son açılanlar, Kaydet (`Ctrl+S`), Farklı kaydet (`Ctrl+Shift+S`), Excel'den aktar (`Ctrl+I`), Dışa aktar, Yazdır, Dosya konumunu aç.
- Projeler `.fzc` dosyasına kaydedilir. Başlıkta `•` kaydedilmemiş değişiklik olduğunu gösterir; kapatırken sorulur.
- Program beklenmedik şekilde kapanırsa bir sonraki açılışta kaydedilmemiş çalışmayı geri yüklemeyi önerir.
- `.fzc` dosyasını exe'nin üzerine sürüklemek o dosyayla açar.
- Excel ve PNG kaydederken klasör sorulur; son kullanılan klasör hatırlanır.
- Ayarlar ve kurtarma kaydı: `%LOCALAPPDATA%\FaaliyetCizelgesi`
- Sekmeler: `Ctrl+1` liste, `Ctrl+2` adam*saat, `Ctrl+3` akış şeması.

### Tarayıcı (HTML)

- Her değişiklik tarayıcının yerel deposuna otomatik kaydedilir.
- **Dosyaya kaydet**, verileri HTML'in içine gömerek yeni bir dosya üretir. Bu dosyayı kime gönderirseniz kayıtlar da onunla gider. Chrome/Edge'de dosya bir kez seçildikten sonra her değişiklik otomatik olarak o dosyaya yazılır.
- **JSON aç / kaydet** ile `.fzc` ve `.json` dosyaları okunur ve yazılır.

İki sürüm aynı dosya biçimini kullanır: exe'nin `.fzc` dosyaları HTML'de, HTML'in kaydettiği dosyalar exe'de açılır. Exe'deki **Dışa aktar → Paylaşılabilir HTML**, programı olmayan birine gönderilebilecek, verileri içinde HTML üretir.
