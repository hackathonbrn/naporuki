import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Categories from "./Categories";

import axios from "axios";

export default function Profile(props) {
  const [description, setDescription] = useState("");
  const [user, setUser] = useState({
    name: "",
    rating: 0,
    phone: "",
    subjects: [],
  });

  useEffect(() => {
    const apiUrl = "http://localhost:8080/api/v1/get-teacher-profile";
    axios
      .get(apiUrl)
      .then((resp) => {
        if (resp.data) {
          console.log(resp);
          setDescription(resp.data.desc);
          setUser(resp.data.user);
        }
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  return (
    <div className="porfile">
      <h2 className="profile__header">Профиль</h2>
      <div className="profile__card">
        <img
          src="https://randomuser.me/api/portraits/lego/2.jpg"
          alt="Фото профиля"
          className="profile__photo"
        />
        <h3 className="profile__subheader">Инфо</h3>
        <div className="profile__info">
            <span>Имя: {user.name}</span>
            <span>Рейтинг: {user.rating}/5</span>
            <span>Телефон: {user.phone}</span>
        </div>


        <h3 className="profile__subheader">Предметы</h3>

        <Categories items={user.subjects} />

        <h3 className="profile__subheader">О себе</h3>
        <p className="profile__description">
          {description} Lorem ipsum dolor sit amet consectetur adipisicing elit.
          Impedit delectus sunt ex quo at consectetur, beatae doloremque minus
          laudantium fugit.
        </p>
        <div className="profile__teacher-wrapper">
        <Link className="profile__teacher" to="/teacher-form">
            Стать преподавателем
        </Link>
      </div>
      </div>
    </div>
  );
}
