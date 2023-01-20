const ALL_BOOKS_URL = '/books/all'
const DELETE_BOOK_URL = '/books/delete'
const UPDATE_BOOK_URL = '/books/update'

const getAvailableBooks = () => {
    if(!isLoggedIn() && isAdmin()){
        location.pathname = '/login'
        return
    }

    fetchData('GET', API_BASE+ALL_BOOKS_URL)
    .then((response) => {
        let tableRow = '';
    
        response.data.forEach((data, index)=>{
            tableRow += `<tr> <th scope="row"> <input class="form-check-input books-check" type="checkbox" data-id="${data.id}"> </th> <td>${index + 1}</td> <td>${data.title}</td>  <td>${data.author}</td>  <td> <button data-id="${data.id}" data-quantity="${data.quantity}" data-title="${data.title}" data-author="${data.author}" ${data.quantity < 1 ? 'disabled' : ''} onClick="decrementQuantity(this)" class="btn btn-danger btn-sm" type="button"> - </button> ${data.quantity} <button data-id="${data.id}" data-quantity="${data.quantity}" data-title="${data.title}" data-author="${data.author}" onClick="incrementQuantity(this)" class="btn btn-primary btn-sm" type="button"> + </button> </td> <td> <a href="/admin/books/update/${data.id}"> <button data-id="${data.id}" class="btn btn-primary" type="button"> Update </button> </a></td></tr>`
        })
        document.getElementById('books-avaialbe').innerHTML = tableRow
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });

}

const incrementQuantity = (e) => {
    const quantity = parseInt(e.dataset.quantity) + 1
    const id = parseInt(e.dataset.id)

    const data = {bookId: id, quantity: quantity, title: e.dataset.title, author:  e.dataset.author}
    console.log(data)
    updateBook(data)
}

const decrementQuantity = (e) => {
    const quantity = parseInt(e.dataset.quantity) - 1
    const id = parseInt(e.dataset.id)

    const data = {bookId: id, quantity: quantity, title: e.dataset.title, author:  e.dataset.author }
    updateBook(data)
}

const updateBook = (data) => {
    fetchData('PATCH', API_BASE+UPDATE_BOOK_URL, JSON.stringify(data))
    .then((response) => {
        getAvailableBooks()
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });
}

const deleteSelected = () => {
    let ids = []

    $('.books-check:checkbox:checked').each(function (){
            ids.push({id : $(this).data('id')})
    })

    if(ids.length == 0){
        alert('Please select a book');
        return
    }

    fetchData('DELETE', API_BASE+DELETE_BOOK_URL, JSON.stringify(ids))
    .then((response) => {
        getAvailableBooks()
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });
}

const toggelSelectAll = (e) => {
    $('.books-check').prop('checked', e.checked);
}

getAvailableBooks()
