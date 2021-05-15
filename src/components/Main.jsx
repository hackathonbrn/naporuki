import React from 'react';

import Categories from './Categories';
import Search from './Search';
import Popular from './Popular';

import { Redirect } from "react-router-dom";

export default function Main(props) {


  if (!props.isAuth) return (<Redirect to="/login" />);


    return (
    <div>
        <main className="main">

        <h2 style={{marginBottom: "30px"}}>Объявления</h2>

        <Search />

        <section className="categories">
          <h2 className="categories__title">Дисциплины</h2>
          <Categories items={['Математика', 'Физика', 'Биология', 'Русский язык', 'Химия', 'Информатика', 'Английский язык']}/>
        </section>



        <section className="popular">
          <h2 className="popular__title">Популярное</h2>
          <Popular 
          items={
            [
              {name: 'Иван Степанов', rating: 4.9, photo: 'https://randomuser.me/api/portraits/men/43.jpg'},
              {name: 'Геннадий Горин', rating: 5.0, photo: 'https://randomuser.me/api/portraits/men/32.jpg'},
              {name: 'Лариса Долина', rating: 4.7, photo: 'https://randomuser.me/api/portraits/women/43.jpg'},
            ]
          }
          />  
        </section>
      </main>

      </div>
    )
}
