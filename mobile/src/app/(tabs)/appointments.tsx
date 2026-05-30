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

type Appointment = {
  id: string;
  studentId: string;
  teacherId: string;
  date: string;
  time: string;
  status: string;
  teacher?: { id: string; name: string };
};

export default function AppointmentsScreen() {
  const [items, setItems] = useState<Appointment[]>([]);
  const [loading, setLoading] = useState(true);
  const [refreshing, setRefreshing] = useState(false);
  const [editing, setEditing] = useState<Appointment | null>(null);
  const [date, setDate] = useState('');
  const [time, setTime] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const load = useCallback(async () => {
    try {
      const { data } = await api.get('/appointments');
      setItems(data.data ?? []);
    } catch (err: any) {
      Alert.alert('Hata', err?.response?.data?.error ?? 'Randevular yüklenemedi');
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

  function openEdit(a: Appointment) {
    setEditing(a);
    setDate(a.date);
    setTime(a.time);
  }

  function closeEdit() {
    setEditing(null);
  }

  async function submitEdit() {
    if (!editing) return;
    if (!date.trim() || !time.trim()) {
      Alert.alert('Eksik bilgi', 'Tarih ve saat gerekli.');
      return;
    }
    try {
      setSubmitting(true);
      await api.put(`/appointments/${editing.id}`, {
        date: date.trim(),
        time: time.trim(),
      });
      closeEdit();
      await load();
    } catch (err: any) {
      Alert.alert('Hata', err?.response?.data?.error ?? 'Randevu güncellenemedi');
    } finally {
      setSubmitting(false);
    }
  }

  async function confirmCancel(a: Appointment) {
    Alert.alert(
      'Randevuyu iptal et',
      `${a.date} ${a.time} randevusunu iptal etmek istediğine emin misin?`,
      [
        { text: 'Vazgeç', style: 'cancel' },
        {
          text: 'İptal Et',
          style: 'destructive',
          onPress: async () => {
            try {
              await api.delete(`/appointments/${a.id}`);
              await load();
            } catch (err: any) {
              Alert.alert('Hata', err?.response?.data?.error ?? 'İptal edilemedi');
            }
          },
        },
      ],
    );
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
        data={items}
        keyExtractor={(a) => a.id}
        contentContainerStyle={styles.list}
        ListHeaderComponent={<Text style={styles.title}>Randevularım</Text>}
        ListEmptyComponent={<Text style={styles.empty}>Henüz randevun yok.</Text>}
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={onRefresh} />}
        renderItem={({ item }) => (
          <View style={styles.card}>
            <View style={styles.rowBetween}>
              <Text style={styles.teacher}>{item.teacher?.name ?? '—'}</Text>
              <Text style={[styles.status, statusColor(item.status)]}>
                {statusLabel(item.status)}
              </Text>
            </View>
            <Text style={styles.datetime}>
              {item.date} • {item.time}
            </Text>

            {item.status === 'upcoming' && (
              <View style={styles.actions}>
                <TouchableOpacity style={styles.editBtn} onPress={() => openEdit(item)}>
                  <Text style={styles.editBtnText}>Güncelle</Text>
                </TouchableOpacity>
                <TouchableOpacity style={styles.cancelBtn} onPress={() => confirmCancel(item)}>
                  <Text style={styles.cancelBtnText}>İptal</Text>
                </TouchableOpacity>
              </View>
            )}
          </View>
        )}
      />

      <Modal visible={!!editing} animationType="slide" transparent onRequestClose={closeEdit}>
        <View style={styles.modalOverlay}>
          <View style={styles.modalCard}>
            <Text style={styles.modalTitle}>Randevu Güncelle</Text>
            <Text style={styles.modalSubtitle}>{editing?.teacher?.name}</Text>

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
              <TouchableOpacity style={[styles.modalBtn, styles.modalCancelBtn]} onPress={closeEdit}>
                <Text style={styles.modalCancelBtnText}>Vazgeç</Text>
              </TouchableOpacity>
              <TouchableOpacity
                style={[styles.modalBtn, styles.confirmBtn, submitting && styles.disabled]}
                onPress={submitEdit}
                disabled={submitting}>
                {submitting ? (
                  <ActivityIndicator color="#fff" />
                ) : (
                  <Text style={styles.confirmBtnText}>Kaydet</Text>
                )}
              </TouchableOpacity>
            </View>
          </View>
        </View>
      </Modal>
    </SafeAreaView>
  );
}

function statusLabel(s: string) {
  if (s === 'upcoming') return 'Yaklaşan';
  if (s === 'past') return 'Geçmiş';
  if (s === 'cancelled') return 'İptal Edildi';
  return s;
}

function statusColor(s: string) {
  if (s === 'upcoming') return { color: '#15803d' };
  if (s === 'past') return { color: '#6b7280' };
  if (s === 'cancelled') return { color: '#dc2626' };
  return { color: '#374151' };
}

const styles = StyleSheet.create({
  safe: { flex: 1, backgroundColor: '#fff' },
  center: { flex: 1, justifyContent: 'center', alignItems: 'center' },
  list: { padding: 16, paddingBottom: 32 },
  title: { fontSize: 22, fontWeight: '700', color: '#111827', marginBottom: 12 },
  empty: { textAlign: 'center', color: '#6b7280', marginTop: 48 },
  card: {
    backgroundColor: '#f9fafb',
    borderRadius: 12,
    padding: 16,
    marginBottom: 12,
    borderWidth: 1,
    borderColor: '#e5e7eb',
  },
  rowBetween: { flexDirection: 'row', justifyContent: 'space-between', alignItems: 'center' },
  teacher: { fontSize: 17, fontWeight: '600', color: '#111827' },
  status: { fontSize: 12, fontWeight: '600' },
  datetime: { fontSize: 14, color: '#374151', marginTop: 6 },
  actions: { flexDirection: 'row', gap: 8, marginTop: 12 },
  editBtn: {
    flex: 1,
    borderWidth: 1,
    borderColor: '#2563eb',
    borderRadius: 8,
    paddingVertical: 10,
    alignItems: 'center',
  },
  editBtnText: { color: '#2563eb', fontWeight: '600' },
  cancelBtn: {
    flex: 1,
    borderWidth: 1,
    borderColor: '#dc2626',
    borderRadius: 8,
    paddingVertical: 10,
    alignItems: 'center',
  },
  cancelBtnText: { color: '#dc2626', fontWeight: '600' },

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
  modalCancelBtn: { borderWidth: 1, borderColor: '#d1d5db' },
  modalCancelBtnText: { color: '#374151', fontWeight: '600' },
  confirmBtn: { backgroundColor: '#2563eb' },
  confirmBtnText: { color: '#fff', fontWeight: '600' },
  disabled: { opacity: 0.6 },
});
