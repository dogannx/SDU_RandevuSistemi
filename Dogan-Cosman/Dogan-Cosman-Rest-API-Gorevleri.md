# Doğan Coşman'ın REST API Metotları

**API Test Videosu:** [Link buraya eklenecek](https://example.com)

## G1 — Kayıt Olma
- **Endpoint:** `POST /auth/register`
- **Request Body:**
  ```json
  {
    "name": "Doğan Coşman",
    "email": "dogan@example.com",
    "password": "Guvenli123!"
  }
  ```
- **Response:** `201 Created` — Öğrenci hesabı başarıyla oluşturuldu

## G2 — Giriş Yapma
- **Endpoint:** `POST /auth/login`
- **Request Body:**
  ```json
  {
    "email": "dogan@example.com",
    "password": "Guvenli123!"
  }
  ```
- **Response:** `200 OK` — Giriş başarılı, token döndürülür

## G3 — Öğrenci Profilini Görüntüleme
- **Endpoint:** `GET /students/{studentId}`
- **Path Parameters:**
  - `studentId` (string, required) — Öğrenci ID'si
- **Authentication:** Bearer Token gerekli
- **Response:** `200 OK` — Öğrenci profil bilgileri döndürülür

## G4 — Tüm Öğrencileri Listeleme
- **Endpoint:** `GET /students`
- **Authentication:** Bearer Token gerekli (yetkili kullanıcı)
- **Response:** `200 OK` — Tüm öğrencilerin listesi döndürülür

## G5 — Tüm Öğretmenleri Listeleme
- **Endpoint:** `GET /teachers`
- **Response:** `200 OK` — Tüm öğretmenlerin ad ve ders bilgileri döndürülür

## G6 — Randevu Alma
- **Endpoint:** `POST /appointments`
- **Authentication:** Bearer Token gerekli
- **Request Body:**
  ```json
  {
    "teacherId": "t123",
    "date": "2025-03-15",
    "time": "14:00"
  }
  ```
- **Response:** `201 Created` — Randevu başarıyla oluşturuldu

## G7 — Randevuları Listeleme
- **Endpoint:** `GET /appointments/me`
- **Authentication:** Bearer Token gerekli
- **Response:** `200 OK` — Öğrenciye ait tüm randevular döndürülür

## G8 — Randevu Güncelleme
- **Endpoint:** `PUT /appointments/{appointmentId}`
- **Path Parameters:**
  - `appointmentId` (string, required) — Randevu ID'si
- **Authentication:** Bearer Token gerekli
- **Request Body:**
  ```json
  {
    "date": "2025-03-20",
    "time": "10:00"
  }
  ```
- **Response:** `200 OK` — Randevu başarıyla güncellendi

## G9 — Randevu İptal Etme
- **Endpoint:** `DELETE /appointments/{appointmentId}`
- **Path Parameters:**
  - `appointmentId` (string, required) — Randevu ID'si
- **Authentication:** Bearer Token gerekli
- **Response:** `204 No Content` — Randevu başarıyla iptal edildi

## G10 — Randevu Önerisi Alma
- **Endpoint:** `POST /appointments/suggest`
- **Authentication:** Bearer Token gerekli
- **Request Body:**
  ```json
  {
    "lessonName": "Matematik",
    "availableSlots": ["Pazartesi 09:00-12:00", "Çarşamba 14:00-17:00"]
  }
  ```
- **Response:** `200 OK` — Yapay zekâ tabanlı önerilen randevu saatleri listesi döndürülür
