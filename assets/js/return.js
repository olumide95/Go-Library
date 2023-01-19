const BORROWED_BOOKS_URL = '/books/borrowed-books'
const RETURN_BOOK_URL = '/books/return'

const getBorrowedBooks = () => {
    if(!isLoggedIn()){
        location.pathname = '/login'
        return
    }

    fetchData('GET', API_BASE+BORROWED_BOOKS_URL)
    .then((response) => {
        let tableRow = '';
    
        response.data.forEach((data, index)=>{
            console.log(data)
            tableRow += `<tr> <th scope="row">${index + 1}</th> <td>${data.Book.title}</td> <td>${data.BorrowedAt}</td> <td> <button data-logid="${data.id}" onclick="returnBook(this);" class="btn btn-primary" type="button" ${data.ReturnedAt ? 'disabled' : ''} > Return </button></td></tr>`
        })
        document.getElementById('books-borrowed').innerHTML = tableRow
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });

}

const returnBook = (e) => {

fetchData('PATCH', API_BASE+RETURN_BOOK_URL, JSON.stringify({logId: parseInt(e.dataset.logid)}))
    .then((response) => {
        console.log(response)
        getBorrowedBooks()
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });
}

getBorrowedBooks()
