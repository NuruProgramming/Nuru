package object

import "fmt"

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	msg := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "Kosa: ")
	return msg + e.Message
}
func (e *Error) Type() ObjectType { return ERROR_OBJ }
