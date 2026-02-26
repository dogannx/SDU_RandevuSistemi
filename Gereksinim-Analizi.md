# Gereksinim Analizi — Birebir Ders Randevu Sistemi

## Gereksinim Kuralları

- Her ekip üyesi en az **1 POST**, **1 GET**, **1 PUT** ve **1 DELETE** gereksinimine sahip olmalıdır.
- Toplam gereksinim sayısı: **10**

---

## Fonksiyonel Gereksinimler

| No | Gereksinim Adı | Açıklama | İşlem Türü | Sorumlu Kişi |
|---|---|---|---|---|
| G1 | Kayıt olma | Öğrenci, sisteme adını, e-postasını ve şifresini girerek yeni bir hesap oluşturabilmelidir. | POST | [Sorumlu Kişi] |
| G2 | Giriş yapma | Öğrenci, daha önce oluşturduğu e-posta ve şifre bilgileriyle sisteme giriş yapabilmelidir. | POST | [Sorumlu Kişi] |
| G3 | Öğrenci profilini görüntüleme | Öğrenci, kendi adını, e-postasını ve temel bilgilerini içeren profil sayfasını görebilmelidir. | GET | [Sorumlu Kişi] |
| G4 | Tüm öğrencileri listeleme | Yetkili kullanıcı, sisteme kayıtlı tüm öğrencilerin ad ve temel bilgilerini liste halinde görebilmelidir. | GET | [Sorumlu Kişi] |
| G5 | Tüm öğretmenleri listeleme | Öğrenci, sistemdeki tüm öğretmenlerin adlarını ve verdikleri dersleri liste halinde görebilmelidir. | GET | [Sorumlu Kişi] |
| G6 | Randevu alma | Öğrenci, seçtiği öğretmen ve uygun saat için bir randevu oluşturabilmelidir. | POST | [Sorumlu Kişi] |
| G7 | Randevuları listeleme | Öğrenci, sadece kendisine ait geçmiş ve gelecek tüm randevuları liste halinde görebilmelidir. | GET | [Sorumlu Kişi] |
| G8 | Randevu güncelleme | Öğrenci, daha önce aldığı bir randevunun saatini veya tarihini uygun başka bir zamanla değiştirebilmelidir. | PUT | [Sorumlu Kişi] |
| G9 | Randevu iptal etme | Öğrenci, artık gitmek istemediği bir randevuyu sistemden iptal edebilmelidir. | DELETE | [Sorumlu Kişi] |
| G10 | Randevu önerisi alma (Yapay zeka tabanlı) | Öğrenci, haftalık boş zamanlarını ve ders adını seçerek kendisi için en uygun randevu saatleri için öneri alabilmelidir. | POST | [Sorumlu Kişi] |

---

## Ekip Üyelerinin Gereksinim Dağılımı

| Ekip Üyesi | Sorumlu Gereksinimler |
|---|---|
| [Grup Üyesi 1] | G1, G2 |
| [Grup Üyesi 2] | G3, G4 |
| [Grup Üyesi 3] | G5, G6 |
| [Grup Üyesi 4] | G7, G8 |
| [Grup Üyesi 5] | G9 |
| [Grup Üyesi 6] | G10 |

> **Not:** Sorumlu kişi sütunlarını ve dağılım tablosunu kendi ekip bilgilerinizle güncelleyin.
