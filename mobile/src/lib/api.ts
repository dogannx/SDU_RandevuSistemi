import axios from 'axios';
import Constants from 'expo-constants';
import * as SecureStore from 'expo-secure-store';

const ACCESS_TOKEN_KEY = 'randevu_access_token';
const STUDENT_KEY = 'randevu_student';

const baseURL =
  (Constants.expoConfig?.extra?.apiBaseUrl as string | undefined) ??
  'https://sdu-randevusistemi.onrender.com/api/v1';

export const api = axios.create({
  baseURL,
  timeout: 15000,
});

function isPublicPath(url: string | undefined): boolean {
  if (!url) return false;
  return url.startsWith('/auth/') || url === '/teachers';
}

api.interceptors.request.use(async (config) => {
  const token = await SecureStore.getItemAsync(ACCESS_TOKEN_KEY);
  if (!token && !isPublicPath(config.url)) {
    return Promise.reject({ __silent: true, message: 'no-token' });
  }
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

api.interceptors.response.use(
  (res) => res,
  async (err) => {
    if (err?.response?.status === 401) {
      await SecureStore.deleteItemAsync(ACCESS_TOKEN_KEY);
      await SecureStore.deleteItemAsync(STUDENT_KEY);
      err.__silent = true;
    }
    return Promise.reject(err);
  },
);

export async function saveAccessToken(token: string) {
  await SecureStore.setItemAsync(ACCESS_TOKEN_KEY, token);
}

export async function clearAccessToken() {
  await SecureStore.deleteItemAsync(ACCESS_TOKEN_KEY);
}

export async function getAccessToken() {
  return SecureStore.getItemAsync(ACCESS_TOKEN_KEY);
}

export async function saveStudent(student: unknown) {
  await SecureStore.setItemAsync(STUDENT_KEY, JSON.stringify(student));
}

export async function getStudent<T = unknown>(): Promise<T | null> {
  const raw = await SecureStore.getItemAsync(STUDENT_KEY);
  return raw ? (JSON.parse(raw) as T) : null;
}

export async function clearStudent() {
  await SecureStore.deleteItemAsync(STUDENT_KEY);
}
