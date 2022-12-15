package helper

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/doniantoro/gogix/domain"
)

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

func ResponseErrorWithJSON(err error, message string) {
	var w http.ResponseWriter

	code := MapStatusCode(err)
	if message == "" {
		message = err.Error()
	}
	payload := domain.Response{
		Status:  code,
		Message: message,
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponseWithJSON(code int, message string, payload interface{}) {
	var w http.ResponseWriter

	payload = domain.Response{
		Status:  code,
		Message: message,
		Data:    payload,
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
