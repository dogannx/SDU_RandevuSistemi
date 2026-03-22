<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()
const router = useRouter()

const name = ref('')
const email = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)

async function handleRegister() {
  error.value = ''

  if (password.value.length < 8) {
    error.value = 'Şifre en az 8 karakter olmalıdır'
    return
  }

  loading.value = true
  try {
    await auth.register(name.value, email.value, password.value)
    router.push('/')
  } catch (e: any) {
    error.value = e.response?.data?.error || 'Kayıt başarısız'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="auth-page">
    <div class="auth-card">
      <h1>📚 Randevu Sistemi</h1>
      <h2>Kayıt Ol</h2>

      <div v-if="error" class="alert alert-error">{{ error }}</div>

      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label>Ad Soyad</label>
          <input v-model="name" type="text" placeholder="Doğan Coşman" required />
        </div>
        <div class="form-group">
          <label>E-posta</label>
          <input v-model="email" type="email" placeholder="ornek@email.com" required />
        </div>
        <div class="form-group">
          <label>Şifre (en az 8 karakter)</label>
          <input v-model="password" type="password" placeholder="********" required />
        </div>
        <button type="submit" class="btn btn-primary btn-block" :disabled="loading">
          {{ loading ? 'Kayıt yapılıyor...' : 'Kayıt Ol' }}
        </button>
      </form>

      <p class="auth-link">
        Zaten hesabın var mı? <RouterLink to="/login">Giriş Yap</RouterLink>
      </p>
    </div>
  </div>
</template>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.auth-card {
  background: #fff;
  padding: 40px;
  border-radius: 16px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15);
  width: 100%;
  max-width: 400px;
}

.auth-card h1 {
  text-align: center;
  font-size: 24px;
  margin-bottom: 4px;
}

.auth-card h2 {
  text-align: center;
  font-size: 18px;
  color: #64748b;
  margin-bottom: 24px;
  font-weight: 400;
}

.btn-block {
  width: 100%;
  padding: 12px;
  font-size: 16px;
}

.auth-link {
  text-align: center;
  margin-top: 16px;
  font-size: 14px;
  color: #64748b;
}

.auth-link a {
  color: #2563eb;
  text-decoration: none;
  font-weight: 500;
}
</style>
