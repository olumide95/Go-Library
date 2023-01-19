const ALL_BOOKS_URL = '/books/all'

fetchData('GET', API_BASE+ALL_BOOKS_URL)
.then((response) => {
    console.log(response)
})
.catch((response) => {
    console.log(response)
});


