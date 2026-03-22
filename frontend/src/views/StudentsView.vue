<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '@/api/client'

interface Student {
  id: string
  name: string
  email: string
  createdAt: string
}

const students = ref<Student[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await api.get('/students')
    students.value = res.data.data || []
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div>
    <h1 class="page-title">Öğrenciler</h1>

    <div v-if="loading" class="card">Yükleniyor...</div>

    <div v-else-if="students.length === 0" class="card" style="text-align: center; color: #64748b;">
      Henüz kayıtlı öğrenci yok.
    </div>

    <div v-else class="students-list">
      <div v-for="student in students" :key="student.id" class="card student-card">
        <div class="student-info">
          <div class="student-avatar">{{ student.name.charAt(0) }}</div>
          <div>
            <h3>{{ student.name }}</h3>
            <p class="student-email">{{ student.email }}</p>
          </div>
        </div>
        <span class="student-date">
          Kayıt: {{ new Date(student.createdAt).toLocaleDateString('tr-TR') }}
        </span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.students-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.student-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.student-info {
  display: flex;
  align-items: center;
  gap: 14px;
}

.student-avatar {
  width: 42px;
  height: 42px;
  border-radius: 50%;
  background: #eff6ff;
  color: #2563eb;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 700;
  flex-shrink: 0;
}

.student-card h3 {
  font-size: 15px;
  margin-bottom: 2px;
}

.student-email {
  font-size: 13px;
  color: #64748b;
}

.student-date {
  font-size: 12px;
  color: #94a3b8;
}
</style>
