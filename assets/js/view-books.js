const ALL_BOOKS_URL = '/books/all'
const DELETE_BOOK_URL = '/books/delete'

const getAvailableBooks = () => {
    if(!isLoggedIn() && isAdmin()){
        location.pathname = '/login'
        return
    }

    fetchData('GET', API_BASE+ALL_BOOKS_URL)
    .then((response) => {
        let tableRow = '';
    
        response.data.forEach((data, index)=>{
            tableRow += `<tr> <th scope="row"> <input class="form-check-input books-check" type="checkbox" data-id="${data.id}"> </th> <td>${index + 1}</td> <td>${data.title}</td>  <td>${data.author}</td>  <td>${data.quantity}</td> <td> <a href="/admin/books/update/${data.id}"> <button data-id="${data.id}" class="btn btn-primary" type="button"> Update </button> </a></td></tr>`
        })
        document.getElementById('books-avaialbe').innerHTML = tableRow
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
