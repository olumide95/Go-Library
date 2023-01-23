package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	controller "github.com/olumide95/go-library/api/controllers"
	"github.com/olumide95/go-library/models"
	"github.com/olumide95/go-library/repository"
	"github.com/olumide95/go-library/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("/book/store", func() {
	var server *gin.Engine
	var bc *controller.BookController

	BeforeEach(func() {
		server = gin.Default()

		br := repository.NewBookRepository(Db)
		blr := repository.NewBookLogRepository(Db)
		bc = &controller.BookController{
			BookUsecase: usecase.NewbookUsecase(br, blr),
		}

		err := Db.AutoMigrate(&models.Book{})
		Ω(err).To(Succeed())

	})

	When("the /store endpoint is called", func() {

		BeforeEach(func() {
			server.POST("books/store", bc.StoreBooks)
		})

		Context("When a correct payload is sent", func() {

			It("Returns a status 201", func() {

				data := `[{ "title": "Book 1", 	"author: "Author 1", quantity: 1 }]`

				req, _ := http.NewRequest("POST", "/books/store", strings.NewReader(data))
				resp := httptest.NewRecorder()
				server.ServeHTTP(resp, req)

				Expect(resp.Code).Should(Equal(http.StatusCreated))
			})
		})

	})

	// Context("When a PATC", func() {

	// 	It("Returns the empty path", func() {
	// 		resp, err := http.Get(server.URL() + "/hello")
	// 		Expect(err).ShouldNot(HaveOccurred())
	// 		Expect(resp.StatusCode).Should(Equal(http.StatusOK))
	// 		body, err := ioutil.ReadAll(resp.Body)
	// 		resp.Body.Close()
	// 		Expect(err).ShouldNot(HaveOccurred())
	// 		Expect(string(body)).To(Equal(msg + "hello!"))
	// 	})
	// })

	// Context("When get request is sent to read path but there is no file", func() {

	// 	It("Returns internal server error", func() {
	// 		resp, err := http.Get(server.URL() + "/read")
	// 		Expect(err).ShouldNot(HaveOccurred())
	// 		Expect(resp.StatusCode).Should(Equal(http.StatusInternalServerError))
	// 		body, err := ioutil.ReadAll(resp.Body)
	// 		resp.Body.Close()
	// 		Expect(err).ShouldNot(HaveOccurred())
	// 		Expect(string(body)).To(Equal("open data.txt: no such file or directory\n"))
	// 	})
	// })

	// Context("When get request is sent to read path but file exists", func() {

	// 	It("Reads data from file successfully", func() {
	// 		resp, err := http.Get(server.URL() + "/read")
	// 		Expect(err).ShouldNot(HaveOccurred())
	// 		Expect(resp.StatusCode).Should(Equal(http.StatusOK))
	// 		body, err := ioutil.ReadAll(resp.Body)
	// 		resp.Body.Close()
	// 		Expect(err).ShouldNot(HaveOccurred())
	// 		Expect(string(body)).To(Equal("Content in file is...\r\nHi there!"))
	// 	})
	// })
})
