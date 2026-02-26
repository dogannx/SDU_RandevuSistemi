# Gereksinim Analizi

Birebir Ders Randevu Sistemi, öğrencilerin öğretmenlerle birebir ders randevusu alabildiği bir platformdur. Aşağıda sistemin desteklediği 10 fonksiyonel gereksinim listelenmiştir.

| No | Gereksinim Adı | Açıklama | İşlem Türü |
|---|---|---|---|
| G1 | Kayıt Olma | Öğrenci, sisteme adını, e-postasını ve şifresini girerek yeni bir hesap oluşturabilmelidir. | POST |
| G2 | Giriş Yapma | Öğrenci, daha önce oluşturduğu e-posta ve şifre bilgileriyle sisteme giriş yapabilmelidir. | POST |
| G3 | Öğrenci Profilini Görüntüleme | Öğrenci, kendi adını, e-postasını ve temel bilgilerini içeren profil sayfasını görebilmelidir. | GET |
| G4 | Tüm Öğrencileri Listeleme | Yetkili kullanıcı, sisteme kayıtlı tüm öğrencilerin ad ve temel bilgilerini liste halinde görebilmelidir. | GET |
| G5 | Tüm Öğretmenleri Listeleme | Öğrenci, sistemdeki tüm öğretmenlerin adlarını ve verdikleri dersleri liste halinde görebilmelidir. | GET |
| G6 | Randevu Alma | Öğrenci, seçtiği öğretmen ve uygun saat için bir randevu oluşturabilmelidir. | POST |
| G7 | Randevuları Listeleme | Öğrenci, sadece kendisine ait geçmiş ve gelecek tüm randevuları liste halinde görebilmelidir. | GET |
| G8 | Randevu Güncelleme | Öğrenci, daha önce aldığı bir randevunun saatini veya tarihini uygun başka bir zamanla değiştirebilmelidir. | PUT |
| G9 | Randevu İptal Etme | Öğrenci, artık gitmek istemediği bir randevuyu sistemden iptal edebilmelidir. | DELETE |
| G10 | Randevu Önerisi Alma | Öğrenci, haftalık boş zamanlarını ve ders adını seçerek kendisi için en uygun randevu saatleri için yapay zekâ tabanlı öneri alabilmelidir. | POST |

## Gereksinim Dosyası

- [Doğan Coşman'ın Gereksinimleri](Dogan-Cosman/Dogan-Cosman-Gereksinimler.md)
