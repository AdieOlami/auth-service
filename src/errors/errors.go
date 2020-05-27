package errors

import (
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   true,
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   true,
	}
}

func NewInteralServerError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   true,
	}
}

func NewInternalServerError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   true,
	}
}

// func NewRestErrorFromBytes(message string) *Error {
// 	return &Error{
// 		Message: message,
// 		Status:  http.StatusNotFound,
// 		Error:   true,
// 	}
// }

func NewRestErrorFromBytes(bytes []byte) (*Error, error) {
	// var apiErr Error{}
	// if err := json.Unmarshal(bytes, &apiErr); err != nil {
	// 	return nil, errors.New("invalid json")
	// }
	return nil, nil
}

func New(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   true,
	}
}
