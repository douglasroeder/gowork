package services

import (
	"testing"

	"github.com/douglasroeder/gowork/app"
	"github.com/douglasroeder/gowork/models"
	"github.com/khaiql/dbcleaner"
	"github.com/stretchr/testify/suite"
)

// Cleaner truncates table before/after usage
var Cleaner = dbcleaner.New()

var goWork = app.NewApp()

// TestSuite holds the testing structure
type TestSuite struct {
	suite.Suite
}

// SetupSuite prepares our app to be tested
func (suite *TestSuite) SetupSuite() {
	goWork.DB.AutoMigrate(&models.Category{})
	sqlite := NewSQLiteEngine(goWork.DB)

	Cleaner.SetEngine(sqlite)
}

// SetupTest acquire a lock to perform table operations
func (suite *TestSuite) SetupTest() {
	Cleaner.Acquire("categories")
}

// TearDownTest cleans up all testing tables
func (suite *TestSuite) TearDownTest() {
	Cleaner.Clean("categories")
}

// TestRunSuite triggers the execution of all specs
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
