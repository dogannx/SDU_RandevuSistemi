import { StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { useAuth } from '@/contexts/auth-context';

export default function HomeScreen() {
  const { student, logout } = useAuth();

  return (
    <SafeAreaView style={styles.safe}>
      <View style={styles.container}>
        <Text style={styles.title}>Birebir Ders Randevu Sistemi</Text>
        <Text style={styles.subtitle}>
          Hoş geldin{student?.name ? `, ${student.name}` : ''}!
        </Text>

        <View style={styles.placeholder}>
          <Text style={styles.placeholderText}>
            Öğretmenler, randevular ve AI öneri sayfaları sonraki fazlarda gelecek.
          </Text>
        </View>

        <TouchableOpacity style={styles.logoutButton} onPress={logout}>
          <Text style={styles.logoutText}>Çıkış Yap</Text>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  safe: { flex: 1, backgroundColor: '#fff' },
  container: { flex: 1, padding: 24, justifyContent: 'center' },
  title: { fontSize: 24, fontWeight: '700', color: '#111827' },
  subtitle: { fontSize: 16, color: '#6b7280', marginTop: 8, marginBottom: 32 },
  placeholder: {
    backgroundColor: '#f3f4f6',
    borderRadius: 12,
    padding: 20,
    marginBottom: 32,
  },
  placeholderText: { color: '#374151', lineHeight: 22 },
  logoutButton: {
    borderWidth: 1,
    borderColor: '#dc2626',
    borderRadius: 8,
    paddingVertical: 14,
    alignItems: 'center',
  },
  logoutText: { color: '#dc2626', fontWeight: '600', fontSize: 16 },
});
