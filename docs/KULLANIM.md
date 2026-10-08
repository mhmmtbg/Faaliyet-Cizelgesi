# Kullanım

## Ekran düzeni

| Bölge | İçerik |
|---|---|
| Üst çubuk | Dosya bilgisi, ↶ Geri al / ↷ Yinele, **+ Yeni faaliyet**, **Dosya** menüsü, **⚙ Proje ayarları** |
| Sekmeler | **Çizelge**, **Liste**, **Akış şeması**, **Yük ve adam*saat**, **Rapor** |
| Süzgeçler (sekmelerin sağında) | Sorumlu, proje, kaynak, etiket, metin arama ve "Kapatılanlar". Tüm sekmelere uygulanır. |
| Sağ çekmece | Faaliyet ekleme ve düzenleme formu. `Esc` ile kapanır. |

Kısayollar: `Ctrl+Z` geri al, `Ctrl+Y` yinele, `Insert` yeni faaliyet, `Alt+1…5` sekmeler (masaüstünde `Ctrl+1…5`), `Ctrl+tekerlek` çizelgede ölçek, haftalık görünümde `←` `→` önceki / sonraki hafta, `Esc` pencere / çekmece kapatır.

## Faaliyet ekleme ve düzenleme

**+ Yeni faaliyet** sağdaki çekmeceyi açar. Üç tür kayıt vardır:

| Tür | Açıklama |
|---|---|
| **İş** | Başlangıç, bitiş ve süresi olan faaliyet |
| **◆ Kilometre taşı** | Tek tarih (ör. TRR, müşteri kabulü). Gün sonu kabul edilir: öncülü aynı gün bitebilir, ardılı ertesi iş günü başlar. |
| **▤ Grup / özet** | İş paketi. Tarihi, adam*saati ve ilerlemesi alt kayıtlarından hesaplanır. |

Form alanları:

- **Üst grup:** Kaydı bir gruba bağlar (gruplar iç içe olabilir).
- **Başlangıç / Bitiş / Süre:** Süre değişince bitiş, bitiş değişince süre güncellenir.
- **Süre birimi:**
  - *İş günü* (varsayılan): hafta sonu ve tatiller sayılmaz.
  - *Takvim günü*: kesintisiz süren işler (ör. 10 günlük termal çevrim) için her gün sayılır.
- **Sorumlu, Proje, Kaynak / tesis:** Daha önce girilen adları önerir. Kaynak, test tezgâhı, oda, cihaz gibi paylaşılan bir varlıktır.
- **Adam*saat, Etiketler** (virgülle).
- **İlerleme %**, **gerçek başlangıç / bitiş**.
- **Özel alanlar:** Proje ayarlarında tanımlananlar.
- **Açıklama**, **Öncül faaliyetler**.

Yeni kayıt eklendikten sonra form boşalır ve aynı sorumlu / proje / grup ile bir sonraki iş günü önerilir; art arda giriş için uygundur.

Her kayda otomatik bir numara verilir (`#1`, `#2`…). Numaralar her yerde ve Excel'de görünür; öncül bağlarken kullanılır.

## Zaman çizelgesi

- **Görünüm:** Tüm faaliyetler, **Haftalık**, Çeyrek (‹ › ile dönem değiştirme) ya da Tarih aralığı.
- **Haftalık görünüm:**
  - Pazartesi–pazar tek hafta gösterilir; ölçek 7 günü ekrana sığdırır (pencere boyutu değişince yeniden sığar).
  - ‹ › düğmeleri ya da `←` `→` tuşları önceki / sonraki haftaya geçer. **Bu hafta** içinde bulunulan haftaya, **Tarihe git…** seçilen tarihin haftasına gider.
  - Gün başlıklarında gün adı ve tarih, geniş çubuklarda tarih aralığı, süre ve ilerleme yazar. Hafta dışına taşan işler kenarda kesik görünür.
  - `+` / `−` ile daha da genişletilebilir; **Sığdır** yeniden 7 güne döner.
  - İş kırılımı satırlarında yalnızca o haftada görünen faaliyetler ve grupları listelenir.
- **Satırlar:** Sorumlu, proje, kaynak, **iş kırılımı** (her kayıt bir satır; gruplar ▾/▸ ile açılıp kapanır, ⊞/⊟ hepsini açar/kapatır) ya da tek satır.
- **Renk:** Sorumlu, proje ya da kaynak.
- **Çakışma:** Aynı sorumlu, aynı kaynak, sorumlu veya kaynak, aynı proje ya da sorumlu veya proje. Kilometre taşları ve gruplar çakışmaya girmez.
- **Oklar / Kritik yol / Temel plan** kutuları gösterimi açıp kapatır.
- **Ölçek:** Kaydırıcı, `−` `+` ya da `Ctrl+tekerlek`; gün genişliği 2–240 piksel. **Bugün** düğmesi bugüne kaydırır; **Sığdır** tüm aralığı ekrana oturtur.

