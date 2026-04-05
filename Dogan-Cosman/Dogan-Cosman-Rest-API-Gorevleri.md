# Doğan Coşman'ın REST API Metotları

**API Test Videosu:** [API Test Videosu](https://youtu.be/VQznhG9tUfI)

## G1 — Kayıt Olma
- **Endpoint:** `POST /api/v1/auth/register`
- **Request Body:**
  ```json
  {
    "name": "Doğan Coşman",
    "email": "dogan@example.com",
    "password": "Guvenli123!"
  }
  ```
- **Response:** `201 Created` — Öğrenci hesabı başarıyla oluşturuldu ✅

## G2 — Giriş Yapma
- **Endpoint:** `POST /api/v1/auth/login`
- **Request Body:**
  ```json
  {
    "email": "dogan@example.com",
    "password": "Guvenli123!"
  }
  ```
- **Response:** `200 OK` — Giriş başarılı, JWT token döndürülür ✅

## G3 — Öğrenci Profilini Görüntüleme
- **Endpoint:** `GET /api/v1/students/{studentId}`
- **Path Parameters:**
  - `studentId` (string, required) — Öğrenci ID'si
- **Authentication:** Bearer Token gerekli
- **Response:** `200 OK` — Öğrenci profil bilgileri döndürülür ✅

## G4 — Tüm Öğrencileri Listeleme
- **Endpoint:** `GET /api/v1/students`
- **Authentication:** Bearer Token gerekli
- **Response:** `200 OK` — Tüm öğrencilerin listesi döndürülür ✅

## G5 — Tüm Öğretmenleri Listeleme
- **Endpoint:** `GET /api/v1/teachers`
- **Response:** `200 OK` — Tüm öğretmenlerin ad ve ders bilgileri döndürülür ✅

## G6 — Randevu Alma
- **Endpoint:** `POST /api/v1/appointments`
- **Authentication:** Bearer Token gerekli
- **Request Body:**
  ```json
  {
    "teacherId": "uuid",
    "date": "2026-04-01",
    "time": "14:00"
  }
  ```
- **Response:** `201 Created` — Randevu başarıyla oluşturuldu ✅

## G7 — Randevuları Listeleme
- **Endpoint:** `GET /api/v1/appointments`
- **Authentication:** Bearer Token gerekli
- **Response:** `200 OK` — Öğrenciye ait tüm randevular döndürülür ✅

## G8 — Randevu Güncelleme
- **Endpoint:** `PUT /api/v1/appointments/{appointmentId}`
- **Path Parameters:**
  - `appointmentId` (string, required) — Randevu ID'si
- **Authentication:** Bearer Token gerekli
- **Request Body:**
  ```json
  {
    "date": "2026-04-05",
    "time": "10:00"
  }
  ```
- **Response:** `200 OK` — Randevu başarıyla güncellendi ✅

## G9 — Randevu İptal Etme
- **Endpoint:** `DELETE /api/v1/appointments/{appointmentId}`
- **Path Parameters:**
  - `appointmentId` (string, required) — Randevu ID'si
- **Authentication:** Bearer Token gerekli
- **Response:** `204 No Content` — Randevu başarıyla iptal edildi ✅

## G10 — Randevu Önerisi Alma
- **Endpoint:** `POST /api/v1/appointments/suggest`
- **Authentication:** Bearer Token gerekli
- **Request Body:**
  ```json
  {
    "lessonName": "Matematik",
    "availableSlots": ["Pazartesi 09:00-12:00", "Çarşamba 14:00-17:00"]
  }
  ```
- **Response:** `200 OK` — Yapay zekâ tabanlı önerilen randevu saatleri listesi döndürülür ✅
