package object

import (
	"bufio"
	"io"
	"os"
)

type File struct {
	Filename string
	Reader   *bufio.Reader // To read the file
	Writer   *bufio.Writer // To write on the file
	Handle   *os.File      // To handle the actual file (open, close etc)
}

func (f *File) Type() ObjectType { return FILE_OBJ }
func (f *File) Inspect() string  { return f.Filename }
func (f *File) Method(method string, args []Object) Object {
	switch method {
	case "soma":
		return f.read(args)
	case "funga":
		return f.close(args)
	}
	return nil
}

func (f *File) read(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	if f.Reader == nil {
		return nil
	}
	txt, _ := io.ReadAll(f.Reader)
	return &String{Value: string(txt)}
}

func (f *File) close(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	_ = f.Handle.Close()
	return &Null{}
}
