package disbot

import "errors"

var (
	ErrUnknownCommandKey    = errors.New("unknown command key")
	ErrUndefinedCommandType = errors.New("undefined command type")
)
