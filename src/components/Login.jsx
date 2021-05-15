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
    const apiUrl = "http://localhost:8080/api/v1/register-teacher";
    axios.post(apiUrl, {name: name, phone: phone, password: password}).then((resp) => {
        setLoading(false);
        if (resp.data) {
            // console.log(resp.data);
            document.cookie = `jwt=${resp.data}`;
            return (<Redirect to="/" />) 
        }
    });
  };

  return (
    <div className="login__wrapper">
      <h2 className="login__header">Регистрация</h2>
      <form className="login__form" onSubmit={handleSubmit}>
        <label className="login__label" htmlFor="name">Имя</label>
        <input required className="login__input" disabled={loading} type="text" placeholder="Введите имя" value={name} onChange={(e) => setName(e.target.value)} name="name" id="name" />
        <label className="login__label" htmlFor="phone">Номер телефона</label>
        <input required className="login__input" disabled={loading} type="tel" placeholder="+7" value={phone} onChange={(e) => setPhone(e.target.value)} name="phone" id="phone" />
        <label className="login__label" htmlFor="password">Пароль</label>
        <input required className="login__input" disabled={loading} type="password" placeholder="Введите пароль" value={password} onChange={(e) => setPassword(e.target.value)} name="password" id="password" />
        <input className="login__submit" disabled={loading} type="submit" value="Сохранить" />
      </form>
    </div>
  );
}
