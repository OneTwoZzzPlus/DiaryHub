"use client";

import { useState, useEffect } from 'react';
import Cookies from 'js-cookie';

export default function Home() {
  return (
    <div>
      <p>Ваш JWT токен </p>{getToken()}
      <ul>
          <li><a href="/login">Вход</a></li>
          <li><a href="/register">Регистрация</a></li>
      </ul>
    </div>
  );
}


function getToken() {
  const [token, setToken] = useState<string | null>(null);
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
    const jwtToken = Cookies.get('token');
    setToken(jwtToken?.toString() || null);
  }, []);

  return (    
      isClient ? (
        token ? (
          <i>{token}</i>
        ) : (
          <b>не найден</b>
        )
      ) : (
        <p>загружается...</p>
      )
  );
}