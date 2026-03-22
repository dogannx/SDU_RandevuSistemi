# API Tasarımı - OpenAPI Specification

**OpenAPI Spesifikasyon Dosyası:** [randevu-api.yaml](randevu-api.yaml)

Bu doküman, Birebir Ders Randevu Sistemi için OpenAPI Specification (OAS) 3.0 standardına göre hazırlanmış REST API tasarımını içermektedir.

## Endpoint Özeti

| No | Gereksinim | HTTP Metodu | Endpoint |
|---|---|---|---|
| G1 | Kayıt Olma | `POST` | `/api/v1/auth/register` |
| G2 | Giriş Yapma | `POST` | `/api/v1/auth/login` |
| G3 | Öğrenci Profilini Görüntüleme | `GET` | `/api/v1/students/{studentId}` |
| G4 | Tüm Öğrencileri Listeleme | `GET` | `/api/v1/students` |
| G5 | Tüm Öğretmenleri Listeleme | `GET` | `/api/v1/teachers` |
| G6 | Randevu Alma | `POST` | `/api/v1/appointments` |
| G7 | Randevuları Listeleme | `GET` | `/api/v1/appointments` |
| G8 | Randevu Güncelleme | `PUT` | `/api/v1/appointments/{appointmentId}` |
| G9 | Randevu İptal Etme | `DELETE` | `/api/v1/appointments/{appointmentId}` |
| G10 | Randevu Önerisi Alma | `POST` | `/api/v1/appointments/suggest` |

## OpenAPI Specification (YAML)

```yaml
openapi: 3.0.3
info:
  title: Birebir Ders Randevu Sistemi API
  description: |
    Birebir Ders Randevu Sistemi için RESTful API.

    ## Özellikler
    - Öğrenci kaydı ve girişi
    - Öğrenci ve öğretmen listeleme
    - Randevu oluşturma, listeleme, güncelleme ve iptal etme
    - Yapay zekâ tabanlı randevu önerisi
    - JWT tabanlı kimlik doğrulama
  version: 1.0.0
  contact:
    name: Doğan Coşman
    email: dogan@example.com

servers:
  - url: https://randevu-api-mic1.onrender.com/api/v1
    description: Production sunucusu
  - url: http://localhost:8080/api/v1
    description: Yerel geliştirme sunucusu

tags:
  - name: auth
    description: Kimlik doğrulama işlemleri
  - name: students
    description: Öğrenci yönetimi işlemleri
  - name: teachers
    description: Öğretmen listeleme işlemleri
  - name: appointments
    description: Randevu yönetimi işlemleri

paths:

  /auth/register:
    post:
      tags:
        - auth
      summary: G1 — Kayıt Olma
      operationId: registerStudent
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/RegisterRequest'
      responses:
        '201':
          description: Öğrenci hesabı başarıyla oluşturuldu
        '400':
          $ref: '#/components/responses/BadRequest'
        '409':
          description: Bu e-posta adresi zaten kullanımda

  /auth/login:
    post:
      tags:
        - auth
      summary: G2 — Giriş Yapma
      operationId: loginStudent
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/LoginRequest'
      responses:
        '200':
          description: Giriş başarılı — JWT token döndürülür
        '401':
          $ref: '#/components/responses/Unauthorized'

  /students:
    get:
      tags:
        - students
      summary: G4 — Tüm Öğrencileri Listeleme
      operationId: listStudents
      security:
        - bearerAuth: []
      responses:
        '200':
          description: Öğrenci listesi başarıyla döndürüldü
        '401':
          $ref: '#/components/responses/Unauthorized'

  /students/{studentId}:
    get:
      tags:
        - students
      summary: G3 — Öğrenci Profilini Görüntüleme
      operationId: getStudentById
      security:
        - bearerAuth: []
      parameters:
        - $ref: '#/components/parameters/StudentIdParam'
      responses:
        '200':
          description: Öğrenci profil bilgileri döndürüldü
        '404':
          $ref: '#/components/responses/NotFound'

  /teachers:
    get:
      tags:
        - teachers
      summary: G5 — Tüm Öğretmenleri Listeleme
      operationId: listTeachers
      responses:
        '200':
          description: Öğretmen listesi başarıyla döndürüldü

  /appointments:
    get:
      tags:
        - appointments
      summary: G7 — Randevuları Listeleme
      operationId: listMyAppointments
      security:
        - bearerAuth: []
      responses:
        '200':
          description: Randevu listesi başarıyla döndürüldü

    post:
      tags:
        - appointments
      summary: G6 — Randevu Alma
      operationId: createAppointment
      security:
        - bearerAuth: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/AppointmentCreate'
      responses:
        '201':
          description: Randevu başarıyla oluşturuldu

  /appointments/{appointmentId}:
    put:
      tags:
        - appointments
      summary: G8 — Randevu Güncelleme
      operationId: updateAppointment
      security:
        - bearerAuth: []
      parameters:
        - $ref: '#/components/parameters/AppointmentIdParam'
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/AppointmentUpdate'
      responses:
        '200':
          description: Randevu başarıyla güncellendi

    delete:
      tags:
        - appointments
      summary: G9 — Randevu İptal Etme
      operationId: cancelAppointment
      security:
        - bearerAuth: []
      parameters:
        - $ref: '#/components/parameters/AppointmentIdParam'
      responses:
        '204':
          description: Randevu başarıyla iptal edildi

  /appointments/suggest:
    post:
      tags:
        - appointments
      summary: G10 — Randevu Önerisi Alma
      operationId: suggestAppointment
      security:
        - bearerAuth: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/SuggestRequest'
      responses:
        '200':
          description: Yapay zekâ tabanlı öneriler döndürüldü

components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT

  parameters:
    StudentIdParam:
      name: studentId
      in: path
      required: true
      schema:
        type: string
    AppointmentIdParam:
      name: appointmentId
      in: path
      required: true
      schema:
        type: string

  schemas:
    RegisterRequest:
      type: object
      required: [name, email, password]
      properties:
        name:
          type: string
          example: Doğan Coşman
        email:
          type: string
          format: email
          example: dogan@example.com
        password:
          type: string
          minLength: 8
          example: Guvenli123!

    LoginRequest:
      type: object
      required: [email, password]
      properties:
        email:
          type: string
          format: email
        password:
          type: string

    AppointmentCreate:
      type: object
      required: [teacherId, date, time]
      properties:
        teacherId:
          type: string
          example: t123
        date:
          type: string
          format: date
          example: "2026-03-15"
        time:
          type: string
          example: "14:00"

    AppointmentUpdate:
      type: object
      properties:
        date:
          type: string
          format: date
          example: "2026-03-20"
        time:
          type: string
          example: "10:00"

    SuggestRequest:
      type: object
      required: [lessonName, availableSlots]
      properties:
        lessonName:
          type: string
          example: Matematik
        availableSlots:
          type: array
          items:
            type: string
          example:
            - Pazartesi 09:00-12:00
            - Çarşamba 14:00-17:00

    Error:
      type: object
      required: [code, message]
      properties:
        code:
          type: string
          example: NOT_FOUND
        message:
          type: string
          example: Kaynak bulunamadı

  responses:
    BadRequest:
      description: Geçersiz istek
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Error'
    Unauthorized:
      description: Yetkisiz erişim
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Error'
    NotFound:
      description: Kaynak bulunamadı
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Error'
```