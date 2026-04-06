package evaluator

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/object"
)

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
