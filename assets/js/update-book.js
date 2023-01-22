const GET_BOOKS_URL = '/books/get-book'
const UPDATE_BOOK_URL = '/books/update'

if(!isLoggedIn() || !isAdmin()){
    location.pathname = '/login'
}

const getBook = () => {

    fetchData('GET', API_BASE+GET_BOOKS_URL+'?bookId='+bookId())
    .then((response) => {
        const data = response.data
        document.getElementById("title").value = data.title;
        document.getElementById("author").value = data.author;
        document.getElementById("quantity").value = data.quantity;
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });

}

document.getElementById('update-book-form')?.addEventListener('submit', (e) => {
    e.preventDefault();

    const data = new FormData(e.target);
    const request = Object.fromEntries(data.entries());
    request.quantity = parseInt(request.quantity)
    request.bookId = parseInt(bookId())

    fetchData('PATCH', API_BASE+UPDATE_BOOK_URL, JSON.stringify(request))
    .then((reponse) => {
       alert(reponse.message)
       getBook()
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });
});

const bookId = () => {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get('id')
}

getBook()
