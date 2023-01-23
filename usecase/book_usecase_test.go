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

	BeforeEach(func() {

		br := repository.NewBookRepository(Db)
		blr := repository.NewBookLogRepository(Db)
		bu = usecase.NewbookUsecase(br, blr)

		err := Db.AutoMigrate(&models.Book{})
		Ω(err).To(Succeed())

	})

	Context("#Create", func() {

		It("Creates a book record in the DB", func() {
			book := models.Book{Title: "Test Title", Author: "Test Author", Quantity: 2}
			err := bu.Create(&book)
			Ω(err).To(Succeed())
		})
	})

})
