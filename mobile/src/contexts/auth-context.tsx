import { createContext, useCallback, useContext, useEffect, useMemo, useState, type ReactNode } from 'react';

import { api, clearAccessToken, getAccessToken, saveAccessToken } from '@/lib/api';

type Student = {
  id: string;
  name: string;
  email: string;
};

type AuthContextValue = {
  student: Student | null;
  token: string | null;
  isLoading: boolean;
  login: (email: string, password: string) => Promise<void>;
  register: (name: string, email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
};

const AuthContext = createContext<AuthContextValue | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [student, setStudent] = useState<Student | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    (async () => {
      const stored = await getAccessToken();
      if (stored) {
        setToken(stored);
      }
      setIsLoading(false);
    })();
  }, []);

  const login = useCallback(async (email: string, password: string) => {
    const { data } = await api.post('/auth/login', { email, password });
    const accessToken = data.data?.accessToken ?? data.accessToken;
    const studentInfo = data.data?.student ?? data.student;
    await saveAccessToken(accessToken);
    setToken(accessToken);
    setStudent(studentInfo);
  }, []);

  const register = useCallback(async (name: string, email: string, password: string) => {
    const { data } = await api.post('/auth/register', { name, email, password });
    const accessToken = data.data?.accessToken ?? data.accessToken;
    const studentInfo = data.data?.student ?? data.student;
    await saveAccessToken(accessToken);
    setToken(accessToken);
    setStudent(studentInfo);
  }, []);

  const logout = useCallback(async () => {
    await clearAccessToken();
    setToken(null);
    setStudent(null);
  }, []);

  const value = useMemo<AuthContextValue>(
    () => ({ student, token, isLoading, login, register, logout }),
    [student, token, isLoading, login, register, logout],
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuth() {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error('useAuth AuthProvider içinde kullanılmalı');
  return ctx;
}
