package model

import "errors"

// NotFoundError represents an HTTP error defined with
// a code and a message
type HTTPError struct {
	Code    int64
	Message string
	Error   error
}

var (
	// ErrNotFound is an error returned when trying to access
	// an unexistent resource
	ErrNotFound = errors.New("not found")
	// ErrForbidden is an error returned when trying to access
	// a resources without permissions
	ErrForbidden = errors.New("forbidden")
	// ErrInternalServerError is an error triggered when the
	// server is not feeling well.
	ErrInternalServerError = errors.New("internal server error")

	// TODO(mb): move these in a dedicated place

	// CardBrandMastercard represents the Mastercard card brand
	CardBrandMastercard = "mastercard"
	// CardBrandVisa represents the Visa card brand
	CardBrandVisa = "visa"
	// CardBrandAMEX represents the American Express card brand
	CardBrandAMEX = "amex"
)
