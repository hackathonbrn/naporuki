import React, { useState, useEffect } from 'react';

import Categories from './Categories';
import Search from './Search';
import Popular from './Popular';

import axios from 'axios';

export default function Main(props) {

  const [profiles, setProfiles] = useState([]);

  useEffect(() => {
    const apiUrl = "http://localhost:8080/api/v1/get-all-profiles";
    axios.get(apiUrl).then((resp) => {
      if (resp.data) {
        setProfiles([]);
      };
    });
  }, []);

    return (
        <main className="main">

        <h2 style={{marginBottom: "30px"}}>Объявления</h2>

        <Search />

        <section className="categories">
          <h2 className="categories__title">Дисциплины</h2>
          <Categories items={['Математика', 'Физика', 'Биология', 'Русский язык', 'Химия', 'Информатика', 'Английский язык']}/>
        </section>

        <section className="popular">
          <h2 className="popular__title">Пользователи</h2>
          <Popular 
          items={profiles}
          />  
        </section>
      </main>
    )
}
