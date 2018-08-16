package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult_GenericResponse(t *testing.T) {
	anyModel := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   123,
		Name: "test",
	}
	error := ""
	result := NewResult(200, anyModel, error)

	outputJSON, _ := json.Marshal(result)
	assert.Equal(t, `{"status_code":200,"error":"","payload":{"id":123,"name":"test"}}`, string(outputJSON))
}
