import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Register",
  description: "Страница регистрации",
};

export default function Home() {
  return (
    <div className="block">
        <h3>Регистрация</h3>
        <input type="email" id="regEmail" placeholder="Email"/>
        <input type="password" id="regPassword" placeholder="Password"/>
        <button id="registerBtn">Зарегистрироваться</button>
        <p><a href="login">Уже есть аккаунт? Войдите!</a></p>
        <div className="result" id="registerResult"></div>
    </div>
  );
}
