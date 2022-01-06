package errs

import (
	"net/http"
	"testing"
)

func assertError(errMessage string, errCode int, appErr AppError, t *testing.T) {
	if appErr.Message != errMessage {
		t.Error("Invalid error message")
	}

	if appErr.Code != errCode {
		t.Error("Invalid error code")
	}
}

func Test_should_return_not_found_error(t *testing.T) {
	errMessage := "Customer not found"
	appErr := NewNotFoundError(errMessage)

	assertError(errMessage, http.StatusNotFound, *appErr, t)
}

func Test_should_return_unexpected_error(t *testing.T) {
	errMessage := "Something went wrong"
	appErr := NewUnexpectedError(errMessage)

	assertError(errMessage, http.StatusInternalServerError, *appErr, t)
}
