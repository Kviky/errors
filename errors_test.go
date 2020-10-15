package errors

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	cer "github.com/go-openapi/errors"
	. "github.com/stretchr/testify/assert"

	"github.com/Kviky/errors/models"
)

func TestCreateProblemDetails(t *testing.T) {
	details := CreateProblemDetails("")
	EqualValues(t, http.StatusInternalServerError, details.Status)
	EqualValues(t, internalServerError, details.Code)

	details = CreateProblemDetails(BadRequest)
	EqualValues(t, http.StatusBadRequest, details.Status)
	EqualValues(t, badRequest, details.Code)

	details = CreateProblemDetails(NameAlreadyTaken)
	EqualValues(t, http.StatusBadRequest, details.Status)
	EqualValues(t, badRequest, details.Code)
	EqualValues(t, "Requested name is already taken! Please, specify another name.", details.Detail)
}

func TestNewInvalidParam(t *testing.T) {
	param := NewInvalidParam("email")
	EqualValues(t, "email", *param.Param)
	EqualValues(t, "Param has incorrect type", param.Reason)
}

func TestNewMissingParam(t *testing.T) {
	param := NewMissingParam("email")
	EqualValues(t, "email", *param.Param)
	EqualValues(t, "Param missing", param.Reason)
}

func TestServeError(t *testing.T) {
	emptyError := func(w http.ResponseWriter, r *http.Request) {
		ServeError(w, r, errors.New(""))
	}

	handler := http.HandlerFunc(emptyError)
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(rr, r)
	EqualValues(t, http.StatusInternalServerError, rr.Code)

	validationError := func(w http.ResponseWriter, r *http.Request) {
		invalidType := cer.InvalidType("email", "body", "string", "")
		err := cer.CompositeValidationError(invalidType)
		ServeError(w, r, err)
	}

	rr = httptest.NewRecorder()
	handler = validationError
	handler.ServeHTTP(rr, r)
	EqualValues(t, http.StatusBadRequest, rr.Code)

	methodError := func(w http.ResponseWriter, r *http.Request) {
		allowed := cer.MethodNotAllowed("test", nil)
		err := cer.CompositeValidationError(allowed)
		ServeError(w, r, err)
	}

	rr = httptest.NewRecorder()
	handler = methodError
	handler.ServeHTTP(rr, r)
	EqualValues(t, http.StatusMethodNotAllowed, rr.Code)
}

func Test_errorAsJSON(t *testing.T) {
	err := cer.New(http.StatusBadRequest, "invalid param")
	bytes := errorAsJSON(err)
	Equal(t, `{"code":400,"message":"invalid param"}`, string(bytes))
}

func Test_flattenComposite(t *testing.T) {
	invalidType := cer.InvalidType("email", "body", "string", "")

	firstErr := cer.CompositeValidationError(invalidType)
	err := cer.CompositeValidationError(firstErr)
	err = cer.CompositeValidationError(err)

	NotEqual(t, firstErr.Error(), err.Error())

	composite := flattenComposite(err)
	Equal(t, firstErr.Error(), composite.Error())
}

func Test_unknownError(t *testing.T) {
	emptyError := func(w http.ResponseWriter, r *http.Request) {
		unknownError(w, r, errors.New(""))
	}

	handler := http.HandlerFunc(emptyError)
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(rr, r)
	EqualValues(t, http.StatusInternalServerError, rr.Code)
}

func Test_writeResponse(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {
		writeResponse(&models.ProblemDetails{Status: http.StatusInternalServerError}, w)
	}

	handler := http.HandlerFunc(h)
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(rr, r)
	EqualValues(t, http.StatusInternalServerError, rr.Code)
	Equal(t, `{"status":500}`, rr.Body.String())
}
