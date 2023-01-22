const API_BASE = 'http://localhost:8080' 
const SIGNIN_URL = '/login'
const SIGNUP_URL = '/signup'
const ADMIN_ROLE = 'Admin'
const USER_ROLE = 'User'

const isLoggedIn = () => {
    return !!localStorage.getItem('access_token');
}

const isAdmin = () => {
    const user = JSON.parse(localStorage.getItem('user'))
    return user?.role == ADMIN_ROLE;
}

const isUser = () => {
    const user = JSON.parse(localStorage.getItem('user'))
    return user?.role == USER_ROLE;
}

const logOut = () => {
    localStorage.removeItem('user')
    localStorage.removeItem('access_token')
    location.pathname = '/'
}

window.onload = (event) => {
    
    const access_token = localStorage.getItem('access_token')
    if(access_token){
        document.querySelectorAll('.no-auth').forEach(function(el) {
            el.style.display = 'none';
         });

        const user = JSON.parse(localStorage.getItem('user'))
       
        if(user.role == USER_ROLE){

            document.querySelectorAll('.admin-auth').forEach(function(el) {
                el.style.display = 'none';
            });

        }

        if(user.role == ADMIN_ROLE){

            document.querySelectorAll('.user-auth').forEach(function(el) {
                el.style.display = 'none';
            });

        }
        return;
    }

    document.querySelectorAll('.auth').forEach(function(el) {
        el.style.display = 'none';
     });
};

document.getElementById('signin-form')?.addEventListener('submit', (e) => {
    e.preventDefault();

    const data = new FormData(e.target);
    const request = Object.fromEntries(data.entries());

    fetchData('POST', API_BASE+SIGNIN_URL, JSON.stringify(request))
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

    fetchData('POST', API_BASE+SIGNUP_URL, JSON.stringify(request))
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

    const bearer_token = "Bearer " + localStorage.getItem('access_token')

    return fetch(url, {
        method: method,
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin', 
        headers: {
            Authorization: bearer_token, 
            'Content-Type': 'application/json'},
        body: data
      }).then((response) => {
        
        if (response.ok) {
          return response.json();
        }

        if (response.status == 401) {
            location.pathname = 'login'
            return
        }

        return Promise.reject(response);
      })
}