import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
  } from "react-router-dom";

import Main from './Main';

export default function Menu(props) {
    return (
        <Router>
            <footer className="footer">
                <Link to="/" className="footer__button button">❏</Link>
                <Link to="/achievements" className="footer__button button">★</Link>
                <Link to="/profile" className="footer__button button">◓</Link>
            </footer>
            <Switch>
                <Route path="/achievements">
                    <Achievements />
                </Route>
                <Route path="/profile">
                    <Profile />
                </Route>
                <Route path="/">
                    <Main />
                </Route>
                <Route path="/">
                    <Main />
                </Route>
            </Switch>
        </Router>
    )
    
    function Achievements() {
     return <h2>Ачивки</h2>;
    }
    
    function Profile() {
        return <h2>Профиль</h2>;
    }
}