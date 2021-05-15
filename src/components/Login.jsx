import React, { useState } from 'react';
import { Redirect, Link } from "react-router-dom";

import axios from 'axios';


export default function Login(props) {

    const [phone, setPhone] = useState('');
    const [password, setPassword] = useState('');
    const [loading, setLoading] = useState(false);

    if (props.isAuth) return <Redirect to="/dashboard" />;

    const handleLogin = (e) => {
        e.preventDefault();
        setLoading(true);
        const apiUrl = "http://localhost:8080/api/v1/login";
        axios.post(apiUrl, {phone: phone, password: password}).then((resp) => {
            setLoading(false);
            if (resp.data) {
                document.cookie = `jwt=${resp.data}`;
            }
            return (<Redirect to="/dashboard" />);
        });
      };

    return (
        <div className="login__wrapper">
        <h2 className="login__header">Войти</h2>
        <form className="login__form" onSubmit={handleLogin}>
            <label className="login__label" htmlFor="phone">Номер телефона</label>
            <input required className="login__input" disabled={loading} type="tel" placeholder="+7" value={phone} onChange={(e) => setPhone(e.target.value)} name="phone" id="phone" />
            <label className="login__label" htmlFor="password">Пароль</label>
            <input required className="login__input" disabled={loading} type="password" placeholder="Введите пароль" value={password} onChange={(e) => setPassword(e.target.value)} name="password" id="password" />
            <input className="login__submit" disabled={loading} type="submit" value="Войти" />
        </form>
        <Link className="login__small" to="/register">Еще не зарегистрированы?</Link>
        </div>
    )
}
