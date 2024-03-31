package builtins

import (
	"fmt"
	"github.com/AvicennaJr/Nuru/object"
	"io"
	"os"
	"bufio"
)

func Aina(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("Samahani, tunahitaji hoja 1, wewe umeweka %d", len(args))
	}

	return &object.String{Value: string(args[0].Type())}
}

func Fungua(args ...object.Object) object.Object {

	if len(args) != 1 {
		return newError("Samahani, tunahitaji hoja 1, wewe umeweka %d", len(args))
	}
	filename := args[0].(*object.String).Value

	file, err := os.ReadFile(filename)
	if err != nil {
		return &object.Error{Message: "Tumeshindwa kusoma faili"}
	}
	return &object.File{Filename: filename, Content: string(file)}
}

func Jaza(args ...object.Object) object.Object {

	if len(args) > 1 {
		return newError("Samahani, kiendesha hiki kinapokea hoja 0 au 1, wewe umeweka %d", len(args))
	}

	if len(args) > 0 && args[0].Type() != object.STRING_OBJ {
		return newError(fmt.Sprintf(`Tafadhali tumia alama ya nukuu: "%s"`, args[0].Inspect()))
	}
	if len(args) == 1 {
		prompt := args[0].(*object.String).Value
		fmt.Fprint(os.Stdout, prompt)
	}

	buffer := bufio.NewReader(os.Stdin)

	line, _, err := buffer.ReadLine()
	if err != nil && err != io.EOF {
		return newError("Nimeshindwa kusoma uliyo yajaza")
	}

	return &object.String{Value: string(line)}
}
