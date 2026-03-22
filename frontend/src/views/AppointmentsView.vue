<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import api from '@/api/client'

interface Appointment {
  id: string
  teacherId: string
  date: string
  time: string
  status: string
  teacher?: { id: string; name: string; subjects: string[] }
}

interface Teacher {
  id: string
  name: string
  subjects: string[]
}

const route = useRoute()
const appointments = ref<Appointment[]>([])
const teachers = ref<Teacher[]>([])
const loading = ref(true)
const error = ref('')
const success = ref('')

// Yeni randevu form
const showForm = ref(false)
const newTeacherId = ref('')
const newDate = ref('')
const newTime = ref('')

// Düzenleme
const editingId = ref('')
const editDate = ref('')
const editTime = ref('')

onMounted(async () => {
  await loadData()

  // Öğretmen sayfasından yönlendirme
  if (route.query.teacherId) {
    newTeacherId.value = route.query.teacherId as string
    showForm.value = true
  }
})

async function loadData() {
  loading.value = true
  try {
    const [apptRes, teacherRes] = await Promise.all([
      api.get('/appointments'),
      api.get('/teachers'),
    ])
    appointments.value = apptRes.data.data || []
    teachers.value = teacherRes.data.data || []
  } finally {
    loading.value = false
  }
}

async function createAppointment() {
  error.value = ''
  success.value = ''
  try {
    await api.post('/appointments', {
      teacherId: newTeacherId.value,
      date: newDate.value,
      time: newTime.value,
    })
    success.value = 'Randevu oluşturuldu!'
    showForm.value = false
    newTeacherId.value = ''
    newDate.value = ''
    newTime.value = ''
    await loadData()
  } catch (e: any) {
    error.value = e.response?.data?.error || 'Randevu oluşturulamadı'
  }
}

function startEdit(appt: Appointment) {
  editingId.value = appt.id
  editDate.value = appt.date
  editTime.value = appt.time
}

async function updateAppointment() {
  error.value = ''
  success.value = ''
  try {
    await api.put(`/appointments/${editingId.value}`, {
      date: editDate.value,
      time: editTime.value,
    })
    success.value = 'Randevu güncellendi!'
    editingId.value = ''
    await loadData()
  } catch (e: any) {
    error.value = e.response?.data?.error || 'Randevu güncellenemedi'
  }
}

async function cancelAppointment(id: string) {
  if (!confirm('Bu randevuyu iptal etmek istediğinize emin misiniz?')) return
  error.value = ''
  success.value = ''
  try {
    await api.delete(`/appointments/${id}`)
    success.value = 'Randevu iptal edildi'
    await loadData()
  } catch (e: any) {
    error.value = e.response?.data?.error || 'Randevu iptal edilemedi'
  }
}

const timeSlots = [
  '08:00', '09:00', '10:00', '11:00', '12:00',
  '13:00', '14:00', '15:00', '16:00', '17:00',
]
</script>

<template>
  <div>
    <div class="page-header">
      <h1 class="page-title">Randevularım</h1>
      <button class="btn btn-primary" @click="showForm = !showForm">
        {{ showForm ? 'İptal' : '+ Yeni Randevu' }}
      </button>
    </div>

    <div v-if="error" class="alert alert-error">{{ error }}</div>
    <div v-if="success" class="alert alert-success">{{ success }}</div>

    <!-- Yeni Randevu Formu -->
    <div v-if="showForm" class="card">
      <h3 style="margin-bottom: 16px;">Yeni Randevu Oluştur</h3>
      <form @submit.prevent="createAppointment">
        <div class="form-row">
          <div class="form-group">
            <label>Öğretmen</label>
            <select v-model="newTeacherId" required>
              <option value="">Seçiniz</option>
              <option v-for="t in teachers" :key="t.id" :value="t.id">
                {{ t.name }} ({{ t.subjects.join(', ') }})
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>Tarih</label>
            <input v-model="newDate" type="date" required />
          </div>
          <div class="form-group">
            <label>Saat</label>
            <select v-model="newTime" required>
              <option value="">Seçiniz</option>
              <option v-for="t in timeSlots" :key="t" :value="t">{{ t }}</option>
            </select>
          </div>
        </div>
        <button type="submit" class="btn btn-primary">Randevu Oluştur</button>
      </form>
    </div>

    <!-- Randevu Listesi -->
    <div v-if="loading" class="card">Yükleniyor...</div>

    <div v-else-if="appointments.length === 0" class="card empty-state">
      <p>Henüz randevunuz yok.</p>
    </div>

    <div v-else class="appointments-list">
      <div v-for="appt in appointments" :key="appt.id" class="card appt-card">
        <template v-if="editingId === appt.id">
          <form class="edit-form" @submit.prevent="updateAppointment">
            <div class="form-row">
              <div class="form-group">
                <label>Tarih</label>
                <input v-model="editDate" type="date" required />
              </div>
              <div class="form-group">
                <label>Saat</label>
                <select v-model="editTime" required>
                  <option v-for="t in timeSlots" :key="t" :value="t">{{ t }}</option>
                </select>
              </div>
            </div>
            <div class="btn-group">
              <button type="submit" class="btn btn-primary">Kaydet</button>
              <button type="button" class="btn btn-secondary" @click="editingId = ''">İptal</button>
            </div>
          </form>
        </template>

        <template v-else>
          <div class="appt-info">
            <div class="appt-teacher">{{ appt.teacher?.name || 'Öğretmen' }}</div>
            <div class="appt-datetime">📅 {{ appt.date }} &nbsp; 🕐 {{ appt.time }}</div>
          </div>
          <div class="appt-actions">
            <span :class="['badge', `badge-${appt.status}`]">
              {{ appt.status === 'upcoming' ? 'Aktif' : appt.status === 'cancelled' ? 'İptal' : 'Geçmiş' }}
            </span>
            <template v-if="appt.status === 'upcoming'">
              <button class="btn btn-secondary btn-sm" @click="startEdit(appt)">Düzenle</button>
              <button class="btn btn-danger btn-sm" @click="cancelAppointment(appt.id)">İptal Et</button>
            </template>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<style scoped>
.page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.page-header .page-title {
  margin-bottom: 0;
}

.form-row {
  display: flex;
  gap: 12px;
  margin-bottom: 12px;
}

.form-row .form-group {
  flex: 1;
  margin-bottom: 0;
}

.appointments-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.appt-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.appt-teacher {
  font-weight: 600;
  font-size: 15px;
  margin-bottom: 4px;
}

.appt-datetime {
  font-size: 13px;
  color: #64748b;
}

.appt-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.btn-sm {
  padding: 6px 12px;
  font-size: 13px;
}

.btn-group {
  display: flex;
  gap: 8px;
}

.edit-form {
  width: 100%;
}

.empty-state {
  text-align: center;
  color: #64748b;
  padding: 40px;
}
</style>
