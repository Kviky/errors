package errors

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Kviky/errors/models"

	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-openapi/errors"
)

const (
	supportText = " Please try again later or contact support at info@kviky.com!"
)

// Types of instances
const (
	InstApp = "app"
	InstApi = "api"
	InstIdentity = "identity"
	InstClient = "client"
	InstDB = "database"
)

// List of error codes
const (
	unauthorized         = "Unauthorized"
	temporaryRedirect    = "Temporary Redirect"
	permanentRedirect    = "Permanent Redirect"
	badRequest           = "Bad Request"
	notImplemented       = "Not Implemented"
	forbidden            = "Forbidden"
	notFound             = "Not Found"
	gone                 = "Gone"
	preconditionFailed   = "Precondition Failed"
	unprocessableEntry   = "Unprocessable Entity"
	lengthRequired       = "Length Required"
	tooManyRequests      = "Too Many Requests"
	internalServerError  = "Internal Server Error"
	serviceUnavailable   = "Service Unavailable"
	methodNotAllowed     = "Method Not Allowed"
)

// List of 400 errors
const (
	BadRequest = "Bad request!"
	CharterHasListings = "Charter cannot be deleted!"
	CharterNotCreated = "Charter not created!"
	InvalidMsgFormat = "Invalid message format!"
	InvalidBodyParam = "Invalid body parameter!"
	InvalidHeaderParam = "Invalid header parameter!"
	InvalidQueryParam = "Invalid query parameter!"
	InvalidPathParam = "Invalid path parameter!"
	ListingNotCreated = "Listing not created!"
	LocationNotCreated = "Location not created!"
	MandatoryParamIncorrect = "Mandatory parameter incorrect!"
	MandatoryParamMissing = "Mandatory parameter missing!"
	NameAlreadyTaken = "Name is already taken!"
	PortAlreadyExists = "Port name exists already!"
	ReservationNotCreated = "Reservation not created!"
)

// List of 401 errors
const (
	InvalidAuthToken = "Invalid authorization token!"
	MissingAuthToken = "Missing authorization token!"
	UnauthorizedAccess = "Unauthorized access!"
)

// List of 403 errors
const (
	ForbiddenAction = "Forbidden action!"
	ForbiddenResource = "Forbidden resource!"
)

// list of 404 errors
const (
	CharterNotFound = "Charter not found!"
	ListingNotFound = "Listing not found!"
	LocationNotFound = "Location not found!"
	ResourceNotFound = "Resource not found!"
	UserNotFound = "User not found!"
	UsersNotFound = "Users not found!"
)

// list of 405 errors
const (
	MethodNotAllowed = "Method not allowed!"
)

// List of 500 errors
const (
	SystemFailure = "System failure!"
	UnspecifiedFailure   = "Unspecified failure!"
)

