<script setup lang="ts">
import { ref } from 'vue'
import api from '@/api/client'

interface Suggestion {
  teacherId: string
  teacherName: string
  date: string
  time: string
  score: number
}

const lessonName = ref('')
const slots = ref([{ day: 'Pazartesi', start: '09:00', end: '12:00' }])
const suggestions = ref<Suggestion[]>([])
const loading = ref(false)
const error = ref('')
const searched = ref(false)

const days = ['Pazartesi', 'Salı', 'Çarşamba', 'Perşembe', 'Cuma']
const hours = ['08:00', '09:00', '10:00', '11:00', '12:00', '13:00', '14:00', '15:00', '16:00', '17:00', '18:00']

function addSlot() {
  slots.value.push({ day: 'Pazartesi', start: '09:00', end: '12:00' })
}

function removeSlot(index: number) {
  slots.value.splice(index, 1)
}

async function getSuggestions() {
  error.value = ''
  loading.value = true
  searched.value = true
  suggestions.value = []

  const availableSlots = slots.value.map((s) => `${s.day} ${s.start}-${s.end}`)

  try {
    const res = await api.post('/appointments/suggest', {
      lessonName: lessonName.value,
      availableSlots,
    })
    suggestions.value = res.data.data.suggestions || []
  } catch (e: any) {
    error.value = e.response?.data?.error || 'Öneri alınamadı'
  } finally {
    loading.value = false
  }
}

async function bookSuggestion(s: Suggestion) {
  try {
    await api.post('/appointments', {
      teacherId: s.teacherId,
      date: s.date,
      time: s.time,
    })
    alert(`Randevu oluşturuldu: ${s.teacherName} - ${s.date} ${s.time}`)
    // Öneriyi listeden kaldır
    suggestions.value = suggestions.value.filter((x) => x !== s)
  } catch (e: any) {
    alert(e.response?.data?.error || 'Randevu oluşturulamadı')
  }
}
</script>

<template>
  <div>
    <h1 class="page-title">AI Randevu Önerisi</h1>

    <div class="card">
      <form @submit.prevent="getSuggestions">
        <div class="form-group">
          <label>Ders Adı</label>
          <input v-model="lessonName" type="text" placeholder="Matematik" required />
        </div>

        <div class="slots-section">
          <label>Müsait Zaman Dilimlerin</label>
          <div v-for="(slot, index) in slots" :key="index" class="slot-row">
            <select v-model="slot.day">
              <option v-for="d in days" :key="d" :value="d">{{ d }}</option>
            </select>
            <select v-model="slot.start">
              <option v-for="h in hours" :key="h" :value="h">{{ h }}</option>
            </select>
            <span>—</span>
            <select v-model="slot.end">
              <option v-for="h in hours" :key="h" :value="h">{{ h }}</option>
            </select>
            <button v-if="slots.length > 1" type="button" class="btn-remove" @click="removeSlot(index)">✕</button>
          </div>
          <button type="button" class="btn btn-secondary" style="margin-top: 8px;" @click="addSlot">
            + Zaman Dilimi Ekle
          </button>
        </div>

        <button type="submit" class="btn btn-primary" :disabled="loading" style="margin-top: 16px;">
          {{ loading ? 'Öneriler alınıyor...' : 'Öneri Al' }}
        </button>
      </form>
    </div>

    <div v-if="error" class="alert alert-error">{{ error }}</div>

    <!-- Öneriler -->
    <div v-if="searched && !loading">
      <h2 style="font-size: 18px; margin: 20px 0 12px;">
        {{ suggestions.length > 0 ? 'Önerilen Randevular' : 'Uygun randevu bulunamadı' }}
      </h2>

      <div v-for="s in suggestions" :key="`${s.teacherId}-${s.date}-${s.time}`" class="card suggestion-card">
        <div class="suggestion-info">
          <div class="suggestion-teacher">{{ s.teacherName }}</div>
          <div class="suggestion-datetime">📅 {{ s.date }} &nbsp; 🕐 {{ s.time }}</div>
          <div class="suggestion-score">
            Uyum: <strong>{{ Math.round(s.score * 100) }}%</strong>
          </div>
        </div>
        <button class="btn btn-primary" @click="bookSuggestion(s)">Bu Randevuyu Al</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.slots-section {
  margin-bottom: 8px;
}

.slots-section > label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 8px;
  color: #374151;
}

.slot-row {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.slot-row select {
  padding: 8px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 14px;
}

.btn-remove {
  background: none;
  border: none;
  color: #dc2626;
  cursor: pointer;
  font-size: 16px;
  padding: 4px 8px;
}

.suggestion-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 8px;
}

.suggestion-teacher {
  font-weight: 600;
  margin-bottom: 4px;
}

.suggestion-datetime {
  font-size: 13px;
  color: #64748b;
  margin-bottom: 4px;
}

.suggestion-score {
  font-size: 13px;
  color: #16a34a;
}
</style>
