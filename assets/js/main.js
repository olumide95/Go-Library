const API_BASE = 'http://localhost:8080' 
const SIGNIN_URL = '/login'
const SIGNUP_URL = '/signup'
const ADMIN_ROLE = 'Admin'

document.getElementById('signin-form').addEventListener('submit', (e) => {
    e.preventDefault();

    const data = new FormData(e.target);
    const request = Object.fromEntries(data.entries());

    fetchData('POST', SIGNIN_URL, JSON.stringify(request))
    .then((reponse) => {
        localStorage.setItem('user', JSON.stringify(reponse.user))
        localStorage.setItem('access_token', reponse.access_token)

        if(reponse.user.role == ADMIN_ROLE){
            location.pathname = '/admin'
            return
        }

        location.pathname = '/user'
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });
});

document.getElementById('signup-form')?.addEventListener('submit', (e) => {
    e.preventDefault();

    const data = new FormData(e.target);
    const request = Object.fromEntries(data.entries());

    fetchData('POST', SIGNUP_URL, JSON.stringify(request))
    .then((reponse) => {
        localStorage.setItem('user', JSON.stringify(reponse.user))
        localStorage.setItem('access_token', reponse.access_token)
        if(reponse.user.role == ADMIN_ROLE){
            location.pathname = '/admin'
            return
        }

        location.pathname = '/user'
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });
});

const fetchData = (method, url, data) => {
    return fetch(SIGNIN_URL, {
        method: method,
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin', 
        headers: {
          'Content-Type': 'application/json'
        },
        body: data
      }).then((response) => {
        if (response.ok) {
          return response.json();
        }
        return Promise.reject(response);
      })
}