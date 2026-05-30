import { useRouter } from 'expo-router';
import { ScrollView, StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { useAuth } from '@/contexts/auth-context';

export default function HomeScreen() {
  const { student, logout } = useAuth();
  const router = useRouter();

  return (
    <SafeAreaView style={styles.safe}>
      <ScrollView contentContainerStyle={styles.container}>
        <Text style={styles.title}>Profil</Text>

        <View style={styles.card}>
          <Text style={styles.label}>Ad Soyad</Text>
          <Text style={styles.value}>{student?.name ?? '—'}</Text>

          <View style={styles.divider} />

          <Text style={styles.label}>E-posta</Text>
          <Text style={styles.value}>{student?.email ?? '—'}</Text>

          {student?.createdAt && (
            <>
              <View style={styles.divider} />
              <Text style={styles.label}>Kayıt Tarihi</Text>
              <Text style={styles.value}>
                {new Date(student.createdAt).toLocaleDateString('tr-TR')}
              </Text>
            </>
          )}
        </View>

        <TouchableOpacity
          style={styles.linkButton}
          onPress={() => router.push('/students')}>
          <Text style={styles.linkButtonText}>Tüm Öğrenciler</Text>
          <Text style={styles.chev}>›</Text>
        </TouchableOpacity>

        <TouchableOpacity style={styles.logoutButton} onPress={logout}>
          <Text style={styles.logoutText}>Çıkış Yap</Text>
        </TouchableOpacity>
      </ScrollView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  safe: { flex: 1, backgroundColor: '#fff' },
  container: { padding: 24, paddingTop: 48 },
  title: { fontSize: 28, fontWeight: '700', color: '#111827', marginBottom: 24 },
  card: {
    backgroundColor: '#f9fafb',
    borderRadius: 12,
    padding: 20,
    borderWidth: 1,
    borderColor: '#e5e7eb',
    marginBottom: 24,
  },
  label: { fontSize: 12, color: '#6b7280', fontWeight: '600', textTransform: 'uppercase' },
  value: { fontSize: 16, color: '#111827', marginTop: 4 },
  divider: { height: 1, backgroundColor: '#e5e7eb', marginVertical: 14 },
  linkButton: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    backgroundColor: '#eff6ff',
    borderRadius: 8,
    padding: 16,
    marginBottom: 24,
  },
  linkButtonText: { color: '#2563eb', fontWeight: '600', fontSize: 16 },
  chev: { color: '#2563eb', fontSize: 22, fontWeight: '600' },
  logoutButton: {
    borderWidth: 1,
    borderColor: '#dc2626',
    borderRadius: 8,
    paddingVertical: 14,
    alignItems: 'center',
  },
  logoutText: { color: '#dc2626', fontWeight: '600', fontSize: 16 },
});
