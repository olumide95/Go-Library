package controller_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	controller "github.com/olumide95/go-library/api/controllers"
	"github.com/olumide95/go-library/api/middleware"
	"github.com/olumide95/go-library/api/util"
	"github.com/olumide95/go-library/models"
	"github.com/olumide95/go-library/repository"
	"github.com/olumide95/go-library/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("/books", func() {
	var server *gin.Engine
	var ur models.UserRepository
	var bc *controller.BookController
	var br models.BookRepository
	var blr models.BookLogRepository
	var brTx models.BookRepository
	var access_token string

	BeforeEach(func() {
		server = gin.Default()

		ur = repository.NewUserRepository(Db)
		br = repository.NewBookRepository(Db)
		blr = repository.NewBookLogRepository(Db)

		brTx = repository.NewBookRepository(Db)
		blrTx := repository.NewBookLogRepository(Db)
		bc = &controller.BookController{
			BookUsecase: usecase.NewbookUsecase(brTx, blrTx),
		}

		err := Db.AutoMigrate(&models.Book{}, &models.BookLog{}, &models.User{})
		Ω(err).To(Succeed())
	})

	When("/store endpoint is called", func() {

		BeforeEach(func() {
			server.POST("books/store", bc.StoreBooks)
		})

		Context("When a correct payload is sent", func() {

			It("Returns a status 201", func() {

				data := `[{"title": "Book 1", "author": "Author 1", "quantity" : 1 }]`

				req, _ := http.NewRequest("POST", "/books/store", strings.NewReader(data))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				server.ServeHTTP(resp, req)

				Expect(resp.Code).Should(Equal(http.StatusCreated))
			})
		})

		Context("When an incorrect payload is sent", func() {

			It("Returns a status 400", func() {

				data := `[{"title": "Book 1", "author": "Author 1", "quantity" : "1" }]`

				req, _ := http.NewRequest("POST", "/books/store", strings.NewReader(data))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				server.ServeHTTP(resp, req)

				Expect(resp.Code).Should(Equal(http.StatusBadRequest))
			})
		})

	})

	When("/borrow endpoint is called", func() {

		BeforeEach(func() {

			server.PATCH("books/borrow", middleware.NewAuthenticatedMiddlware(Db).Check,
				middleware.DBTransactionMiddleware(Db), bc.BorrowBook)
		})

		Context("authentication token not valid", func() {

			It("Returns a status 401", func() {

				data := `{"bookId": 1}`

				req, _ := http.NewRequest("PATCH", "/books/borrow", strings.NewReader(data))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "token")
				resp := httptest.NewRecorder()
				server.ServeHTTP(resp, req)

				Expect(resp.Code).Should(Equal(http.StatusUnauthorized))
			})
		})

		Context("authentication token is valid", func() {

			var book models.Book
			var user models.User

			BeforeEach(func() {

				book = models.Book{ID: 1, Title: "Test Title", Author: "Test Author", Quantity: 2}
				br.Create(&book)

				user = models.User{ID: 1, Name: "Test", Email: "borrow@email.com", Role: "User", Password: "password"}
				ur.Create(&user)

				os.Setenv("ACCESS_TOKEN_PRIVATE_KEY", "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlCUEFJQkFBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VXNjbGFFKzlaUUg5Q2VpOGIxcUVmCnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUUpCQUw4ZjRBMUlDSWEvQ2ZmdWR3TGMKNzRCdCtwOXg0TEZaZXMwdHdtV3Vha3hub3NaV0w4eVpSTUJpRmI4a25VL0hwb3piTnNxMmN1ZU9wKzVWdGRXNApiTlVDSVFENm9JdWxqcHdrZTFGY1VPaldnaXRQSjNnbFBma3NHVFBhdFYwYnJJVVI5d0loQVBOanJ1enB4ckhsCkUxRmJxeGtUNFZ5bWhCOU1HazU0Wk1jWnVjSmZOcjBUQWlFQWhML3UxOVZPdlVBWVd6Wjc3Y3JxMTdWSFBTcXoKUlhsZjd2TnJpdEg1ZGdjQ0lRRHR5QmFPdUxuNDlIOFIvZ2ZEZ1V1cjg3YWl5UHZ1YStxeEpXMzQrb0tFNXdJZwpQbG1KYXZsbW9jUG4rTkVRdGhLcTZuZFVYRGpXTTlTbktQQTVlUDZSUEs0PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
				os.Setenv("ACCESS_TOKEN_PUBLIC_KEY", "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VQpzY2xhRSs5WlFIOUNlaThiMXFFZnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==")
				access_token, _ = util.CreateToken(user.Email)
			})

			Context("incorrect request param is sent", func() {

				var resp *httptest.ResponseRecorder

				BeforeEach(func() {
					data := `{"bookI": 1}`

					req, _ := http.NewRequest("PATCH", "/books/borrow", strings.NewReader(data))
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("Authorization", "Bearer "+access_token)
					resp = httptest.NewRecorder()
					server.ServeHTTP(resp, req)
				})

				It("Returns a status 400", func() {
					Expect(resp.Code).Should(Equal(http.StatusBadRequest))
				})

				It("does not decrements the borrowed book", func() {
					updatedBook, _ := br.GetByID(book.ID)
					Expect(updatedBook.Quantity).Should(Equal(uint16(2)))
				})

				It("cdoes not create a book log record", func() {
					log, _ := blr.All()
					Expect(log).To(HaveLen(0))
				})
			})

			Context("correct request param is sent", func() {

				var resp *httptest.ResponseRecorder

				BeforeEach(func() {
					data := `{"bookId": 1}`

					req, _ := http.NewRequest("PATCH", "/books/borrow", strings.NewReader(data))
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("Authorization", "Bearer "+access_token)
					resp = httptest.NewRecorder()
					server.ServeHTTP(resp, req)
				})

				It("Returns a status 200", func() {
					Expect(resp.Code).Should(Equal(http.StatusOK))
				})

				It("Decrements the borrowed book record quantity by 1", func() {
					updatedBook, _ := br.GetByID(book.ID)
					Expect(updatedBook.Quantity).Should(Equal(uint16(1)))
				})

				It("creates a book log record", func() {
					log, _ := blr.All()
					Expect(log[0].BookId).To(Equal(book.ID))
					Expect(log[0].UserId).To(Equal(user.ID))
					Expect(log[0].ReturnedAt).To(BeNil())
					Expect(log).To(HaveLen(1))
				})
			})
		})

	})
})
