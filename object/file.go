package object

import (
	"os"
)

type File struct {
	Filename string
	Content  string
}

func (f *File) Type() ObjectType { return FILE_OBJ }
func (f *File) Inspect() string  { return f.Filename }
func (f *File) Method(method string, args []Object) Object {
	switch method {
	case "soma":
		return f.read(args)
	case "andika":
		return f.write(args)
	case "ongeza":
		return f.append(args)
	}
	return nil
}

func (f *File) read(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	return &String{Value: f.Content}
}

func (f *File) write(args []Object) Object {
	if len(args) != 1 {
		return newError("Samahani, tunahitaji Hoja 1, wewe umeweka %d", len(args))
	}
	content, ok := args[0].(*String)
	if !ok {
		return newError("Samahani, hoja lazima iwe Tungo")
	}
	err := os.WriteFile(f.Filename, []byte(content.Value), 0644)
	if err != nil {
		return newError("Hitilafu katika kuandika faili: %s", err.Error())
	}
	f.Content = content.Value
	return &Boolean{Value: true}
}

func (f *File) append(args []Object) Object {
	if len(args) != 1 {
		return newError("Samahani, tunahitaji Hoja 1, wewe umeweka %d", len(args))
	}
	content, ok := args[0].(*String)
	if !ok {
		return newError("Samahani, hoja lazima iwe Tungo")
	}
	file, err := os.OpenFile(f.Filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return newError("Hitilafu katika kufungua faili: %s", err.Error())
	}
	defer file.Close()
	_, err = file.WriteString(content.Value)
	if err != nil {
		return newError("Hitilafu katika kuongeza kwa faili: %s", err.Error())
	}
	f.Content += content.Value
	return &Boolean{Value: true}
}
