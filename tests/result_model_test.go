package tests

import (
	"encoding/json"

	"github.com/douglasroeder/gowork/models"
)

func (suite *TestSuite) TestResultModel_GenericResponse() {
	anyModel := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   123,
		Name: "test",
	}
	errors := []string{}
	result := models.NewResult(200, anyModel, errors)

	outputJSON, _ := json.Marshal(result)
	suite.Equal(`{"status_code":200,"errors":[],"payload":{"id":123,"name":"test"}}`, string(outputJSON))
}
