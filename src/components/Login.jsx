import React from 'react'
import { Redirect } from 'react-router'

export default function Login(props) {
    if (props.isAuth) return (<Redirect to='/' />);
    return (
        <div>
            <form action="http://go:8080/register-teacher" method="post">
                <input type="text" name="name" />
                <input type="tel" name="phone" />
                <input type="password" name="password" />
                <input type="submit" value="Send" />
            </form>
        </div>
    )
}
