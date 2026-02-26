# Gereksinim Analizi

Birebir Ders Randevu Sistemi için tüm gereksinimler aşağıda listelenmiştir. Her gereksinim bir API metoduna karşılık gelecek şekilde tanımlanmıştır.

## Gereksinim Sayıları (En Az)

- **1 Kişi:** 10 gereksinim
- **2 Kişi:** 16 gereksinim
- **3 Kişi:** 21 gereksinim
- **4 Kişi:** 24 gereksinim
- **5 Kişi:** 30 gereksinim

## Gereksinimlerde Uyulması Gereken Kurallar

1. **İsimler anlamlı olmalı:** Gereksinim isimleri net ve anlaşılır olmalıdır.
2. **Açıklamalar net olmalı:** Her gereksinimin açıklaması açık ve anlaşılır şekilde yazılmalıdır.
3. **Açıklamalar teknik jargon ve kısaltmalar içermemeli:** Gereksinim açıklamaları herkesin anlayabileceği basit bir dille yazılmalıdır.
4. **Gereksinim isimleri çok uzun olmamalı ve bir eylem bildirmeli:**
   - İsimler kısa ve öz olmalıdır
   - Bir eylem fiili içermelidir
   - Örnekler: "Kayıt Olma", "Giriş Yapma", "Profil Güncelleme", "Hesap Silme"

# Tüm Gereksinimler

| No | Gereksinim Adı | Açıklama | İşlem Türü | Sorumlu Kişi |
|---|---|---|---|---|
| G1 | Kayıt Olma | Öğrenci, sisteme adını, e-postasını ve şifresini girerek yeni bir hesap oluşturabilmelidir. | POST | [Sorumlu Kişi] |
| G2 | Giriş Yapma | Öğrenci, daha önce oluşturduğu e-posta ve şifre bilgileriyle sisteme giriş yapabilmelidir. | POST | [Sorumlu Kişi] |
| G3 | Öğrenci Profilini Görüntüleme | Öğrenci, kendi adını, e-postasını ve temel bilgilerini içeren profil sayfasını görebilmelidir. | GET | [Sorumlu Kişi] |
| G4 | Tüm Öğrencileri Listeleme | Yetkili kullanıcı, sisteme kayıtlı tüm öğrencilerin ad ve temel bilgilerini liste halinde görebilmelidir. | GET | [Sorumlu Kişi] |
| G5 | Tüm Öğretmenleri Listeleme | Öğrenci, sistemdeki tüm öğretmenlerin adlarını ve verdikleri dersleri liste halinde görebilmelidir. | GET | [Sorumlu Kişi] |
| G6 | Randevu Alma | Öğrenci, seçtiği öğretmen ve uygun saat için bir randevu oluşturabilmelidir. | POST | [Sorumlu Kişi] |
| G7 | Randevuları Listeleme | Öğrenci, sadece kendisine ait geçmiş ve gelecek tüm randevuları liste halinde görebilmelidir. | GET | [Sorumlu Kişi] |
| G8 | Randevu Güncelleme | Öğrenci, daha önce aldığı bir randevunun saatini veya tarihini uygun başka bir zamanla değiştirebilmelidir. | PUT | [Sorumlu Kişi] |
| G9 | Randevu İptal Etme | Öğrenci, artık gitmek istemediği bir randevuyu sistemden iptal edebilmelidir. | DELETE | [Sorumlu Kişi] |
| G10 | Randevu Önerisi Alma | Öğrenci, haftalık boş zamanlarını ve ders adını seçerek kendisi için en uygun randevu saatleri için yapay zekâ tabanlı öneri alabilmelidir. | POST | [Sorumlu Kişi] |

# Gereksinim Dağılımları

1. [Grup Üyesi 1'in Gereksinimleri](Grup-Uyesi-1/Grup-Uyesi-1-Gereksinimler.md)
2. [Grup Üyesi 2'nin Gereksinimleri](Grup-Uyesi-2/Grup-Uyesi-2-Gereksinimler.md)
3. [Grup Üyesi 3'ün Gereksinimleri](Grup-Uyesi-3/Grup-Uyesi-3-Gereksinimler.md)
4. [Grup Üyesi 4'ün Gereksinimleri](Grup-Uyesi-4/Grup-Uyesi-4-Gereksinimler.md)
5. [Grup Üyesi 5'in Gereksinimleri](Grup-Uyesi-5/Grup-Uyesi-5-Gereksinimler.md)
6. [Grup Üyesi 6'nın Gereksinimleri](Grup-Uyesi-6/Grup-Uyesi-6-Gereksinimler.md)
