package models

import (
	"encoding/json"
	"errors"
	"testing"

	cer "github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	. "github.com/stretchr/testify/assert"
)

func TestInvalidParam_MarshalBinary(t *testing.T) {
	param := InvalidParam{}
	binary, err := param.MarshalBinary()
	NoError(t, err)
	Equal(t, `{"param":null}`, string(binary))

	email := "email"
	param.Param = &email
	param.Reason = "invalid email"

	binary, err = param.MarshalBinary()
	NoError(t, err)
	Equal(t, `{"param":"email","reason":"invalid email"}`, string(binary))
}

func TestInvalidParam_UnmarshalBinary(t *testing.T) {
	param := InvalidParam{}
	var s *string

	bytes, err := json.Marshal(&param)
	NoError(t, err)

	err = param.UnmarshalBinary(bytes)
	NoError(t, err)
	Equal(t, s, param.Param)

	email := "email"
	param.Param = &email
	param.Reason = "invalid param"

	bytes, err = json.Marshal(&param)
	NoError(t, err)

	param.Param = nil
	param.Reason = ""

	err = param.UnmarshalBinary(bytes)
	NoError(t, err)
	Equal(t, "invalid param", param.Reason)
	Equal(t, email, *param.Param)
}

func TestInvalidParam_Validate(t *testing.T) {
	param := InvalidParam{}
	registry := strfmt.NewFormats()
	cerr := cer.CompositeValidationError(errors.New("param in body is required"))

	err := param.Validate(registry)
	Error(t, err)
	EqualError(t, err, cerr.Error())

	s := "email"
	param.Param = &s
	err = param.Validate(registry)
	NoError(t, err)
}

func TestInvalidParam_validateParam(t *testing.T) {
	param := InvalidParam{}
	registry := strfmt.NewFormats()

	err := param.validateParam(registry)
	Error(t, err)
	EqualError(t, err, "param in body is required")

	s := "email"
	param.Param = &s
	err = param.validateParam(registry)
	NoError(t, err)
}
