<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api/client'

interface Teacher {
  id: string
  name: string
  subjects: string[]
  bio?: string
}

const router = useRouter()
const teachers = ref<Teacher[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await api.get('/teachers')
    teachers.value = res.data.data
  } finally {
    loading.value = false
  }
})

function bookAppointment(teacher: Teacher) {
  router.push({
    path: '/appointments',
    query: { teacherId: teacher.id, teacherName: teacher.name },
  })
}
</script>

<template>
  <div>
    <h1 class="page-title">Öğretmenler</h1>

    <div v-if="loading" class="card">Yükleniyor...</div>

    <div v-else class="teachers-grid">
      <div v-for="teacher in teachers" :key="teacher.id" class="card teacher-card">
        <div class="teacher-info">
          <div class="teacher-avatar">{{ teacher.name.charAt(0) }}</div>
          <div>
            <h3>{{ teacher.name }}</h3>
            <p class="teacher-bio">{{ teacher.bio }}</p>
            <div class="subjects">
              <span v-for="subject in teacher.subjects" :key="subject" class="subject-tag">
                {{ subject }}
              </span>
            </div>
          </div>
        </div>
        <button class="btn btn-primary" @click="bookAppointment(teacher)">Randevu Al</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.teachers-grid {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.teacher-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.teacher-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.teacher-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: #eff6ff;
  color: #2563eb;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 700;
  flex-shrink: 0;
}

.teacher-card h3 {
  font-size: 16px;
  margin-bottom: 4px;
}

.teacher-bio {
  font-size: 13px;
  color: #64748b;
  margin-bottom: 6px;
}

.subjects {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
}

.subject-tag {
  background: #f0fdf4;
  color: #16a34a;
  padding: 2px 10px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}
</style>
