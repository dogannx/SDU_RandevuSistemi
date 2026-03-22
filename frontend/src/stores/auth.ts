import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '@/api/client'
import router from '@/router'

interface Student {
  id: string
  name: string
  email: string
  createdAt: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || '')
  const student = ref<Student | null>(
    JSON.parse(localStorage.getItem('student') || 'null'),
  )

  const isLoggedIn = computed(() => !!token.value)

  async function register(name: string, email: string, password: string) {
    const res = await api.post('/auth/register', { name, email, password })
    setAuth(res.data.data)
  }

  async function login(email: string, password: string) {
    const res = await api.post('/auth/login', { email, password })
    setAuth(res.data.data)
  }

  function setAuth(data: { token: string; student: Student }) {
    token.value = data.token
    student.value = data.student
    localStorage.setItem('token', data.token)
    localStorage.setItem('student', JSON.stringify(data.student))
  }

  function logout() {
    token.value = ''
    student.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('student')
    router.push('/login')
  }

  return { token, student, isLoggedIn, register, login, logout }
})