// CreateProblemDetails - Helper function to create ProblemDetails object
func CreateProblemDetails(errorName string) *models.ProblemDetails {
	problem := &models.ProblemDetails{}
	problem.Title = errorName
	problem.Type = "/"

	switch errorName{

	// 400 ERRORS
	case BadRequest:
		problem.Detail = "There was a problem with the request!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = instClient
	case CharterHasListings:
		problem.Detail = "Charter cannot be deleted, because it still has some active listings!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case CharterNotCreated:
		problem.Detail = "There was a problem to create charter profile!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case InvalidBodyParam:
		problem.Detail = "The HTTP request contains an unsupported body parameter!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case InvalidHeaderParam:
		problem.Detail = "The HTTP request contains an unsupported header parameter!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case InvalidQueryParam:
		problem.Detail = "The HTTP request contains an unsupported query parameter in the URI!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case InvalidPathParam:
		problem.Detail = "The HTTP request contains an unsupported path parameter in the URI!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case ListingNotCreated:
		problem.Detail = "There was a problem to create listing!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case LocationNotCreated:
		problem.Detail = "There was a problem to create location!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case MandatoryParamIncorrect:
		problem.Detail = "Mandatory parameter has semantically incorrect value!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case MandatoryParamMissing:
		problem.Detail = "Parameter which is defined as mandatory is missing!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case InvalidMsgFormat:
		problem.Detail = "The HTTP request has an invalid format!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case NameAlreadyTaken:
		problem.Detail = "Requested name is already taken! Please, specify another name."
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case PortAlreadyExists:
		problem.Detail = "Requested port/marina name already exists for this country and city!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient
	case ReservationNotCreated:
		problem.Detail = "There was a problem to create reservation!"
		problem.Status = 400
		problem.Code = badRequest
		problem.Instance = InstClient

	// 401 ERRORS
	case InvalidAuthToken:
		problem.Detail = "Authorization token is invalid!"
		problem.Status = 401
		problem.Code = unauthorized
		problem.Instance = InstClient
	case MissingAuthToken:
		problem.Detail = "Authorization token is missing!"
		problem.Status = 401
		problem.Code = unauthorized
		problem.Instance = InstClient
	case UnauthorizedAccess:
		problem.Detail = "The request doesn't have permissions to access resources!"
		problem.Status = 401
		problem.Code = unauthorized
		problem.Instance = InstApi

	// 403 ERRORS
	case ForbiddenAction:
		problem.Detail = "You don't have a permission to make this action!"
		problem.Status = 403
		problem.Code = forbidden
		problem.Instance = InstClient
	case ForbiddenResource:
		problem.Detail = "You don't have a permission to access this resource!"
		problem.Status = 403
		problem.Code = forbidden
		problem.Instance = InstClient

	// 404 ERRORS
	case CharterNotFound:
		problem.Detail = "The charter indicated in the request does not exist!"
		problem.Status = 404
		problem.Code = notFound
		problem.Instance = InstClient
	case ListingNotFound:
		problem.Detail = "The listing indicated in the request does not exist!"
		problem.Status = 404
		problem.Code = notFound
		problem.Instance = InstClient
	case LocationNotFound:
		problem.Detail = "The location indicated in the request does not exist!"
		problem.Status = 404
		problem.Code = notFound
		problem.Instance = InstClient
	case ResourceNotFound:
		problem.Detail = "Requested resource does not exist!"
		problem.Status = 404
		problem.Code = notFound
		problem.Instance = InstClient
	case UserNotFound:
		problem.Detail = "The user indicated in the request does not exist!"
		problem.Status = 404
		problem.Code = notFound
		problem.Instance = InstClient
	case UsersNotFound:
		problem.Detail = "Requested users does not exist!"
		problem.Status = 404
		problem.Code = notFound
		problem.Instance = InstClient

	// 405 ERRORS
	case MethodNotAllowed:
		problem.Detail = "Requested method is not allowed. Check the response header `Allow` for allowed methods!"
		problem.Status = 405
		problem.Code = methodNotAllowed
		problem.Instance = InstClient

	// 500 ERRORS
	case UnspecifiedFailure:
		problem.Detail = "The request is rejected due to unspecified reason at the system!"
		problem.Status = 500
		problem.Code = internalServerError
		problem.Instance = InstApi

	// DEFUALT ERROR 500
	default:
		problem.Title = SystemFailure
		problem.Detail = "We are sorry, but there is an internal problem with the application!" + supportText
		problem.Status = 500
		problem.Code = internalServerError
		problem.Instance = InstApi
	}

	return problem
}


func NewMissingParam(name string) *models.InvalidParam {
	return &models.InvalidParam{
		Param:  &name,
		Reason: "Param missing",
	}
}

func NewInvalidParam(name string) *models.InvalidParam {
	return &models.InvalidParam{
		Param:  &name,
		Reason: "Param has incorrect type",
	}
}

func writeResponse(problem *models.ProblemDetails, rw http.ResponseWriter) {

	rw.WriteHeader(int(problem.Status))
	data, _ := problem.MarshalBinary()
	_, _ = rw.Write(data)
}

func unknownError(rw http.ResponseWriter, r *http.Request, err error) {

	log.WithField("util", "errors").Errorf("Unknown error: %v", err.Error())

	problem := CreateProblemDetails(SystemFailure)
	problem.Type = 	r.RequestURI
	writeResponse(problem, rw)
}

func flattenComposite(errs *errors.CompositeError) *errors.CompositeError {
	var res []error
	for _, er := range errs.Errors {
		switch e := er.(type) {
		case *errors.CompositeError:
			if len(e.Errors) > 0 {
				flat := flattenComposite(e)
				if len(flat.Errors) > 0 {
					res = append(res, flat.Errors...)
				}
			}
		default:
			if e != nil {
				res = append(res, e)
			}
		}
	}
	return errors.CompositeValidationError(res...)
}

func errorAsJSON(err errors.Error) []byte {
	b, _ := json.Marshal(struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	}{err.Code(), err.Error()})
	return b
}

