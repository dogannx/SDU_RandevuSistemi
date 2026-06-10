import { Stack } from 'expo-router';
import { useCallback, useEffect, useState } from 'react';
import {
  ActivityIndicator,
  Alert,
  FlatList,
  RefreshControl,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { api } from '@/lib/api';

type Student = {
  id: string;
  name: string;
  email: string;
  createdAt: string;
};

export default function StudentsScreen() {
  const [items, setItems] = useState<Student[]>([]);
  const [loading, setLoading] = useState(true);
  const [refreshing, setRefreshing] = useState(false);

  const load = useCallback(async () => {
    try {
      const { data } = await api.get('/students');
      setItems(data.data ?? []);
    } catch (err: any) {
      if (err?.__silent) return;
      Alert.alert('Hata', err?.response?.data?.error ?? 'Öğrenciler yüklenemedi');
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

  if (loading) {
    return (
      <SafeAreaView style={styles.safe}>
        <Stack.Screen options={{ title: 'Tüm Öğrenciler' }} />
        <View style={styles.center}>
          <ActivityIndicator size="large" color="#2563eb" />
        </View>
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView style={styles.safe}>
      <Stack.Screen options={{ title: 'Tüm Öğrenciler' }} />
      <FlatList
        data={items}
        keyExtractor={(s) => s.id}
        contentContainerStyle={styles.list}
        ListEmptyComponent={<Text style={styles.empty}>Kayıtlı öğrenci yok.</Text>}
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={onRefresh} />}
        renderItem={({ item }) => (
          <View style={styles.card}>
            <Text style={styles.name}>{item.name}</Text>
            <Text style={styles.email}>{item.email}</Text>
            <Text style={styles.date}>
              Kayıt: {new Date(item.createdAt).toLocaleDateString('tr-TR')}
            </Text>
          </View>
        )}
      />
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  safe: { flex: 1, backgroundColor: '#fff' },
  center: { flex: 1, justifyContent: 'center', alignItems: 'center' },
  list: { padding: 16 },
  empty: { textAlign: 'center', color: '#6b7280', marginTop: 48 },
  card: {
    backgroundColor: '#f9fafb',
    borderRadius: 12,
    padding: 14,
    marginBottom: 10,
    borderWidth: 1,
    borderColor: '#e5e7eb',
  },
  name: { fontSize: 16, fontWeight: '600', color: '#111827' },
  email: { fontSize: 13, color: '#6b7280', marginTop: 4 },
  date: { fontSize: 12, color: '#9ca3af', marginTop: 6 },
});
