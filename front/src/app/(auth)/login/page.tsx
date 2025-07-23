"use client";

import Cookies from 'js-cookie';
import { useState, useRef } from "react";

export default function Home() {
  const inputRef = useRef<HTMLInputElement>(null);
  const [result, setResult] = useState<string>("");

  const handleLogin = () => {
    setResult("Загрузка...");
    const email = (document.getElementById("loginEmail") as HTMLInputElement)?.value;
    const password = (document.getElementById("loginPassword") as HTMLInputElement)?.value;

    const postData = {
      email: email,
      password: password,
      appId: 1,
    };

    fetch("http://" + process.env.NEXT_PUBLIC_ADDRESS_AUTH + "/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(postData),
    })
      .then((response) => {
        if (!response.ok) {
          if (response.status == 400) {
            throw new Error("некорректный ввод");
          }
          if (response.status == 500) {
            throw new Error("внутренняя ошибка сервера");
          }
          throw new Error("ошибка сети");
        } 
        return response.json();
      })
      .then((data) => {
        setResult(data["token"]);
        Cookies.set('token', data["token"], {
          expires: 1/24,
          // secure: true, // только HTTPS
          sameSite: 'strict'
        });
      })
      .catch((error) => {
        setResult("Ошибка: " + error.message);
      });
  };

  return (
    <div className="block">
        <h3>Вход</h3>
        <input ref={inputRef} type="email" id="loginEmail" placeholder="Email"/>
        <input ref={inputRef} type="password" id="loginPassword" placeholder="Password"/>
        <button id="loginBtn" onClick={handleLogin}>Войти</button>
        <p style={{ textAlign: "center" }}>
          <a href="register" text-align="center">
            Нет аккаунта? Зарегистрируйтесь!
          </a>
        </p>
        <div className="result" id="loginResult">{result}</div>
    </div>
  );
}