Gösterimler:

| İşaret | Anlamı |
|---|---|
| Gri sütun | Hafta sonu |
| Turuncu taralı sütun | Tatil / kapalı gün (üzerine gelince adı görünür) |
| Çubuğun altındaki koyu şerit | İlerleme yüzdesi |
| Kırmızı taralı çubuk, ⚠ | Çakışma |
| Kırmızı sağ kenar | Gecikti: bitiş tarihi geçti, %100 değil |
| Siyah çerçeve, kalın siyah ok | Kritik yol |
| Çubuğun altındaki gri çizgi | Temel plandaki tarih |
| ◆ ve kesikli dikey çizgi | Kilometre taşı |
| Koyu özet çubuğu | Grup |
| Soluk gri, kesikli kenar, ✓ | Kapatılmış faaliyet (yalnızca "Kapatılanlar" işaretliyken; taşınamaz) |

**Fareyle düzenleme:**

| Hareket | Sonuç |
|---|---|
| Çubuğa tıklama | Ayrıntı penceresi |
| Çubuğu sürükleme | Tarihleri taşır. İş günü faaliyetlerinde süre iş günü olarak korunur, hafta sonuna bırakılan çubuk iş gününe oturur. |
| Çubuğun sağ kenarını sürükleme | Bitişi, yani süreyi değiştirir |
| Çubuğun sağındaki ○ tutamağı başka çubuğa sürükleme | Öncül → ardıl bağı kurar |

Taşıma ardılları etkiliyorsa zincir onay penceresi açılır. `Esc` sürüklemeyi iptal eder.

## Öncül / ardıl bağları

Bağ türü *bitişten başlangıca*dır: öncül bitince ardıl başlar.

- İsteğe bağlı **gecikme** girilebilir. İş günü birimli ardılda iş günü, takvim günü birimlide gün sayılır. Negatif gecikme örtüşmeye izin verir.
- Ardıl, öncülün bitişinden sonraki ilk iş gününde başlayabilir.

**Bağ kurmanın üç yolu:**

1. **Sürükle-bırak:** Çizelgede ○ tutamaktan ya da akış şemasında kutudan başka bir kayda bırakın. Geçerli hedef yeşil, geçersiz hedef (döngü, var olan bağ, grup) kırmızı çerçeve alır.
2. **Form:** "Öncül faaliyetler" bölümünden seçip gecikme yazın ve **Ekle**'ye basın.
3. **Excel:** "Öncül" sütununa aynı dosyadaki numaraları yazın: `3`, `3; 5`, gecikmeli `3+2`.

**Tarih değişince:** Bağlı faaliyeti olan bir işin tarihi değişince (çubuğu sürükleme, sağ kenarından süre değiştirme, formda kaydetme ya da toplu kaydırma) bir uyarı çıkar: "#6 faaliyetini 2 iş günü ileri kaydırdınız — bağlı faaliyetlerin tarihleri de güncellensin mi?"

Pencerede etkilenecek faaliyetler, şu anki ve yeni tarihleriyle listelenir. Seçenekler:

| Seçenek | Ne yapar |
|---|---|
| **Bağlı faaliyetleri de aynı miktarda kaydır** (varsayılan) | Ardıllar, işin bitişi ne kadar kaydıysa o kadar kayar. Aradaki gün farkı korunur; işi öne alırsanız ardıllar da öne gelir. |
| **Öncülleri de aynı miktarda kaydır** | İşaretlenirse öncüller de başlangıç kayması kadar kayar. |
| **Yalnızca gerekenleri ileri it** | Eski davranış: yalnızca öncülünden önce başlamaya kalkan ardıllar itilir, boşluğu olanlar yerinde kalır. |
| **Yalnızca bu faaliyeti taşı / kaydet** | Bağlı faaliyetlere dokunmaz. |

- Kayma, iş günü birimli faaliyetlerde iş günü (tatiller atlanır), takvim günü birimlilerde gün olarak uygulanır.
- Kapatılmış ve tamamlanmış (%100 ya da gerçek bitişi girilmiş) faaliyetler taşınmaz.
- Kaydırma sonucunda başka bir öncülüyle çakışan iş varsa o da ileri itilir.

