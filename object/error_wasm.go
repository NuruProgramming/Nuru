//go:build wasm && js
package object

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	// msg := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "Kosa: ") // removes ANSI codes
	return "Kosa: " + e.Message
}
func (e *Error) Type() ObjectType { return ERROR_OBJ }
