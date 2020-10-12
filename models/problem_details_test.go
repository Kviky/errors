package models

import (
	"encoding/json"
	"net/http"
	"testing"

	cer "github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	. "github.com/stretchr/testify/assert"
)

func TestProblemDetails_MarshalBinary(t *testing.T) {
	details := ProblemDetails{}
	binary, err := details.MarshalBinary()
	NoError(t, err)
	Equal(t, `{}`, string(binary))

	details.Status = http.StatusOK
	details.Code = "OK"
	binary, err = details.MarshalBinary()
	NoError(t, err)
	Equal(t, `{"code":"OK","status":200}`, string(binary))
}

func TestProblemDetails_UnmarshalBinary(t *testing.T) {
	details := ProblemDetails{}

	bytes, err := json.Marshal(&details)
	NoError(t, err)

	err = details.UnmarshalBinary(bytes)
	NoError(t, err)
	Equal(t, ProblemDetails{}, details)

	details.Code = "OK"
	details.Status = http.StatusOK

	bytes, err = json.Marshal(&details)
	NoError(t, err)

	err = details.UnmarshalBinary(bytes)
	NoError(t, err)
	Equal(t, ProblemDetails{
		Status: http.StatusOK,
		Code:   "OK",
	}, details)
}

func TestProblemDetails_Validate(t *testing.T) {
	details := ProblemDetails{}
	registry := strfmt.NewFormats()

	err := details.Validate(registry)
	NoError(t, err)

	var s *string
	params := make([]*InvalidParam, 0, 1)

	param := &InvalidParam{
		Param:  s,
		Reason: "empty",
	}
	details.InvalidParams = append(params, param)
	err = details.Validate(registry)
	Error(t, err)

	cerr := cer.CompositeValidationError(param.Validate(registry))
	Error(t, err)

	EqualError(t, err, cerr.Error())
}

func TestProblemDetails_validateInvalidParams(t *testing.T) {
	details := ProblemDetails{}
	registry := strfmt.NewFormats()

	var s *string
	params := make([]*InvalidParam, 0, 1)

	param := &InvalidParam{
		Param:  s,
		Reason: "empty",
	}

	details.InvalidParams = append(params, param)
	err := details.validateInvalidParams(registry)
	Error(t, err)

	errv := param.Validate(registry)
	Error(t, errv)

	EqualError(t, err, errv.Error())
}
