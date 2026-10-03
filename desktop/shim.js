/* Exe köprüsü: sayfaya window.masaustu arayüzünü sağlar (dosya pencereleri, kayıt, menü). */
(function(){
  if (window.masaustu) return;
  var bekleyen = {}, sayac = 0;
  window.__ncb = function(id, sonuc){ var f = bekleyen[id]; delete bekleyen[id]; if (f) f(sonuc); };
  function diyalog(tur, baslik, ad, filtre){
    return new Promise(function(coz){ var id = ++sayac; bekleyen[id] = coz; window.nativeDiyalog(id, tur, baslik || "", ad || "", filtre || ""); });
  }
  function soru(mesaj, detay, dugme, tur){
    return new Promise(function(coz){ var id = ++sayac; bekleyen[id] = coz; window.nativeSoru(id, mesaj || "", detay || "", dugme, tur || ""); });
  }
  function b64Bayt(b64){ var s = atob(b64), u = new Uint8Array(s.length); for (var i = 0; i < s.length; i++) u[i] = s.charCodeAt(i); return u; }
  function blobB64(blob){
    return new Promise(function(coz, red){ var r = new FileReader(); r.onload = function(){ coz(String(r.result).split(",")[1] || ""); }; r.onerror = red; r.readAsDataURL(blob); });
  }
  var PROJE = "Faaliyet çizelgesi (*.fzc)|*.fzc|Eski kayıtlar (*.json;*.html)|*.json;*.html;*.htm|Tüm dosyalar (*.*)|*.*";
  var TURLER = { xlsx:"Excel çalışma kitabı", png:"PNG görseli", json:"JSON dosyası", html:"HTML sayfası", csv:"CSV dosyası", fzc:"Faaliyet çizelgesi" };
  function tekFiltre(uz, tur){ return (tur || TURLER[uz] || uz.toUpperCase()) + " (*." + uz + ")|*." + uz + "|Tüm dosyalar (*.*)|*.*"; }

  var menuCb = null, bildirimCb = null;
  window.__menu = function(k){ if (menuCb) menuCb(k); };

  window.masaustu = {
    baslangic: function(){ return window.nativeBaslangic(); },
    ac: function(yol){ return yol ? window.nativeOku(yol) : diyalog("ac", "Aç", "", PROJE); },
    ikiliAc: function(s){
      s = s || {};
      var uz = (s.uzantilar || ["*"]).map(function(e){ return "*." + e; }).join(";");
      return diyalog("acIkili", s.baslik || "Aç", "", (s.tur || "Dosya") + " (" + uz + ")|" + uz + "|Tüm dosyalar (*.*)|*.*")
        .then(function(r){ return r ? { ad:r.ad, veri:b64Bayt(r.b64) } : null; });
    },
    kaydetYeri: function(s){
      s = s || {};
      var uz = s.uzanti || "fzc";
      return diyalog("kaydetYeri", s.baslik || "Farklı kaydet", s.ad || ("Yeni çizelge." + uz), uz === "fzc" ? PROJE : tekFiltre(uz, s.tur));
    },
    yaz: function(yol, icerik, projeMi){ return window.nativeYaz(yol, icerik, false, !!projeMi); },
    soru: function(s){
      var n = (s.dugmeler || []).length;
      return soru(s.mesaj, s.detay, n >= 3 ? 3 : n, s.tur);
    },
    kurtarmaYaz: function(v){ window.nativeKurtarmaYaz(v); },
    kurtarmaOku: function(){ return window.nativeKurtarmaOku(); },
    kurtarmaSil: function(){ window.nativeKurtarmaSil(); },
    kirli: function(d){ window.nativeKirli(!!d, document.title); },
    kapat: function(){ window.nativeKapat(); },
    konumuAc: function(yol){ window.nativeKonum(yol); },
    dosyaYolu: function(){ return ""; },
    menu: function(cb){ menuCb = cb; },
    bildirim: function(cb){ bildirimCb = cb; }
  };

  /* Excel / PNG / JSON indirmeleri: Windows'un Farklı kaydet penceresi */
  var bloblar = {};
  var ozgunOlustur = URL.createObjectURL.bind(URL);
  URL.createObjectURL = function(b){ var u = ozgunOlustur(b); if (b instanceof Blob) bloblar[u] = b; return u; };
  function indir(blob, ad){
    var uz = (ad.split(".").pop() || "").toLowerCase();
    diyalog("kaydetYeri", "Farklı kaydet", ad, tekFiltre(uz)).then(function(yol){
      if (!yol) return;
      return blobB64(blob).then(function(b64){ return window.nativeYaz(yol, b64, true, false); }).then(function(r){
        if (r && r.ok && bildirimCb) bildirimCb("Kaydedildi: " + r.ad);
      });
    });
  }
  var ozgunTikla = HTMLAnchorElement.prototype.click;
  HTMLAnchorElement.prototype.click = function(){
    if (this.hasAttribute("download") && bloblar[this.href]) { indir(bloblar[this.href], this.getAttribute("download") || "dosya"); return; }
    return ozgunTikla.call(this);
  };
  document.addEventListener("click", function(e){
    var a = e.target && e.target.closest ? e.target.closest("a[download]") : null;
    if (a && bloblar[a.href]) { e.preventDefault(); indir(bloblar[a.href], a.getAttribute("download") || "dosya"); }
  }, true);

  /* alert / confirm: uygulama adlı Windows mesaj kutusu (senkron) */
  function mesajKutusu(k, metin){
    try {
      var x = new XMLHttpRequest();
      x.open("POST", location.pathname.replace(/\/?$/, "/") + "__mesaj?k=" + k, false);
      x.send(String(metin == null ? "" : metin));
      return x.responseText === "1";
    } catch (e) { return k === "confirm" ? false : true; }
  }
  window.alert = function(m){ mesajKutusu("alert", m); };
  window.confirm = function(m){ return mesajKutusu("confirm", m); };

  /* kısayollar; sayfa yenileme ve tarayıcı sağ tık menüsü kapalı (kaydedilmemiş iş kaybolmasın) */
  /* Ctrl+Z / Ctrl+Y sayfanın kendisinde işlenir (yazı alanında metin geri alma bozulmasın) */
  var KISA = { n:"yeni", o:"ac", s:"kaydet", e:"excelKaydet", i:"excelAktar", p:"yazdir",
               "1":"sekme:cizelge", "2":"sekme:liste", "3":"sekme:akis", "4":"sekme:yuk", "5":"sekme:rapor" };
  window.addEventListener("keydown", function(e){
    var k = (e.key || "").toLowerCase();
    if (k === "f5" || ((e.ctrlKey || e.metaKey) && k === "r")) { e.preventDefault(); return; }
    if (!(e.ctrlKey || e.metaKey) || e.altKey) return;
    if (k === "s" && e.shiftKey) { e.preventDefault(); window.__menu("farkli"); return; }
    if (e.shiftKey) return;
    if (KISA[k]) { e.preventDefault(); window.__menu(KISA[k]); }
  }, true);
  window.addEventListener("contextmenu", function(e){
    var t = e.target;
    if (t && (t.isContentEditable || /^(INPUT|TEXTAREA)$/.test(t.tagName))) return;
    e.preventDefault();
  });
})();
