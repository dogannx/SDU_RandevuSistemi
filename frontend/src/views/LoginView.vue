<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()
const router = useRouter()

const email = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)

async function handleLogin() {
  error.value = ''
  loading.value = true
  try {
    await auth.login(email.value, password.value)
    router.push('/')
  } catch (e: any) {
    error.value = e.response?.data?.error || 'Giriş başarısız'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="auth-page">
    <div class="auth-card">
      <h1>📚 Randevu Sistemi</h1>
      <h2>Giriş Yap</h2>

      <div v-if="error" class="alert alert-error">{{ error }}</div>

      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label>E-posta</label>
          <input v-model="email" type="email" placeholder="ornek@email.com" required />
        </div>
        <div class="form-group">
          <label>Şifre</label>
          <input v-model="password" type="password" placeholder="********" required />
        </div>
        <button type="submit" class="btn btn-primary btn-block" :disabled="loading">
          {{ loading ? 'Giriş yapılıyor...' : 'Giriş Yap' }}
        </button>
      </form>

      <p class="auth-link">
        Hesabın yok mu? <RouterLink to="/register">Kayıt Ol</RouterLink>
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
