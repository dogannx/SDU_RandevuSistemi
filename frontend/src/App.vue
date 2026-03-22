<script setup lang="ts">
import { RouterView, RouterLink, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()
const route = useRoute()
</script>

<template>
  <div id="app">
    <nav v-if="auth.isLoggedIn" class="navbar">
      <div class="nav-brand">
        <span>📚 Randevu Sistemi</span>
      </div>
      <div class="nav-links">
        <RouterLink to="/" :class="{ active: route.name === 'teachers' }">Öğretmenler</RouterLink>
        <RouterLink to="/appointments" :class="{ active: route.name === 'appointments' }">Randevularım</RouterLink>
        <RouterLink to="/suggest" :class="{ active: route.name === 'suggest' }">AI Öneri</RouterLink>
        <RouterLink to="/students" :class="{ active: route.name === 'students' }">Öğrenciler</RouterLink>
        <RouterLink to="/profile" :class="{ active: route.name === 'profile' }">Profilim</RouterLink>
      </div>
      <div class="nav-user">
        <span class="user-name">{{ auth.student?.name }}</span>
        <button class="btn-logout" @click="auth.logout()">Çıkış</button>
      </div>
    </nav>

    <main :class="{ container: auth.isLoggedIn }">
      <RouterView />
    </main>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background: #f5f7fa;
  color: #333;
}

.navbar {
  background: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  padding: 0 24px;
  height: 60px;
  gap: 24px;
}

.nav-brand span {
  font-weight: 700;
  font-size: 18px;
  color: #2563eb;
}

.nav-links {
  display: flex;
  gap: 8px;
  flex: 1;
}

.nav-links a {
  text-decoration: none;
  color: #64748b;
  padding: 8px 16px;
  border-radius: 6px;
  font-size: 14px;
  transition: all 0.2s;
}

.nav-links a:hover,
.nav-links a.active {
  background: #eff6ff;
  color: #2563eb;
}

.nav-user {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-name {
  font-size: 14px;
  color: #64748b;
}

.btn-logout {
  background: none;
  border: 1px solid #e2e8f0;
  padding: 6px 14px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 13px;
  color: #64748b;
  transition: all 0.2s;
}

.btn-logout:hover {
  background: #fee2e2;
  color: #dc2626;
  border-color: #fecaca;
}

.container {
  max-width: 960px;
  margin: 24px auto;
  padding: 0 24px;
}

.card {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  padding: 24px;
  margin-bottom: 16px;
}

.page-title {
  font-size: 24px;
  font-weight: 700;
  margin-bottom: 20px;
  color: #1e293b;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.2s;
}

.btn-primary {
  background: #2563eb;
  color: #fff;
}

.btn-primary:hover {
  background: #1d4ed8;
}

.btn-danger {
  background: #dc2626;
  color: #fff;
}

.btn-danger:hover {
  background: #b91c1c;
}

.btn-secondary {
  background: #f1f5f9;
  color: #475569;
}

.btn-secondary:hover {
  background: #e2e8f0;
}

.form-group {
  margin-bottom: 16px;
}

.form-group label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 6px;
  color: #374151;
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 14px;
  transition: border-color 0.2s;
}

.form-group input:focus,
.form-group select:focus {
  outline: none;
  border-color: #2563eb;
  box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.1);
}

.alert {
  padding: 12px 16px;
  border-radius: 8px;
  margin-bottom: 16px;
  font-size: 14px;
}

.alert-error {
  background: #fef2f2;
  color: #dc2626;
  border: 1px solid #fecaca;
}

.alert-success {
  background: #f0fdf4;
  color: #16a34a;
  border: 1px solid #bbf7d0;
}

.badge {
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.badge-upcoming {
  background: #dbeafe;
  color: #2563eb;
}

.badge-cancelled {
  background: #fee2e2;
  color: #dc2626;
}

.badge-past {
  background: #f1f5f9;
  color: #64748b;
}
</style>
