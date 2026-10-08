# Değişiklik Günlüğü

## v1.3.0 — 2026-10-08

- **Haftalık görünüm:**
  - Çizelgede pazartesi–pazar tek hafta gösterilir; ölçek 7 günü ekrana sığdırır.
  - Hafta geçişi ‹ › düğmeleri, ← → tuşları, "Bu hafta" ve "Tarihe git…" ile yapılır.
  - Gün başlıklarında gün adı ve tarih, geniş çubuklarda tarih aralığı ve süre görünür; bugünün sütunu vurgulanır.
- **Daha geniş ölçek:** Gün genişliği 240 piksele kadar çıkabilir. Kaydırıcı logaritmiktir; `Ctrl` + tekerlek ile yakınlaştırılır.
- **Bağlı faaliyetleri gün farkını koruyarak kaydırma:**
  - Bir faaliyetin tarihi değişince (sürükleme, form, toplu kaydırma) "x iş günü kaydırdınız, bağlı faaliyetler de güncellensin mi?" diye sorulur.
  - Varsayılan seçimde ardıllar aynı miktarda kayar ve aradaki gün farkı korunur.
  - İstenirse öncüller de kaydırılır.
  - Eski davranış ("yalnızca gerekenleri ileri it") seçenek olarak duruyor.
  - Tamamlanmış ve kapatılmış faaliyetler taşınmaz.
- **Akış şeması:**
  - "Öncül okları" kutusuyla oklar açılıp kapanır.
  - Oklar kutuların ve gri aşama çizgilerinin üzerinden geçmez; kutular arasında oklara ayrı koridor bırakılır.
- **İş kırılımı:** Haftalık, çeyrek ve tarih aralığı görünümlerinde yalnızca o pencerede görünen faaliyetler ve grupları listelenir.

## v1.2.0 — 2026-10-03

**Ekran düzeni**
- Pencereye sığan tam ekran düzen: Çizelge, Liste, Akış şeması, Yük ve adam*saat, Rapor ayrı sekmelerde; her sekme kendi içinde kayar
- Faaliyet formu sağdan açılan çekmecede (`+ Yeni faaliyet`, `Insert`); art arda girişte ortak alanlar ve sonraki iş günü önerilir
- Süzgeçler (sorumlu, proje, kaynak, etiket, arama) sekme çubuğunda, tüm sekmelere uygulanır
- Dosya işlemleri tek "Dosya" menüsünde

**Planlama**
- Geri al / yinele (`Ctrl+Z` / `Ctrl+Y`, 100 adım) — taşıma, bağ, toplu düzenleme, silme ve ayar değişiklikleri dahil
- Çubukları sürükleyerek taşıma, sağ kenarından süre değiştirme; ardıllar etkilenirse zincir onayı
- Bağ kurma artık çubuğun sağındaki ○ tutamaktan (akış şemasında kutudan) yapılır
- Çalışma takvimi: hafta sonları, Türkiye resmi tatilleri (dini bayramlar 2025–2027), özel tatil / kapalı günler; süre ve bağ gecikmesi iş günüyle hesaplanır, kesintisiz işler için "takvim günü" seçeneği
- Kritik yol ve bolluk (iş günü); çizelge, akış şeması, liste ve raporda gösterim
- Kilometre taşları (◆), gruplar / iş kırılımı (özet çubuk, aç-kapa, alt kayıtlardan hesaplanan tarih, adam*saat ve ilerleme)
- İlerleme yüzdesi, gerçekleşen başlangıç/bitiş, "gecikti" ve "başlamadı" durumları
- Temel plan kaydı ve sapma (gün)
- Kaynak / tesis alanı; kaynağa göre satır, renk, süzgeç ve çakışma kontrolü
- Etiketler ve proje ayarlarından tanımlanan özel alanlar (metin, sayı, tarih, seçim listesi)
- Listede çoklu seçim ve toplu düzenleme (sorumlu, proje, kaynak, grup, etiket, ilerleme, durum, iş günü kaydırma, silme)

**Raporlama**
- Yük sekmesi: kişi başı haftalık saat yükü (kapasite ve tatillere göre aşım renkleri) ve kaynak doluluğu (aynı gün çift kullanım uyarısı)
- Rapor sekmesi: dönem özeti, ilerleme, gecikenler, başlayacak / bitecek işler, kilometre taşları, kritik yol, çakışmalar, temel plandan sapmalar ve dönem çizelgesi; yalnızca rapor yazdırılır (PDF)
- Excel: tür, üst grup, kaynak, ilerleme, gerçekleşen tarihler, süre birimi, etiketler, özel alanlar, bolluk, kritik, temel plan ve sapma sütunları; içe aktarımda grup adları ve özel alanlar eşleşir; "Tatiller" sayfası
- PNG çıktısı yeni gösterimlerle (tatil gölgesi, kilometre taşı, grup, ilerleme, kritik yol, temel plan)

**Masaüstü**
- Düzen menüsü: Geri al, Yinele, Yeni faaliyet, Temel planı kaydet, Proje ayarları; Görünüm menüsünde 5 sekme (`Ctrl+1…5`)
- Proje ayarları `.fzc` dosyasında saklanır; v1.1 dosyaları olduğu gibi açılır


## v1.1.0 — 2026-10-02

İlk yayımlanan sürüm.

- Tarihe göre zaman çizelgesi; sorumlu/proje bazlı satırlar ve renkler; tüm faaliyetler, çeyrek ve tarih aralığı görünümleri
- Sorumlu veya proje bazlı çakışma tespiti; ayrıntı penceresinde çakışma listesi ve örtüşme şeridi
- Öncül/ardıl bağımlılıkları: gecikme, döngü engeli, zincirleme tarih güncelleme (ileri itme / öne çekme), ihlal işaretleme ve hizalama
- Çizelgede ve akış şemasında sürükleyerek bağ kurma
- Kutuların etrafından dolaşan, dik açılı ok rotalama
- Proje bazlı akış şeması: aşama akışı ve bağ okları, proje seçici, açıklamalar
- Adam*saat özeti (kişi × proje)
- Excel içe/dışa aktarma (No ve Öncül sütunları dahil), boş şablon
- Çizelge ve akış şeması için PNG çıktısı
- Faaliyet kapatma (çizelgeden düşer, kayıt kalır)
- Tarayıcı sürümü: otomatik kayıt, verileri içine gömülü paylaşılabilir HTML
- Kurulumsuz Windows masaüstü sürümü (Go + WebView2, ~7 MB): Windows menü çubuğu, aç/kaydet pencereleri, `.fzc` proje dosyası, son açılanlar, kaydedilmemiş değişiklik uyarısı, çökme kurtarması, exe'ye sürükle-bırak ile açma
