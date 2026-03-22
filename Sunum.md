# Video Sunum

## Sunum Videosu

> **Video Linki:** [Sunum videosu linki buraya eklenecek](https://example.com)

---

## Sunum Yapısı

### 1. Açılış Konuşması (1-2 dakika)

**Konuşma İçeriği:**
- Grup adının tanıtılması
- Projenin genel tanıtımı
- Projenin amacı ve kapsamı
- Sunumun yapısının kısaca açıklanması

**Örnek Konuşma:**
> "Merhaba, ben Doğan Coşman. DC Software olarak Birebir Ders Randevu Sistemi projesini geliştirdim. Bu proje, öğrencilerin öğretmenlerle birebir ders randevusu alabildiği bir web platformudur. Bugün sizlere projemi ve tüm gereksinimlerin çalışır halini göstereceğim."

---

### 2. Kişisel Tanıtım ve Gereksinim Sunumu

#### Doğan Coşman

**Kişisel Tanıtım (30-45 saniye)**
- İsim: Doğan Coşman
- Grup: DC Software (tek kişi)
- Sorumluluk: Backend (Go + Fiber), Frontend (Vue 3), Veritabanı (PostgreSQL)
- Tüm 10 gereksinimin tasarımı ve implementasyonu

**Gereksinim Sunumu — Demo Senaryoları:**

1. **G1 — Kayıt Olma** (1 dakika)
   - Endpoint: `POST /api/v1/auth/register`
   - Demo: Kayıt ekranından ad, e-posta, şifre girerek yeni hesap oluştur
   - Beklenen: 201 Created, JWT token döner, ana sayfaya yönlendirilir

2. **G2 — Giriş Yapma** (1 dakika)
   - Endpoint: `POST /api/v1/auth/login`
   - Demo: Çıkış yap, ardından kayıtlı e-posta/şifre ile tekrar giriş yap
   - Beklenen: 200 OK, JWT token döner, ana sayfaya yönlendirilir

3. **G3 — Öğrenci Profilini Görüntüleme** (30 saniye)
   - Endpoint: `GET /api/v1/students/{studentId}`
   - Demo: Navbar'dan "Profilim" sayfasına git, ad/e-posta/kayıt tarihi göster

4. **G4 — Tüm Öğrencileri Listeleme** (30 saniye)
   - Endpoint: `GET /api/v1/students`
   - Demo: Navbar'dan "Öğrenciler" sayfasına git, kayıtlı öğrenci listesini göster

5. **G5 — Tüm Öğretmenleri Listeleme** (30 saniye)
   - Endpoint: `GET /api/v1/teachers`
   - Demo: Ana sayfadaki öğretmen listesini göster (ad, ders, biyografi)

6. **G6 — Randevu Alma** (1 dakika)
   - Endpoint: `POST /api/v1/appointments`
   - Demo: "Randevularım" sayfasından yeni randevu oluştur (öğretmen seç, tarih/saat belirle)
   - Beklenen: 201 Created, randevu listede görünür

7. **G7 — Randevuları Listeleme** (30 saniye)
   - Endpoint: `GET /api/v1/appointments`
   - Demo: "Randevularım" sayfasında oluşturulan randevuların listesini göster

8. **G8 — Randevu Güncelleme** (1 dakika)
   - Endpoint: `PUT /api/v1/appointments/{appointmentId}`
   - Demo: Mevcut bir randevunun "Düzenle" butonuna bas, tarih/saati değiştir, kaydet
   - Beklenen: 200 OK, güncellenen bilgi listede yansır

9. **G9 — Randevu İptal Etme** (30 saniye)
   - Endpoint: `DELETE /api/v1/appointments/{appointmentId}`
   - Demo: Bir randevunun "İptal Et" butonuna bas, onayla
   - Beklenen: 204 No Content, randevu durumu "İptal" olarak görünür

10. **G10 — Randevu Önerisi Alma** (1-2 dakika)
    - Endpoint: `POST /api/v1/appointments/suggest`
    - Demo: "AI Öneri" sayfasından ders adı (Matematik) ve müsait zaman dilimleri gir, "Öneri Al" butonuna bas
    - Beklenen: Uygun öğretmen + tarih/saat önerileri listelenir, birini seçerek randevu al

---

### 3. Kapanış Konuşması (1 dakika)

**Konuşma İçeriği:**
- Tüm gereksinimlerin tamamlandığının özeti
- Projenin başarıyla tamamlandığının vurgulanması

**Örnek Konuşma:**
> "Bugün sizlere Birebir Ders Randevu Sistemi projemi sundum. Tüm 10 fonksiyonel gereksinim başarıyla tamamlandı ve çalışır durumda gösterildi. Backend Go + Fiber ile, frontend Vue 3 ile geliştirildi, veritabanı olarak PostgreSQL kullanıldı. Teşekkürler!"

---

## Sunum Hazırlık Kontrol Listesi

### Genel Hazırlık
- [ ] Açılış konuşması hazırlandı
- [ ] Tüm gereksinimler çalışır durumda
- [ ] Demo senaryoları hazırlandı
- [ ] Test verileri ve hesaplar hazırlandı (seed öğretmenler + test öğrenci)

### Teknik Hazırlık
- [ ] Video kayıt cihazı/kamera hazır
- [ ] Mikrofon kalitesi test edildi
- [ ] Işıklandırma uygun
- [ ] Arka plan düzenlendi
- [ ] Ekran kayıt yazılımı hazır (demo için)
- [ ] Backend çalışıyor (`make run`)
- [ ] Frontend çalışıyor (`yarn dev`)
- [ ] Veritabanı hazır (migration + seed yapıldı)

### Kişisel Hazırlık
- [ ] Sunum prova edildi
- [ ] Konuşma süresi kontrol edildi (toplam ~12-15 dakika)
- [ ] Demo akışı prova edildi

---

## Video Çekim Teknikleri

### Kişisel Tanıtım Bölümü
- **Kamera Açısı:** Yüz net görünecek şekilde
- **Işık:** Yüzün iyi aydınlatıldığından emin olun
- **Arka Plan:** Temiz ve profesyonel görünüm
- **Görüntü:** Omuz üstü çekim
- **Göz Teması:** Kameraya bakarak konuşun

### Demo Bölümü
- **Ekran Kaydı:** Net ve yüksek çözünürlükte
- **Ses:** Demo sırasında açıklama yapın
- **Hız:** Yavaş ve anlaşılır hareket edin
- **Vurgu:** Önemli noktaları işaret edin

---

## Zaman Yönetimi

- **Açılış:** 1-2 dakika
- **Kişisel tanıtım:** 30-45 saniye
- **Gereksinim sunumu:** ~10 dakika (her gereksinim 30sn-2dk)
- **Kapanış:** 1 dakika
- **Toplam Süre:** Yaklaşık 12-15 dakika
