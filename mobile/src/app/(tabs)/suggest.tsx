import { useState } from 'react';
import {
  ActivityIndicator,
  Alert,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  StyleSheet,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { api } from '@/lib/api';

type Suggestion = {
  teacherId: string;
  teacherName: string;
  date: string;
  time: string;
  score: number;
};

const DAYS = ['Pazartesi', 'Salı', 'Çarşamba', 'Perşembe', 'Cuma'];
const DEFAULT_RANGE = '09:00-17:00';

export default function SuggestScreen() {
  const [lessonName, setLessonName] = useState('');
  const [selectedDays, setSelectedDays] = useState<string[]>([]);
  const [timeRange, setTimeRange] = useState(DEFAULT_RANGE);
  const [loading, setLoading] = useState(false);
  const [suggestions, setSuggestions] = useState<Suggestion[]>([]);
  const [booking, setBooking] = useState<string | null>(null);

  function toggleDay(day: string) {
    setSelectedDays((prev) =>
      prev.includes(day) ? prev.filter((d) => d !== day) : [...prev, day],
    );
  }

  async function getSuggestions() {
    if (!lessonName.trim()) {
      Alert.alert('Eksik bilgi', 'Ders adı gerekli.');
      return;
    }
    if (selectedDays.length === 0) {
      Alert.alert('Eksik bilgi', 'En az bir gün seç.');
      return;
    }
    const slots = selectedDays.map((d) => `${d} ${timeRange.trim()}`);
    try {
      setLoading(true);
      const { data } = await api.post('/appointments/suggest', {
        lessonName: lessonName.trim(),
        availableSlots: slots,
      });
      const payload = data.data ?? data;
      setSuggestions(payload.suggestions ?? []);
      if ((payload.suggestions ?? []).length === 0) {
        Alert.alert('Sonuç yok', 'Bu ders ve slotlar için öneri bulunamadı.');
      }
    } catch (err: any) {
      Alert.alert('Hata', err?.response?.data?.error ?? 'Öneri alınamadı');
    } finally {
      setLoading(false);
    }
  }

  async function bookSuggestion(s: Suggestion) {
    try {
      setBooking(`${s.teacherId}-${s.date}-${s.time}`);
      await api.post('/appointments', {
        teacherId: s.teacherId,
        date: s.date,
        time: s.time,
      });
      Alert.alert('Başarılı', 'Randevu oluşturuldu.');
      setSuggestions((prev) =>
        prev.filter((x) => !(x.teacherId === s.teacherId && x.date === s.date && x.time === s.time)),
      );
    } catch (err: any) {
      Alert.alert('Hata', err?.response?.data?.error ?? 'Randevu oluşturulamadı');
    } finally {
      setBooking(null);
    }
  }

  return (
    <SafeAreaView style={styles.safe}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        style={styles.flex}>
        <ScrollView contentContainerStyle={styles.container}>
          <Text style={styles.title}>AI Randevu Önerisi</Text>
          <Text style={styles.subtitle}>
            Ders adı + uygun saatler ver, sistem en iyi öğretmen + zaman kombinasyonunu önersin.
          </Text>

          <View style={styles.field}>
            <Text style={styles.label}>Ders Adı</Text>
            <TextInput
              style={styles.input}
              value={lessonName}
              onChangeText={setLessonName}
              placeholder="Matematik"
              placeholderTextColor="#9ca3af"
            />
          </View>

          <View style={styles.field}>
            <Text style={styles.label}>Uygun Günler</Text>
            <View style={styles.daysRow}>
              {DAYS.map((d) => {
                const active = selectedDays.includes(d);
                return (
                  <TouchableOpacity
                    key={d}
                    style={[styles.dayChip, active && styles.dayChipActive]}
                    onPress={() => toggleDay(d)}>
                    <Text style={[styles.dayChipText, active && styles.dayChipTextActive]}>
                      {d}
                    </Text>
                  </TouchableOpacity>
                );
              })}
            </View>
          </View>

          <View style={styles.field}>
            <Text style={styles.label}>Saat Aralığı</Text>
            <TextInput
              style={styles.input}
              value={timeRange}
              onChangeText={setTimeRange}
              placeholder="09:00-17:00"
              placeholderTextColor="#9ca3af"
            />
          </View>

          <TouchableOpacity
            style={[styles.submit, loading && styles.disabled]}
            onPress={getSuggestions}
            disabled={loading}>
            {loading ? (
              <ActivityIndicator color="#fff" />
            ) : (
              <Text style={styles.submitText}>Öneri Al</Text>
            )}
          </TouchableOpacity>

          {suggestions.length > 0 && (
            <View style={styles.results}>
              <Text style={styles.resultsTitle}>Öneriler</Text>
              {suggestions.map((s) => {
                const key = `${s.teacherId}-${s.date}-${s.time}`;
                const isBooking = booking === key;
                return (
                  <View key={key} style={styles.suggestionCard}>
                    <View style={styles.rowBetween}>
                      <Text style={styles.suggestionTeacher}>{s.teacherName}</Text>
                      <Text style={styles.score}>%{Math.round(s.score * 100)}</Text>
                    </View>
                    <Text style={styles.suggestionDatetime}>
                      {s.date} • {s.time}
                    </Text>
                    <TouchableOpacity
                      style={[styles.bookBtn, isBooking && styles.disabled]}
                      onPress={() => bookSuggestion(s)}
                      disabled={isBooking}>
                      {isBooking ? (
                        <ActivityIndicator color="#fff" />
                      ) : (
                        <Text style={styles.bookBtnText}>Randevu Al</Text>
                      )}
                    </TouchableOpacity>
                  </View>
                );
              })}
            </View>
          )}
        </ScrollView>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  safe: { flex: 1, backgroundColor: '#fff' },
  flex: { flex: 1 },
  container: { padding: 16, paddingBottom: 48 },
  title: { fontSize: 22, fontWeight: '700', color: '#111827' },
  subtitle: { fontSize: 13, color: '#6b7280', marginTop: 6, marginBottom: 20, lineHeight: 19 },
  field: { marginBottom: 16 },
  label: { fontSize: 13, color: '#374151', marginBottom: 8, fontWeight: '600' },
  input: {
    borderWidth: 1,
    borderColor: '#d1d5db',
    borderRadius: 8,
    paddingHorizontal: 12,
    paddingVertical: 12,
    fontSize: 16,
    color: '#111827',
    backgroundColor: '#f9fafb',
  },
  daysRow: { flexDirection: 'row', flexWrap: 'wrap', gap: 8 },
  dayChip: {
    paddingHorizontal: 12,
    paddingVertical: 8,
    borderRadius: 16,
    borderWidth: 1,
    borderColor: '#d1d5db',
    backgroundColor: '#fff',
  },
  dayChipActive: { backgroundColor: '#2563eb', borderColor: '#2563eb' },
  dayChipText: { color: '#374151', fontSize: 13, fontWeight: '500' },
  dayChipTextActive: { color: '#fff' },
  submit: {
    backgroundColor: '#2563eb',
    borderRadius: 8,
    paddingVertical: 14,
    alignItems: 'center',
    marginTop: 8,
  },
  submitText: { color: '#fff', fontWeight: '600', fontSize: 16 },
  disabled: { opacity: 0.6 },
  results: { marginTop: 32 },
  resultsTitle: { fontSize: 18, fontWeight: '700', color: '#111827', marginBottom: 12 },
  suggestionCard: {
    backgroundColor: '#f9fafb',
    borderRadius: 12,
    padding: 14,
    marginBottom: 10,
    borderWidth: 1,
    borderColor: '#e5e7eb',
  },
  rowBetween: { flexDirection: 'row', justifyContent: 'space-between', alignItems: 'center' },
  suggestionTeacher: { fontSize: 16, fontWeight: '600', color: '#111827' },
  score: {
    fontSize: 12,
    fontWeight: '700',
    color: '#15803d',
    backgroundColor: '#dcfce7',
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 12,
  },
  suggestionDatetime: { fontSize: 14, color: '#374151', marginTop: 6 },
  bookBtn: {
    backgroundColor: '#2563eb',
    borderRadius: 8,
    paddingVertical: 10,
    alignItems: 'center',
    marginTop: 10,
  },
  bookBtnText: { color: '#fff', fontWeight: '600' },
});
