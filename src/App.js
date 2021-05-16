import Header from './components/Header';

import Main from './components/Main';
import Profile from './components/Profile';
import Registration from './components/Registration';
import Login from './components/Login';
import ProfileTeacherForm from './components/ProfileTeacherForm';

import { BrowserRouter as Router, Switch, Route, NavLink, Redirect } from 'react-router-dom';

import React, { useState } from 'react';

import axios from 'axios';

axios.defaults.withCredentials = true;

function App() {

  const [isAuth, setAuth] = useState(null);

  function checkAuth() {
    if (!document.cookie) {
      setAuth(false);
    }
    const apiUrl = 'http://localhost:8080/api/v1/check-auth';
    axios.get(apiUrl).then((resp) => {
      if (resp.data === true) {
        setAuth(true);
      }
    });
  };

  const PrivateRoute = ({component: Component, ...rest}) => {
    return (
        <Route {...rest} render={props => (
            isAuth ?
                <Component {...props} />
            : <Redirect to="/login" />
        )} />
    );
  };

  const PublicRoute = ({component: Component, restricted, ...rest}) => {
    return (
        <Route {...rest} render={props => (
          isAuth && restricted ?
                <Redirect to="/dashboard" />
            : <Component {...props} />
        )} />
    );
  };

  return (
    <div>
      <Header />

      <Router>
        <nav className="footer">
          <NavLink activeClassName="nav-item__active" to="/dashboard" className="nav-item">
            <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 0 24 24" width="24px" fill="#000000">
              <path d="M0 0h24v24H0z" fill="none" />
              <path d="M3 13h8V3H3v10zm0 8h8v-6H3v6zm10 0h8V11h-8v10zm0-18v6h8V3h-8z" />
            </svg>
          </NavLink>
          <NavLink activeClassName="nav-item__active" to="/achievements" className="nav-item">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              enableBackground="new 0 0 24 24"
              height="24px"
              viewBox="0 0 24 24"
              width="24px"
              fill="#000000"
            >
              <rect fill="none" height="24" width="24" />
              <path d="M19,5h-2V3H7v2H5C3.9,5,3,5.9,3,7v1c0,2.55,1.92,4.63,4.39,4.94c0.63,1.5,1.98,2.63,3.61,2.96V19H7v2h10v-2h-4v-3.1 c1.63-0.33,2.98-1.46,3.61-2.96C19.08,12.63,21,10.55,21,8V7C21,5.9,20.1,5,19,5z M5,8V7h2v3.82C5.84,10.4,5,9.3,5,8z M19,8 c0,1.3-0.84,2.4-2,2.82V7h2V8z" />
            </svg>
          </NavLink>
          <NavLink activeClassName="nav-item__active" to="/profile" className="nav-item">
            <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 0 24 24" width="24px" fill="#000000">
              <path d="M0 0h24v24H0z" fill="none" />
              <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z" />
            </svg>
          </NavLink>
        </nav>

        <Switch>
          <Redirect exact from="/" to="/dashboard" />

          {/* <Route exact path="/dashboard">
            <Main isAuth={isAuth} />
          </Route> */}

          <PrivateRoute component={Main} path="/dashboard" exact />

          {/* <Route exact path="/register">
            <Registration isAuth={isAuth} />
          </Route> */}

          <PublicRoute restricted={false} component={Registration} path="/register" exact />

          {/* <Route exact path="/achievements">
            <Achievements />
          </Route> */}

          <PrivateRoute component={Achievements} path="/achievements" exact />
          

          {/* <Route exact path="/profile">
            <Profile isAuth={isAuth} />
          </Route> */}

          <PrivateRoute component={Profile} path="/profile" exact />

          {/* <Route exact path="/login">
            <Login isAuth={isAuth} />
          </Route> */}

          <PublicRoute restricted={false} component={Login} path="/login" exact />


          
          {/* <Route exact path="/teacher-form">
            <ProfileTeacherForm isAuth={isAuth} />
          </Route> */}

          <PrivateRoute component={ProfileTeacherForm} path="/teacher-form" exact />


        </Switch>
      </Router>
    </div>
  );
  function Achievements() {
    return <h2>Достижения</h2>;
  }
}

export default App;
