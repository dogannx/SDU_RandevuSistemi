<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()
</script>

<template>
  <div>
    <h1 class="page-title">Profilim</h1>

    <div class="card profile-card">
      <div class="profile-avatar">{{ auth.student?.name?.charAt(0) || '?' }}</div>
      <div class="profile-info">
        <div class="profile-row">
          <span class="label">Ad Soyad</span>
          <span class="value">{{ auth.student?.name }}</span>
        </div>
        <div class="profile-row">
          <span class="label">E-posta</span>
          <span class="value">{{ auth.student?.email }}</span>
        </div>
        <div class="profile-row">
          <span class="label">Kayıt Tarihi</span>
          <span class="value">{{ new Date(auth.student?.createdAt || '').toLocaleDateString('tr-TR') }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-card {
  display: flex;
  align-items: flex-start;
  gap: 24px;
}

.profile-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32px;
  font-weight: 700;
  flex-shrink: 0;
}

.profile-info {
  flex: 1;
}

.profile-row {
  display: flex;
  padding: 12px 0;
  border-bottom: 1px solid #f1f5f9;
}

.profile-row:last-child {
  border-bottom: none;
}

.profile-row .label {
  width: 120px;
  font-size: 14px;
  color: #64748b;
  font-weight: 500;
}

.profile-row .value {
  font-size: 14px;
  color: #1e293b;
}
</style>
