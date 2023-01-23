package repository_test

import (
	"github.com/olumide95/go-library/models"
	"github.com/olumide95/go-library/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {
	var repo models.BookRepository
	BeforeEach(func() {
		repo = repository.NewBookRepository(Db)
		err := Db.AutoMigrate(&models.Book{})
		Ω(err).To(Succeed())
	})

	It("Creates a book record", func() {
		book := models.Book{Title: "Test Title", Author: "Test Author", Quantity: 2}
		err := repo.Create(&book)
		Ω(err).To(Succeed())
	})

	It("List all book records", func() {
		book := models.Book{Title: "Test Title 1", Author: "Test Author", Quantity: 2}
		repo.Create(&book)

		books, err := repo.All()
		Ω(err).To(Succeed())
		Ω(books).To(HaveLen(1))
		Ω(books[0].Title).To(Equal("Test Title 1"))
	})

	It("deletes a book record", func() {
		book := models.Book{Title: "Test Title 1", Author: "Test Author", Quantity: 2}
		repo.Create(&book)

		l, err := repo.Delete([]uint{1})
		Ω(err).To(Succeed())
		Ω(l).To(BeEquivalentTo(1))
	})
})
