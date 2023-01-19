const ALL_BOOKS_URL = '/books/all'
const BORROW_BOOK_URL = '/books/borrow'

const getAvailableBooks = () => {
    if(!isLoggedIn()){
        location.pathname = '/login'
        return
    }

    fetchData('GET', API_BASE+ALL_BOOKS_URL)
    .then((response) => {
        let tableRow = '';
    
        response.data.forEach((data, index)=>{
            tableRow += `<tr> <th scope="row">${index + 1}</th> <td>${data.title}</td>  <td>${data.author}</td>  <td>${data.quantity}</td> <td> <button data-id="${data.id}" onclick="borrowBook(this);" class="btn btn-primary" type="button" ${data.quantity < 1 ? 'disabled' : ''} > Borrow </button></td></tr>`
        })
        document.getElementById('books-avaialbe').innerHTML = tableRow
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });

}

const borrowBook = (e) => {

fetchData('PATCH', API_BASE+BORROW_BOOK_URL, JSON.stringify({bookId: parseInt(e.dataset.id)}))
    .then((response) => {
        console.log(response)
        getAvailableBooks()
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });
}

getAvailableBooks()
