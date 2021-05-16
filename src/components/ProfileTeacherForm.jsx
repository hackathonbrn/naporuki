import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';

import axios from 'axios';

import Select from "./Select";

export default function ProfileTeacherForm(props) {
    const [description, setDescription] = useState('');
    const [subject, setSubject] = useState('');
    const [loading, setLoading] = useState(false);

    const handleTeacher = (e) => {
        e.preventDefault();
        setLoading(!loading);
        const apiUrl = "http://localhost:8080/api/v1/create-teacher-profile";
        axios.post(apiUrl, {desc: description, subjects: [subject]}).then((resp) => {
            setLoading(!loading);
            if (resp.data === "success") {
                return (<Redirect to="/profile" />);
            }
        });
      };

    return (
        <div className="teacher-form__wrapper">
            <h2 className="teacher-form__header">Стать преподавателем</h2>
            <form className="teacher-form__form" onSubmit={handleTeacher}>
                <label className="teacher-form__label" htmlFor="desc">О себе</label>
                <textarea placeholder="Чем Вы интересуетесь, чем занимаетесь, на каких предметах специализируетесь?" className="teacher-form__field teacher-form__textarea" id="desc" type="text" name="desc" value={description} onChange={(e) => setDescription(e.target.value)} />
                <label className="teacher-form__label" htmlFor="subject">Предмет</label>
                <select className="teacher-form__field" id="subject" name="subjects" value={subject} onChange={(e) => setSubject(e.target.value)}>
                    <Select value="Математика"/>
                    <Select value="Физика" />
                    <Select value="Химия" />
                </select>
                <input className="teacher-form__submit" type="submit" value="Сохранить" />
            </form>
        </div>
    )
}
