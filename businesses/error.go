package businesses

import "errors"

var (
	ErrInternalServer  = errors.New("Something Gone Wrong,Please Contact Administrator")
	ErrNotFoundProduct = errors.New("Product Does'nt Exist")
)
