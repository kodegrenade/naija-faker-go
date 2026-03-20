package errors

import "fmt"

type ErrorCode string

const (
	ErrInvalidLanguage ErrorCode = "INVALID_LANGUAGE"
	ErrInvalidGender   ErrorCode = "INVALID_GENDER"
	ErrInvalidNetwork  ErrorCode = "INVALID_NETWORK"
	ErrInvalidState    ErrorCode = "INVALID_STATE"
	ErrInvalidBank     ErrorCode = "INVALID_BANK"
	ErrInvalidLevel    ErrorCode = "INVALID_LEVEL"
	ErrInvalidType     ErrorCode = "INVALID_TYPE"
	ErrInvalidParam    ErrorCode = "INVALID_PARAM"
)

type NaijaFakerError struct {
	Code    ErrorCode
	Message string
}

func (e *NaijaFakerError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func New(code ErrorCode, message string) *NaijaFakerError {
	return &NaijaFakerError{Code: code, Message: message}
}
