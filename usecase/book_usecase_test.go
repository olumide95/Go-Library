package usecase_test

import (
	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/models"
	"github.com/olumide95/go-library/repository"
	"github.com/olumide95/go-library/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BookUsecase", func() {
	var bu domain.BookUsecase
	var ur models.UserRepository
	var blr models.BookLogRepository

	BeforeEach(func() {

		ur = repository.NewUserRepository(Db)
		br := repository.NewBookRepository(Db)
		blr = repository.NewBookLogRepository(Db)
		bu = usecase.NewbookUsecase(br, blr)

		err := Db.AutoMigrate(&models.Book{}, &models.BookLog{}, &models.User{})
		Ω(err).To(Succeed())
	})

	Context("#CreateBulk", func() {

		It("stores book records in bulk in the DB", func() {
			books := []models.Book{{Title: "Test Title", Author: "Test Author", Quantity: 2}}
			err := bu.CreateBulk(&books)
			Ω(err).To(Succeed())
		})
	})

	Context("#BorrowBook", func() {

		BeforeEach(func() {
			book := models.Book{Title: "Test Title", Author: "Test Author", Quantity: 2}
			bu.Create(&book)

			user1 := models.User{ID: 1, Name: "Test", Email: "borrow@email.com", Role: "User", Password: "password"}
			ur.Create(&user1)

			books, _ := bu.AllBooks()
			bu.BorrowBook(books[0].ID, 1)
		})

		It("reduces the borrowed book quantity by one", func() {
			books, _ := bu.AllBooks()
			Expect(books[0].Quantity).To(Equal(uint16(1)))
		})

		It("creates a book log record with nil returned at column", func() {
			logs, _ := blr.All()

			Ω(logs).To(HaveLen(1))
			Ω(logs[0].ReturnedAt).To(BeNil())
		})
	})

})
