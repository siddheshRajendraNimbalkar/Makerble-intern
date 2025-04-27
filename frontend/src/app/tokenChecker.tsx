'use client';

import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

export function useTokenChecker() {
  const router = useRouter();

  useEffect(() => {
    const accessToken = localStorage.getItem('access_token');
    const tokenExpiry = localStorage.getItem('token_expiry');

    if (!accessToken || !tokenExpiry) {
      router.push('/login');
      return;
    }

    const isExpired = new Date(tokenExpiry).getTime() < new Date().getTime();

    if (isExpired) {
      localStorage.removeItem('access_token');
      localStorage.removeItem('user_id');
      localStorage.removeItem('token_expiry');
      router.push('/login');
    }
  }, [router]);
}