// ServeError the error handler interface implementation
func ServeError(rw http.ResponseWriter, r *http.Request, err error) {
	rw.Header().Set("Content-Type", "application/json")
	
	bodyProblem := CreateProblemDetails(InvalidBodyParam)
	bodyProblem.Type = r.RequestURI

	bodyMissingProblem := CreateProblemDetails(MandatoryParamMissing)
	bodyMissingProblem.Type = r.RequestURI

	queryProblem := CreateProblemDetails(InvalidQueryParam)
	queryProblem.Type = r.RequestURI

	pathProblem := CreateProblemDetails(InvalidPathParam)
	pathProblem.Type = r.RequestURI

	headerProblem := CreateProblemDetails(InvalidHeaderParam)
	headerProblem.Type = r.RequestURI

	problem := CreateProblemDetails(InvalidMsgFormat)
	problem.Type = r.RequestURI
	
	switch e := err.(type) {
	case *errors.CompositeError:
		// er := flattenComposite(e)
		// log.Printf("er: %v", er.Errors)

		for _, errItem := range e.Errors{
			switch valErr := errItem.(type){
			case *errors.Validation:
				invalidParam := &models.InvalidParam{
					Param:  &valErr.Name,
					Reason: valErr.Error(),
				}
				switch valErr.In{
				case "body":
					// log.Printf("request body issue: %+v", valErr)
					// log.Printf("valErr.code(): %v", valErr.Code())
					if valErr.Name == "body"{
						problem.InvalidParams = append(problem.InvalidParams, invalidParam)
					} else {
						// Filter custom openapi errors
						// More details - https://github.com/go-openapi/errors/blob/master/schema.go
						if valErr.Code() == 602 {
							bodyMissingProblem.InvalidParams = append(bodyMissingProblem.InvalidParams, invalidParam)
						} else {
							bodyProblem.InvalidParams = append(bodyProblem.InvalidParams, invalidParam)
						}
					}

				case "query":
					queryProblem.InvalidParams = append(queryProblem.InvalidParams, invalidParam)

				case "path":
					pathProblem.InvalidParams = append(pathProblem.InvalidParams, invalidParam)

				case "header":
					headerProblem.InvalidParams = append(headerProblem.InvalidParams, invalidParam)

				default:
					problem.InvalidParams = append(problem.InvalidParams, invalidParam)
				}

			case *errors.ParseError:
				invalidParam := &models.InvalidParam{
					Param:  &valErr.Name,
					Reason: valErr.Error(),
				}
				switch valErr.In{
				case "body":
					if valErr.Name == "body"{
						problem.InvalidParams = append(problem.InvalidParams, invalidParam)
					} else {
						// Filter custom openapi errors
						// More details - https://github.com/go-openapi/errors/blob/master/schema.go
						if valErr.Code() == 602 {
							bodyMissingProblem.InvalidParams = append(bodyMissingProblem.InvalidParams, invalidParam)
						} else {
							bodyProblem.InvalidParams = append(bodyProblem.InvalidParams, invalidParam)
						}
					}

				default:
					ServeError(rw, r, valErr)
					return
				}

			default:
				ServeError(rw, r, valErr)
				return
			}
		}

		// The queue of the returned problem details is important
		// first let return query problems
		// second missing parameters
		// and then the rest
		if len(bodyProblem.InvalidParams) > 0 {
			writeResponse(bodyProblem, rw)
		} else if len(bodyMissingProblem.InvalidParams) > 0 {
			writeResponse(bodyMissingProblem, rw)
		} else if len(queryProblem.InvalidParams) > 0 {
			writeResponse(queryProblem, rw)
		} else if len(pathProblem.InvalidParams) > 0 {
			writeResponse(pathProblem, rw)
		} else if len(headerProblem.InvalidParams) > 0 {
			writeResponse(headerProblem, rw)
		} else if len(problem.InvalidParams) > 0 {
			writeResponse(problem, rw)
		}else {
			ServeError(rw, r, nil)
		}

	case *errors.MethodNotAllowedError:
		rw.Header().Add("Allow", strings.Join(err.(*errors.MethodNotAllowedError).Allowed, ","))

		methodNotAllowedProblem := CreateProblemDetails(MethodNotAllowed)
		methodNotAllowedProblem.Type = r.RequestURI
		if r == nil || r.Method != http.MethodHead {
			writeResponse(methodNotAllowedProblem, rw)
		}

	// Default error handler
	case errors.Error:

		if e.Code() == 400 {
			badRequestProblem := CreateProblemDetails(BadRequest)
			badRequestProblem.Detail = fmt.Sprintf("%v %v", badRequestProblem.Detail, e.Error())
			badRequestProblem.Type = r.RequestURI
			writeResponse(badRequestProblem, rw)
			return
		}

		if e.Code() == 401 {
			notAuthorizedProblem := CreateProblemDetails(UnauthorizedAccess)
			notAuthorizedProblem.Type = r.RequestURI
			writeResponse(notAuthorizedProblem, rw)
			return
		}

		if e.Code() == 404 {
			notFoundProblem := CreateProblemDetails(ResourceNotFound)
			notFoundProblem.Type = r.RequestURI
			notFoundProblem.Detail = notFoundProblem.Detail + " " + e.Error() 
			writeResponse(notFoundProblem, rw)
			return
		}

		value := reflect.ValueOf(e)
		//log.Printf("error value: %v, error code: %v", value, e.Code())
		if value.Kind() == reflect.Ptr && value.IsNil() {
			unknownError(rw, r, err)
			return
		}
		unknownError(rw, r, err)

	case nil:
		unknownError(rw, r, err)
	default:
		unknownError(rw, r, err)
	}
}