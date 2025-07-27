package cli

import "fmt"

type ErrCommandHandler struct {
	Code    int
	Message string
}

func (e *ErrCommandHandler) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func NewErrorCommandNotFound(command string) *ErrCommandHandler {
	return &ErrCommandHandler{
		Code:    404,
		Message: fmt.Sprintf("Command '%s' not found", command),
	}
}

func NewErrorNoCommandProvided() *ErrCommandHandler {
	return &ErrCommandHandler{
		Code:    400,
		Message: "No command provided",
	}
}