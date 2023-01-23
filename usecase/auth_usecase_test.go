package usecase_test

import (
	"github.com/olumide95/go-library/domain"
	"github.com/olumide95/go-library/models"
	"github.com/olumide95/go-library/repository"
	"github.com/olumide95/go-library/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthUsecase", func() {
	var au domain.AuthUsecase

	BeforeEach(func() {

		ur := repository.NewUserRepository(Db)
		au = usecase.NewauthUsecase(ur)

		err := Db.AutoMigrate(&models.User{})
		Ω(err).To(Succeed())

	})

	Context("#Create", func() {

		It("Creates a user record in the DB", func() {
			user := models.User{Name: "Test", Email: "test@email.com", Role: "User", Password: "password"}
			err := au.CreateUser(&user)
			Ω(err).To(Succeed())
		})

		It("does not create a user record in the DB when there is an existing user with the same email", func() {
			user1 := models.User{Name: "Test", Email: "test@email.com", Role: "User", Password: "password"}
			au.CreateUser(&user1)

			user2 := models.User{Name: "Test", Email: "test@email.com", Role: "User", Password: "password"}
			err := au.CreateUser(&user2)
			Ω(err).NotTo(Succeed())
		})
	})

})
