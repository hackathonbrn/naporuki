import React from 'react';
import { Redirect, Link, withRouter } from "react-router-dom";

import axios from 'axios';


class Login extends React.Component {
    
    constructor(props) {
        super(props);   
        this.state = {      
            phone: '',      
            password: '',
            loading: false,
            isSignedUp: false   
        };  
    }

    signup(){
        const apiUrl = "http://localhost:8080/api/v1/login";
        return axios.post(apiUrl, {phone: this.state.phone, password: this.state.password}).then((resp) => {
            this.setState({loading: false});
            if (resp.data) {
                document.cookie = `jwt=${resp.data}`;
                this.setState({ isSignedUp: true });
                this.props.history.push("/dashboard");
            } 
        });
    }

    handleLogin = (e) => {
        e.preventDefault();
        this.setState({stateloading: true});
        this.signup();
      };

      render() {
        if (this.state.isSignedUp) return <Redirect to="/dashboard" />;
        if (this.props.isAuth) return <Redirect to="/dashboard" />;

        return (
            <div className="login__wrapper">
            <h2 className="login__header">Войти</h2>
            <form className="login__form" onSubmit={this.handleLogin}>
                <label className="login__label" htmlFor="phone">Номер телефона</label>
                <input required className="login__input" disabled={this.state.loading} type="tel" placeholder="+7" value={this.state.phone} onChange={(e) => this.setState({phone:  e.target.value})} name="phone" id="phone" />
                <label className="login__label" htmlFor="password">Пароль</label>
                <input required className="login__input" disabled={this.state.loading} type="password" placeholder="Введите пароль" value={this.state.password} onChange={(e) => this.setState({password:  e.target.value})} name="password" id="password" />
                <input className="login__submit" disabled={this.state.loading} type="submit" value="Войти" />
            </form>
            <Link className="login__small" to="/register">Еще не зарегистрированы?</Link>
            </div>
        )
      }
    
};

export default withRouter(Login);