**İhlaller:** Bir işi öncülünden önceye koyarsanız iş "⛓ bağımlılık" olarak işaretlenir ve oku kırmızı kesikli çizilir. Düzeltmek için ayrıntıdaki **Bağımlılığa göre hizala**'yı, lejanttaki **Hepsini hizala**'yı ya da **Dosya → Bağımlılık ihlallerini düzelt**'i kullanın.

## Kritik yol ve bolluk

- **Bolluk:** Bir işin, projenin bitişini kaydırmadan ne kadar gecikebileceğidir (iş günü). Listede ve ayrıntıda görünür.
- **Kritik yol:** Bolluğu 0 ya da eksi olan işler. Çizelgede siyah çerçeve ve siyah ok, akış şemasında kalın çerçeve, listede ● ile gösterilir.
- Hesap her proje için ayrı yapılır. Bağı olmayan bir işin sınırı kendi projesinin bitişidir. Kapatılmış işler hesaba girmez.

## Temel plan

**Dosya → Temel planı kaydet** o anki tarihleri saklar. Sonra:

- **Temel plan** kutusu çizelgede eski tarihleri gri çizgiyle gösterir.
- Listede **Sapma** sütunu bitişin kaç gün kaydığını gösterir. Aynı bilgi ayrıntıda, raporda ve Excel'de de yer alır.
- Temel plan yeniden kaydedilebilir ya da **Proje ayarları**'ndan silinebilir.

## Liste ve toplu düzenleme

- Varsayılan sıra **iş kırılımı** sırasıdır. Grupların altındaki kayıtlar girintili gösterilir.
- Sütun başlığına tıklamak sıralar. "No" iki kez tıklanınca iş kırılımı sırasına döner.
- ☐ kutularıyla seçim yapılır. Üstte beliren çubukta **Toplu düzenle**, **Kapat**, **Yeniden aç**, **Sil** bulunur.
- Toplu düzenlemede boş bırakılan alanlar değişmez. Değiştirilebilenler:
  - sorumlu, proje, kaynak, üst grup;
  - etiket ekleme ve kaldırma;
  - ilerleme %, durum;
  - **N gün kaydırma** (iş günü faaliyetlerinde iş günü; ardılları etkilerse zincir sorulur).

## Akış şeması

Her proje ayrı bir bant olarak çizilir. Tarihleri örtüşen faaliyetler aynı aşamada alt alta durur.

| İşaret | Anlamı |
|---|---|
| Gri oklar | Aşama sırası |
| Yeşil oklar | Öncül → ardıl bağları |
| Kırmızı kesikli ok | Bağımlılık ihlali |
| Kırmızı çerçeve ve ⚠ | Çakışma |
| Kalın siyah çerçeve | Kritik yol |
| Kutunun altındaki yeşil şerit | İlerleme |
| ◆ | Kilometre taşı |

**Öncül okları** kutusu yeşil bağ oklarını açıp kapatır. Oklar açıkken kutular arasında oklara ayrı koridor bırakılır. Oklar kutuların ve gri aşama çizgilerinin üzerinden geçmez; gerekirse gri çizgiyi dik keser.

Gruplar akış şemasında gösterilmez. Kutuya tıklamak ayrıntı penceresini açar; kutuyu sürükleyip başka kutuya bırakmak bağ kurar.

## Yük ve adam*saat

- **Kişi yükü:**
  - Her faaliyetin adam*saati, süresindeki iş günlerine eşit dağıtılır ve haftalara toplanır.
  - Hücre rengi kapasiteye göre değişir: %50, %100, aşım, %125 üstü.
  - Haftalık kapasite (varsayılan 40 saat) tatil haftalarında iş günü oranında azalır.
- **Kaynak doluluğu:** Her kaynağın haftada kaç gün dolu olduğunu gösterir. Aynı gün birden fazla faaliyet aynı kaynağı kullanıyorsa hücre kırmızıdır.
- Kişi ya da kaynak adına tıklamak o kişiye / kaynağa süzer.
- Altta kişi × proje adam*saat özeti vardır.

## Rapor

**Dönem başı** ve **süre** (1, 2 ya da 4 hafta) seçilir. Üstteki süzgeçler rapora da uygulanır.

Rapor şunları içerir:

- özet kartlar ve genel ilerleme;
- dönem çizelgesi;
- geciken ve başlaması geciken faaliyetler;
- dönemde başlayacak ve bitecek işler;
- kilometre taşları;
- kritik yol;
- çakışmalar;
- temel plandan sapmalar.

**Raporu yazdır / PDF** yalnızca raporu yazdırır. PDF için yazıcı olarak "PDF olarak kaydet" seçin.

