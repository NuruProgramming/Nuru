package object

type File struct {
	Filename string
	Content  string // To read the file
}

func (f *File) Type() ObjectType { return FILE_OBJ }
func (f *File) Inspect() string  { return f.Filename }
func (f *File) Method(method string, args []Object) Object {
	switch method {
	case "soma":
		return f.read(args)
	}
	return nil
}

func (f *File) read(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	return &String{Value: f.Content}
}
