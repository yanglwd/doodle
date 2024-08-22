package shared

import "errors"

var (
	ErrInvalidProtocol = errors.New("invalid protocol")
	ErrInvalidListener = errors.New("must listen first")
)
