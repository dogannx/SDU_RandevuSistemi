# SDU_RandevuSistemi

| No | Gereksinim Adı                 | Açıklama                                                                                                                                     | İşlem Türü
|----|--------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------|------------
| 1  | Kayıt olma                     | Öğrenci, sisteme adını, e‑postasını ve şifresini girerek yeni bir hesap oluşturabilmelidir.                                                 | POST
| 2  | Giriş yapma                    | Öğrenci, daha önce oluşturduğu e‑posta ve şifre bilgileriyle sisteme giriş yapabilmelidir.                                                  | POST
| 3  | Öğrenci profilini görüntüleme  | Öğrenci, kendi adını, e‑postasını ve temel bilgilerini içeren profil sayfasını görebilmelidir.                                              | GET
| 4  | Tüm öğrencileri listeleme      | Yetkili kullanıcı, sisteme kayıtlı tüm öğrencilerin ad ve temel bilgilerini liste halinde görebilmelidir.                                   | GET
| 5  | Tüm öğretmenleri listeleme     | Öğrenci, sistemdeki tüm öğretmenlerin adlarını ve verdikleri dersleri liste halinde görebilmelidir.                                         | GET
| 6  | Randevu alma                   | Öğrenci, seçtiği öğretmen ve uygun saat için bir randevu oluşturabilmelidir.                                                                | POST
| 7  | Randevuları listeleme          | Öğrenci, sadece kendisine ait geçmiş ve gelecek tüm randevuları liste halinde görebilmelidir.                                               | GET
| 8  | Randevu güncelleme             | Öğrenci, daha önce aldığı bir randevunun saatini veya tarihini uygun başka bir zamanla değiştirebilmelidir.                                 | PUT
| 9  | Randevu iptal etme             | Öğrenci, artık gitmek istemediği bir randevuyu sistemden iptal edebilmelidir.                                                               | DELETE
| 10 | Randevu önerisi alma           | Öğrenci, haftalık boş zamanlarını ve ders adını seçerek kendisi için en uygun randevu saatleri için öneri alabilmelidir (yapay zeka tabanlı). | POST
