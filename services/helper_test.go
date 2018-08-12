package services

import (
	"log"
	"testing"

	"github.com/douglasroeder/gowork/app"
	"github.com/douglasroeder/gowork/models"
	"github.com/khaiql/dbcleaner"
	"github.com/stretchr/testify/suite"
	testfixtures "gopkg.in/testfixtures.v2"
)

// Cleaner truncates table before/after usage
var Cleaner = dbcleaner.New()
var fixtures *testfixtures.Context
var goWork = app.NewApp()

// TestSuite holds the testing structure
type TestSuite struct {
	suite.Suite
}

// SetupSuite prepares our app to be tested
func (suite *TestSuite) SetupSuite() {
	goWork.DB.AutoMigrate(&models.Category{}, &models.Contact{})
	sqlite := NewSQLiteEngine(goWork.DB)

	var err error
	fixtures, err = testfixtures.NewFolder(goWork.DB.DB(), &testfixtures.SQLite{}, "../fixtures")
	if err != nil {
		log.Fatal(err)
	}

	Cleaner.SetEngine(sqlite)
}

// SetupTest acquire a lock to perform table operations
func (suite *TestSuite) SetupTest() {
	Cleaner.Acquire("categories")

	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}

// TearDownTest cleans up all testing tables
func (suite *TestSuite) TearDownTest() {
	Cleaner.Clean("categories")
}

// TestRunSuite triggers the execution of all specs
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
