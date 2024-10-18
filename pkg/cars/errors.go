package cars

import "errors"

var ErrCarNameIsRequired = errors.New("car-name-is-required")
var ErrCarIdIsRequired = errors.New("car-id-is-required")
var ErrCarNotFound = errors.New("car-not-found")
