const STORE_BOOK_URL = '/books/store'

if(!isLoggedIn() || !isAdmin()){
    location.pathname = '/login'
}

const addRow = () => {
    const rowId = 'row' + Date.now()
    const row = document.createElement("div");
    row.setAttribute("class", "row pb-2");
    row.setAttribute("id",  rowId);
    row.innerHTML  = `<div class="col-4"> <div class="form-group"> <input type="text" class="form-control" name="title" value="" placeholder="Tile"> </div> </div> <div class="col-3"> <div class="form-group"> <input type="text" class="form-control" name="author" value="" placeholder="Author"> </div> </div> <div class="col-3"> <div class="form-group"> <input type="number" class="form-control" name="quantity" value="" placeholder="Quantity"> </div> </div> <div class="col-1"> <button class="btn btn-danger" type="button"  onclick="removeRow('${rowId}');"> - </button> </div>`
    document.getElementById('book-row').append(row)
}

const removeRow = (rowId) => {
    document.getElementById(rowId).remove();
}

document.getElementById('add-books-form')?.addEventListener('submit', (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);

    const titles = formData.getAll('title')
    const authors = formData.getAll('author')
    const quantities = formData.getAll('quantity')

    let data = []
    titles.forEach((val, index) => {
        data.push({title: val, author: authors[index], quantity: parseInt(quantities[index])})
    })

    fetchData('POST', API_BASE+STORE_BOOK_URL, JSON.stringify(data))
    .then((reponse) => {
       alert(reponse.message)
       location.reload();
    })
    .catch((response) => {
        response.json().then((r) => { alert(r.message) })
    });
});