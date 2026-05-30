import { useCallback, useEffect, useState } from 'react';
import {
  ActivityIndicator,
  Alert,
  FlatList,
  Modal,
  RefreshControl,
  StyleSheet,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { api } from '@/lib/api';

type Teacher = {
  id: string;
  name: string;
  subjects: string[];
  bio?: string | null;
};

export default function TeachersScreen() {
  const [teachers, setTeachers] = useState<Teacher[]>([]);
  const [loading, setLoading] = useState(true);
  const [refreshing, setRefreshing] = useState(false);
  const [selected, setSelected] = useState<Teacher | null>(null);
  const [date, setDate] = useState('');
  const [time, setTime] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const load = useCallback(async () => {
    try {
      const { data } = await api.get('/teachers');
      setTeachers(data.data ?? []);
    } catch (err: any) {
      Alert.alert('Hata', err?.response?.data?.error ?? 'Öğretmenler yüklenemedi');
    }
  }, []);

  useEffect(() => {
    (async () => {
      await load();
      setLoading(false);
    })();
  }, [load]);

  async function onRefresh() {
    setRefreshing(true);
    await load();
    setRefreshing(false);
  }

  function openBooking(t: Teacher) {
    setSelected(t);
    setDate('');
    setTime('');
  }

  function closeBooking() {
    setSelected(null);
  }

  async function submitBooking() {
    if (!selected) return;
    if (!date.trim() || !time.trim()) {
      Alert.alert('Eksik bilgi', 'Tarih ve saat gerekli.');
      return;
    }
    try {
      setSubmitting(true);
      await api.post('/appointments', {
        teacherId: selected.id,
        date: date.trim(),
        time: time.trim(),
      });
      closeBooking();
      Alert.alert('Başarılı', 'Randevu oluşturuldu.');
    } catch (err: any) {
      Alert.alert('Hata', err?.response?.data?.error ?? 'Randevu oluşturulamadı');
    } finally {
      setSubmitting(false);
    }
  }

  if (loading) {
    return (
      <SafeAreaView style={styles.safe}>
        <View style={styles.center}>
          <ActivityIndicator size="large" color="#2563eb" />
        </View>
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView style={styles.safe}>
      <FlatList
        data={teachers}
        keyExtractor={(t) => t.id}
        contentContainerStyle={styles.list}
        ListHeaderComponent={<Text style={styles.title}>Öğretmenler</Text>}
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={onRefresh} />}
        renderItem={({ item }) => (
          <View style={styles.card}>
            <Text style={styles.name}>{item.name}</Text>
            <Text style={styles.subjects}>{item.subjects.join(', ')}</Text>
            {item.bio ? <Text style={styles.bio}>{item.bio}</Text> : null}
            <TouchableOpacity style={styles.bookButton} onPress={() => openBooking(item)}>
              <Text style={styles.bookButtonText}>Randevu Al</Text>
            </TouchableOpacity>
          </View>
        )}
      />

      <Modal visible={!!selected} animationType="slide" transparent onRequestClose={closeBooking}>
        <View style={styles.modalOverlay}>
          <View style={styles.modalCard}>
            <Text style={styles.modalTitle}>Randevu Al</Text>
            <Text style={styles.modalSubtitle}>{selected?.name}</Text>

            <View style={styles.field}>
              <Text style={styles.label}>Tarih (YYYY-MM-DD)</Text>
              <TextInput
                style={styles.input}
                value={date}
                onChangeText={setDate}
                placeholder="2026-06-15"
                placeholderTextColor="#9ca3af"
              />
            </View>

            <View style={styles.field}>
              <Text style={styles.label}>Saat (HH:MM)</Text>
              <TextInput
                style={styles.input}
                value={time}
                onChangeText={setTime}
                placeholder="14:00"
                placeholderTextColor="#9ca3af"
              />
            </View>

            <View style={styles.modalActions}>
              <TouchableOpacity style={[styles.modalBtn, styles.cancelBtn]} onPress={closeBooking}>
                <Text style={styles.cancelBtnText}>Vazgeç</Text>
              </TouchableOpacity>
              <TouchableOpacity
                style={[styles.modalBtn, styles.confirmBtn, submitting && styles.disabled]}
                onPress={submitBooking}
                disabled={submitting}>
                {submitting ? (
                  <ActivityIndicator color="#fff" />
                ) : (
                  <Text style={styles.confirmBtnText}>Onayla</Text>
                )}
              </TouchableOpacity>
            </View>
          </View>
        </View>
      </Modal>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  safe: { flex: 1, backgroundColor: '#fff' },
  center: { flex: 1, justifyContent: 'center', alignItems: 'center' },
  list: { padding: 16, paddingBottom: 32 },
  title: { fontSize: 22, fontWeight: '700', color: '#111827', marginBottom: 12 },
  card: {
    backgroundColor: '#f9fafb',
    borderRadius: 12,
    padding: 16,
    marginBottom: 12,
    borderWidth: 1,
    borderColor: '#e5e7eb',
  },
  name: { fontSize: 17, fontWeight: '600', color: '#111827' },
  subjects: { fontSize: 13, color: '#2563eb', marginTop: 4 },
  bio: { fontSize: 13, color: '#6b7280', marginTop: 8, lineHeight: 19 },
  bookButton: {
    marginTop: 12,
    backgroundColor: '#2563eb',
    borderRadius: 8,
    paddingVertical: 10,
    alignItems: 'center',
  },
  bookButtonText: { color: '#fff', fontWeight: '600' },
  modalOverlay: {
    flex: 1,
    backgroundColor: 'rgba(0,0,0,0.4)',
    justifyContent: 'center',
    padding: 24,
  },
  modalCard: { backgroundColor: '#fff', borderRadius: 12, padding: 20 },
  modalTitle: { fontSize: 20, fontWeight: '700', color: '#111827' },
  modalSubtitle: { fontSize: 14, color: '#6b7280', marginTop: 4, marginBottom: 20 },
  field: { marginBottom: 14 },
  label: { fontSize: 13, color: '#374151', marginBottom: 6, fontWeight: '600' },
  input: {
    borderWidth: 1,
    borderColor: '#d1d5db',
    borderRadius: 8,
    paddingHorizontal: 12,
    paddingVertical: 10,
    fontSize: 16,
    color: '#111827',
    backgroundColor: '#f9fafb',
  },
  modalActions: { flexDirection: 'row', gap: 8, marginTop: 8 },
  modalBtn: { flex: 1, borderRadius: 8, paddingVertical: 12, alignItems: 'center' },
  cancelBtn: { borderWidth: 1, borderColor: '#d1d5db' },
  cancelBtnText: { color: '#374151', fontWeight: '600' },
  confirmBtn: { backgroundColor: '#2563eb' },
  confirmBtnText: { color: '#fff', fontWeight: '600' },
  disabled: { opacity: 0.6 },
});
