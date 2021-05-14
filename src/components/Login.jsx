import axios from "axios";
import React, { useState } from "react";
import { Redirect } from "react-router";
// import axios from 'axios';

export default function Login(props) {
    const [name, setName] = useState('');
    const [phone, setPhone] = useState('');
    const [password, setPassword] = useState('');
    const [loading, setLoading] = useState(false);

  if (props.isAuth) return <Redirect to="/" />;

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true);
    const apiUrl = "http://api:8080/register-teacher";
    axios.post(apiUrl, {name: name, phone: phone, password: password}).then((resp) => {
        setLoading(false);
        if (resp.text === "success") { return (<Redirect to="/" />) }
    });
  };

  return (
    <div className="login__wrapper">
      <h2 className="login__header">Регистрация</h2>
      <form className="login__form" onSubmit={handleSubmit}>
        <input className="login__input" disabled={loading} type="text" placeholder="Ваше имя" value={name} onChange={(e) => setName(e.target.value)} name="name" />
        <input className="login__input" disabled={loading} type="tel" placeholder="Номер телефона" value={phone} onChange={(e) => setPhone(e.target.value)} name="phone" />
        <input className="login__input" disabled={loading} type="password" placeholder="Пароль" value={password} onChange={(e) => setPassword(e.target.value)} name="password" />
        <input className="login__submit" disabled={loading} type="submit" value="Отправить" />
      </form>
    </div>
  );
}
