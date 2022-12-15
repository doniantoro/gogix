package NewHelper

import (
	"encoding/json"
	"errors"
	"net/http"
)

// This Variable is initiate error that will use when mapping and give error response
var ErrPayload = errors.New("Invalid request payload")
var ErrBadRequest = errors.New("Error - 400 Bad Request")
var ErrUnAuthorized = errors.New("401 - Unauthorized")
var ErrPaymentRequired = errors.New("402 - Not Found")
var ErrForbidden = errors.New("403 - Forbidden")
var ErrNotFound = errors.New("404 - Not Found")
var ErrMethodNotAllow = errors.New("405 - Method Not Allowed")
var ErrNotAcceptable = errors.New("406 - Not Acceptable")
var ErrProxy = errors.New("407 - Proxy Authentication Required")
var ErrTimeOut = errors.New("408 - Request Time Out")
var ErrConflict = errors.New("409 - Error Conflict")
var ErrInternalServerError = errors.New("Error - 500 Internal Server Error")
var ErrNotImplemented = errors.New("Error - 501 Not Implemented")
var ErrBadGateway = errors.New("Error - 502 Bad Gateway")
var ErrServiceUnavailable = errors.New("Error - 503 Service Unavailable")
var ErrGatewayTimeOut = errors.New("Error - 504 Gateway Timeout")

// This Struct is used to to response json in http
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// MapStatusCode function has a function to mapping status code with parameter error
// it can use response json or for logging
func MapStatusCode(err error) int {

	if err == ErrBadRequest {
		return 400
	} else if err == ErrUnAuthorized {
		return 401
	} else if err == ErrPaymentRequired {
		return 402
	} else if err == ErrForbidden {
		return 403
	} else if err == ErrNotFound {
		return 404
	} else if err == ErrMethodNotAllow {
		return 405
	} else if err == ErrNotAcceptable {
		return 406
	} else if err == ErrProxy {
		return 407
	} else if err == ErrTimeOut {
		return 408
	} else if err == ErrConflict {
		return 409
	} else if err == ErrUnAuthorized {
		return 401
	} else if err == ErrInternalServerError {
		return 500
	} else if err == ErrNotImplemented {
		return 501
	} else if err == ErrBadGateway {
		return 502
	} else if err == ErrServiceUnavailable {
		return 503
	} else if err == ErrGatewayTimeOut {
		return 504
	} else {
		return 400
	}
}

// MapError function has a function to mapping status code with parameter status code
// it can use to mapping error json or logging
func MapError(code int) error {

	if code == 400 {
		return ErrBadRequest
	} else if code == 401 {
		return ErrUnAuthorized
	} else if code == 402 {
		return ErrPaymentRequired
	} else if code == 403 {
		return ErrForbidden
	} else if code == 404 {
		return ErrNotFound
	} else if code == 405 {
		return ErrMethodNotAllow
	} else if code == 406 {
		return ErrNotAcceptable
	} else if code == 407 {
		return ErrProxy
	} else if code == 408 {
		return ErrTimeOut
	} else if code == 409 {
		return ErrConflict
	} else if code == 401 {
		return ErrUnAuthorized
	} else if code == 500 {
		return ErrInternalServerError
	} else if code == 501 {
		return ErrNotImplemented
	} else if code == 502 {
		return ErrBadGateway
	} else if code == 503 {
		return ErrServiceUnavailable
	} else if code == 504 {
		return ErrGatewayTimeOut
	} else {
		return ErrBadRequest
	}

}

// ResponseErrorWithJSON is a error function to give response json in http with parameter responseWriter , error , and message
// If param error fill with param error that has initiate , it will mapping status code , and the message will automatic fill
//
//	if param error fill with param error that has initiate , but still wanna custome message, you can fill message
func ResponseErrorWithJSON(w http.ResponseWriter, err error, message string) {

	code := MapStatusCode(err)
	if message == "" {
		message = err.Error()
	}
	payload := Response{
		Status:  code,
		Message: message,
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ResponseErrorWithJSON is a success function to give response json in http with parameter responseWriter , code , message and payload
// code is response code when success , can use 200,201 or other
// Message is message success that wanna you give
// Payload is data will show in response http
func ResponseWithJSON(w http.ResponseWriter, code int, message string, payload interface{}) {

	payload = Response{
		Status:  code,
		Message: message,
		Data:    payload,
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