## Proje ayarları (⚙)

- **Çalışılmayan günler:** Varsayılan cumartesi ve pazar.
- **Türkiye resmi tatilleri:**
  - Sabit tatiller her yıl uygulanır.
  - Ramazan ve Kurban bayramı 2025–2027 için tanımlıdır.
  - Planınızda tanımsız bir yıl varsa uyarı çıkar; o yılın bayramlarını Diyanet takviminden bakıp özel tatil olarak ekleyin.
- **Özel tatiller ve kapalı günler:** Tek gün ya da tarih aralığı (ör. idari izin, tesis bakımı).
- **Kapasite:** Kişi başı haftalık saat.
- **Özel alanlar:** Ad ve tür (metin, sayı, tarih, seçim listesi). Formda, ayrıntıda ve Excel'de ek sütun olarak görünür.
- **Temel plan:** Bilgi ve silme.

Takvim değişikliği mevcut tarihleri kaydırmaz; sonraki planlamada kullanılır. Ayar değişiklikleri de geri alınabilir.

## Excel

**İçe aktarma:**

- Başlık satırı ve sütunlar otomatik eşleştirilir. Yanlış eşleşme açılır listelerden düzeltilir.
- Okunabilen sütunlar:
  - No, Öncül, Faaliyet, Başlangıç, Bitiş;
  - Tür (*iş / kilometre taşı / grup*), Durum, Üst grup;
  - Sorumlu, Proje, Kaynak, Adam*saat;
  - İlerleme %, Gerçek başlangıç / bitiş;
  - Süre birimi, Etiketler, Açıklama;
  - özel alanlar (aynı adlı sütundan).
- "Üst grup" sütunu aynı dosyadaki grup satırının adını ya da numarasını alır. Bulunamayan adlar için grup oluşturulur.
- Tarihler Excel tarih hücresi, `12.03.2026`, `2026-03-12`, `12 Mart 2026` biçimlerinde okunur.
- "Durum" sütununda *kapandı / tamamlandı / bitti* yazan satırlar kapatılmış gelir.

**Dışa aktarma:** Faaliyetler iş kırılımı sırasıyla, tüm alanlar, bolluk, kritik, temel plan ve sapma sütunlarıyla yazılır. Ayrıca adam*saat özeti, çakışmalar ve tatiller ayrı sayfalardadır.

**Şablon:** Doğru sütun başlıkları ve grup / iş / kilometre taşı örnekleri içeren boş çalışma kitabı.

## Kaydetme ve dosyalar

### Masaüstü (exe)

- **Dosya** menüsü: Yeni (`Ctrl+N`), Aç (`Ctrl+O`), Son açılanlar, Kaydet (`Ctrl+S`), Farklı kaydet (`Ctrl+Shift+S`), Excel'den aktar (`Ctrl+I`), Dışa aktar, Yazdır (`Ctrl+P`), Dosya konumunu aç.
- **Düzen** menüsü: Geri al, Yinele, Yeni faaliyet, Bağımlılık ihlallerini hizala, Temel planı kaydet, Proje ayarları.
- **Görünüm** menüsü: `Ctrl+1` Çizelge, `Ctrl+2` Liste, `Ctrl+3` Akış şeması, `Ctrl+4` Yük, `Ctrl+5` Rapor.
- Projeler, proje ayarlarıyla birlikte `.fzc` dosyasına kaydedilir. Başlıktaki `•` kaydedilmemiş değişiklik olduğunu gösterir; kapatırken sorulur.
- Program beklenmedik şekilde kapanırsa bir sonraki açılışta kaydedilmemiş çalışmayı geri yüklemeyi önerir.
- `.fzc` dosyasını exe'nin üzerine sürüklemek o dosyayla açar.
- Ayarlar ve kurtarma kaydı `%LOCALAPPDATA%\FaaliyetCizelgesi` klasöründedir.

### Tarayıcı (HTML)

- Her değişiklik tarayıcının yerel deposuna otomatik kaydedilir.
- **Dosyaya kaydet**, verileri HTML'in içine gömerek yeni bir dosya üretir. Bu dosyayı kime gönderirseniz kayıtlar da onunla gider. Chrome/Edge'de dosya bir kez seçildikten sonra her değişiklik otomatik olarak o dosyaya yazılır.
- **Dosya → Proje dosyası aç / kaydet** ile `.fzc` ve `.json` dosyaları okunur ve yazılır.

İki sürüm aynı dosya biçimini kullanır. v1.1 ile kaydedilmiş dosyalar da açılır; eski kayıtlar iş günü birimli sayılır.
