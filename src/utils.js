import axios from 'axios';

const TOKEN_KEY = 'jwt';

export const login = (jwt) => {
  localStorage.setItem(TOKEN_KEY, jwt);
  document.cookie = `jwt=${jwt}`;
};

export const logout = () => {
  localStorage.removeItem(TOKEN_KEY);
};

export const isLogin = () => {
    // return true;
  if (localStorage.getItem(TOKEN_KEY)) {
    document.cookie = `${TOKEN_KEY}=${localStorage.getItem(TOKEN_KEY)}`;
    const apiUrl = 'http://localhost:8080/api/v1/check-auth';
    axios.get(apiUrl).then((resp) => {
      if (resp.data === true) {
        return true;
      }
    });
  }
  return false;
};
