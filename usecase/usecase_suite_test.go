package usecase_test

import (
	"testing"

	"github.com/olumide95/go-library/bootstrap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecase Suite")
}

var Db *gorm.DB
var cleanupDocker func()

var _ = BeforeSuite(func() {
	// setup *gorm.Db with docker
	Db, cleanupDocker = bootstrap.SetupGormWithDocker()
})

var _ = AfterSuite(func() {
	// cleanup resource
	cleanupDocker()
})

var _ = BeforeEach(func() {
	// clear db tables before each test
	err := Db.Exec(`DROP SCHEMA public CASCADE;CREATE SCHEMA public;`).Error
	Ω(err).To(Succeed())
})
