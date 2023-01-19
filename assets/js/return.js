const BORROWED_BOOKS_URL = '/books/borrowed-books'
const RETURN_BOOK_URL = '/books/return'

const getBorrowedBooks = () => {
    fetchData('GET', API_BASE+BORROWED_BOOKS_URL)
    .then((response) => {
        let tableRow = '';
    
        response.data.forEach((data, index)=>{
            console.log(data)
            tableRow += `<tr> <th scope="row">${index + 1}</th> <td>${data.title}</td>  <td>${data.author}</td>  <td>${data.quantity}</td> <td> <button data-id="${data.id}" onclick="returnBook(this);" class="btn btn-primary" type="button" ${data.quantity < 1 ? 'disabled' : ''} > Return </button></td></tr>`
        })
        document.getElementById('books-borrowed').innerHTML = tableRow
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });

}

const returnBook = (e) => {

fetchData('PATCH', API_BASE+RETURN_BOOK_URL, JSON.stringify({bookId: parseInt(e.dataset.id)}))
    .then((response) => {
        console.log(response)
        getBorrowedBooks()
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });
}


getBorrowedBooks()
