import React, { useEffect, useState } from 'react';
import { Link, Redirect } from 'react-router-dom';
import Categories from './Categories';

import axios from 'axios';

export default function Profile(props) {

    const [description, setDescription] = useState('');
    const [user, setUser] = useState({
        name: '',
        rating: 0.0,
        phone: '',
        subjects: []
    });

    useEffect(() => {
        const apiUrl = "http://localhost:8080/api/v1/get-teacher-profile";
        axios.get(apiUrl).then((resp) => {
          if (resp.data) {
              setDescription(resp.desc);
              setUser(resp.user);
          };
        }).catch(error => {
            console.log(error);
        });
      }, []);

    if (!props.isAuth) return (<Redirect to="/login" />);

    return (
        <div className="porfile">
            <div className="profile__card">
                <h2 className="profile__header">Профиль</h2>
                <img src="https://randomuser.me/api/portraits/lego/2.jpg" alt="Фото профиля" className="profile__label" />

                <span>Имя: {user.name}</span>
                <span>Рейтинг: {user.rating}/5</span>
                <span>Телефон: {user.phone}</span>

                <h3 className="profile__subheader">Предметы</h3>
                <Categories items={user.subjects}/>

                <h3 className="profile__subheader">Описание</h3>
                <p>{description}</p>
            </div>

            <Link className="profile__teacher" to="/teacher-form">Стать преподавателем</Link>
        </div>
    )
}