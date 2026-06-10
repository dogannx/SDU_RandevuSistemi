# Doğan Coşman'ın Mobil Backend Görevleri

**Mobil Backend Demo:** *(Demo hoca odasında yüz yüze sunulmuştur.)*

**API Adresi:** [sdu-randevusistemi.onrender.com](https://sdu-randevusistemi.onrender.com)

**Yapı:** Axios instance + `Constants.expoConfig.extra.apiBaseUrl` ile env'den base URL + request interceptor (token inject) + response interceptor (401 SecureStore cleanup).

## Görevler

| No | Görev | Durum |
|---|---|---|
| 1 | Kullanıcı kayıt/giriş API entegrasyonu (POST /auth/register, POST /auth/login) | [x] |
| 2 | Öğretmen listeleme endpoint entegrasyonu (GET /teachers) | [x] |
| 3 | Randevu oluşturma/güncelleme/silme entegrasyonu (POST/PUT/DELETE /appointments) | [x] |
| 4 | Randevu listeleme entegrasyonu (GET /appointments) | [x] |
| 5 | AI öneri endpoint entegrasyonu (POST /appointments/suggest) | [x] |
| 6 | Öğrenci profili/listesi entegrasyonu (GET /students/:id, GET /students) | [x] |
| 7 | Token persist (expo-secure-store) + auth interceptor + 401 cleanup | [x] |
