package object

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// File represents a file in the Nuru language
type File struct {
	Filename    string      // Path to the file
	Content     string      // Content of the file (for read operations)
	FileDesc    *os.File    // Underlying file descriptor
	Mode        string      // Opening mode (r, r+, w, w+, a, a+)
	IsOpen      bool        // Whether the file is currently open
	Permissions os.FileMode // File permissions
}

func (f *File) Type() ObjectType { return FILE_OBJ }
func (f *File) Inspect() string  { return f.Filename }

// Method dispatches to the appropriate method based on the method name
func (f *File) Method(method string, args []Object) Object {
	switch method {
	case "soma":
		return f.read(args)
	case "andika":
		return f.write(args)
	case "ongeza":
		return f.append(args)
	case "funga":
		return f.close(args)
	case "tafuta":
		return f.seek(args)
	case "hali":
		return f.stat(args)
	case "isFungwa":
		return f.isClosed(args)
	default:
		return createError("Samahani, kiendesha hiki hakitumiki na faili: %s", method)
	}
}

// read reads the current content of the file
func (f *File) read(args []Object) Object {
	if len(args) != 0 {
		return createError("Samahani, soma inahitaji Hoja 0, wewe umeweka %d", len(args))
	}

	// If the file is open, read from its current position
	if f.IsOpen && f.FileDesc != nil {
		// Save current position to restore it after reading
		currentPos, err := f.FileDesc.Seek(0, io.SeekCurrent)
		if err != nil {
			return createError("Hitilafu wakati wa kusoma faili: %s", err.Error())
		}

		// Go to the beginning of the file
		_, err = f.FileDesc.Seek(0, io.SeekStart)
		if err != nil {
			return createError("Hitilafu wakati wa kusoma faili: %s", err.Error())
		}

		// Read the content
		content, err := ioutil.ReadAll(f.FileDesc)
		if err != nil {
			return createError("Hitilafu wakati wa kusoma faili: %s", err.Error())
		}

		// Restore original position
		_, err = f.FileDesc.Seek(currentPos, io.SeekStart)
		if err != nil {
			return createError("Hitilafu wakati wa kusoma faili: %s", err.Error())
		}

		return &String{Value: string(content)}
	}

	// If the file is not open, return the cached content
	return &String{Value: f.Content}
}

// write writes data to the file
func (f *File) write(args []Object) Object {
	if len(args) != 1 {
		return createError("Samahani, andika inahitaji Hoja 1, wewe umeweka %d", len(args))
	}

	var content string
	switch data := args[0].(type) {
	case *String:
		content = data.Value
	case *Integer:
		content = fmt.Sprintf("%d", data.Value)
	case *Float:
		content = fmt.Sprintf("%g", data.Value)
	case *Boolean:
		content = fmt.Sprintf("%t", data.Value)
	default:
		content = data.Inspect()
	}

	// If the file is open, write to the file descriptor
	if f.IsOpen && f.FileDesc != nil {
		// Check if the file is writable
		if f.Mode == "r" {
			return createError("Hitilafu: Faili imefunguliwa kwa kusoma tu")
		}

		// Clear the file first if mode is "w" or "w+"
		if f.Mode == "w" || f.Mode == "w+" {
			err := f.FileDesc.Truncate(0)
			if err != nil {
				return createError("Hitilafu wakati wa kuandika kwenye faili: %s", err.Error())
			}
			_, err = f.FileDesc.Seek(0, io.SeekStart)
			if err != nil {
				return createError("Hitilafu wakati wa kuandika kwenye faili: %s", err.Error())
			}
		}

		_, err := f.FileDesc.WriteString(content)
		if err != nil {
			return createError("Hitilafu wakati wa kuandika kwenye faili: %s", err.Error())
		}

		// Update the cached content
		f.Content = content

		return &Boolean{Value: true}
	}

	// If the file is not open, write directly to the file
	err := os.WriteFile(f.Filename, []byte(content), 0644)
	if err != nil {
		return createError("Hitilafu katika kuandika faili: %s", err.Error())
	}

	// Update the cached content
	f.Content = content

	return &Boolean{Value: true}
}

// append adds data to the end of the file
func (f *File) append(args []Object) Object {
	if len(args) != 1 {
		return createError("Samahani, ongeza inahitaji Hoja 1, wewe umeweka %d", len(args))
	}

	var content string
	switch data := args[0].(type) {
	case *String:
		content = data.Value
	case *Integer:
		content = fmt.Sprintf("%d", data.Value)
	case *Float:
		content = fmt.Sprintf("%g", data.Value)
	case *Boolean:
		content = fmt.Sprintf("%t", data.Value)
	default:
		content = data.Inspect()
	}

	// If the file is open, append to the file descriptor
	if f.IsOpen && f.FileDesc != nil {
		// Check if the file is writable
		if f.Mode == "r" {
			return createError("Hitilafu: Faili imefunguliwa kwa kusoma tu")
		}

		// Go to the end of the file
		_, err := f.FileDesc.Seek(0, io.SeekEnd)
		if err != nil {
			return createError("Hitilafu wakati wa kuongeza kwenye faili: %s", err.Error())
		}

		_, err = f.FileDesc.WriteString(content)
		if err != nil {
			return createError("Hitilafu wakati wa kuongeza kwenye faili: %s", err.Error())
		}

		// Update the cached content
		f.Content += content

		return &Boolean{Value: true}
	}

	// If the file is not open, append directly to the file
	file, err := os.OpenFile(f.Filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return createError("Hitilafu katika kufungua faili: %s", err.Error())
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return createError("Hitilafu katika kuongeza kwa faili: %s", err.Error())
	}

	// Update the cached content
	f.Content += content

	return &Boolean{Value: true}
}

// close closes the file
func (f *File) close(args []Object) Object {
	if len(args) != 0 {
		return createError("Samahani, funga inahitaji Hoja 0, wewe umeweka %d", len(args))
	}

	if f.IsOpen && f.FileDesc != nil {
		err := f.FileDesc.Close()
		if err != nil {
			return createError("Hitilafu wakati wa kufunga faili: %s", err.Error())
		}
		f.IsOpen = false
		f.FileDesc = nil
		return &Boolean{Value: true}
	}

	return &Boolean{Value: false}
}

// seek changes the file position
func (f *File) seek(args []Object) Object {
	if len(args) != 2 {
		return createError("Samahani, tafuta inahitaji Hoja 2, wewe umeweka %d", len(args))
	}

	if !f.IsOpen || f.FileDesc == nil {
		return createError("Hitilafu: Faili haijafunguliwa")
	}

	offsetObj, ok := args[0].(*Integer)
	if !ok {
		return createError("Samahani, hoja ya kwanza lazima iwe namba")
	}

	whenceObj, ok := args[1].(*Integer)
	if !ok {
		return createError("Samahani, hoja ya pili lazima iwe namba")
	}

	whence := int(whenceObj.Value)
	if whence < 0 || whence > 2 {
		return createError("Samahani, hoja ya pili lazima iwe 0 (mwanzo), 1 (sasa), au 2 (mwisho)")
	}

	newPos, err := f.FileDesc.Seek(offsetObj.Value, whence)
	if err != nil {
		return createError("Hitilafu wakati wa kubadilisha nafasi ya faili: %s", err.Error())
	}

	return &Integer{Value: newPos}
}

// stat gets file information
func (f *File) stat(args []Object) Object {
	if len(args) != 0 {
		return createError("Samahani, hali inahitaji Hoja 0, wewe umeweka %d", len(args))
	}

	info, err := os.Stat(f.Filename)
	if err != nil {
		return createError("Hitilafu katika kupata taarifa za faili: %s", err.Error())
	}

	// Create a dictionary with file information
	fileInfo := &Dict{Pairs: make(map[HashKey]DictPair)}
	fileInfo.Set(&String{Value: "jina"}, &String{Value: info.Name()})
	fileInfo.Set(&String{Value: "ni_sarafa"}, &Boolean{Value: info.IsDir()})
	fileInfo.Set(&String{Value: "ukubwa"}, &Integer{Value: info.Size()})
	fileInfo.Set(&String{Value: "ruhusu"}, &Integer{Value: int64(info.Mode())})
	fileInfo.Set(&String{Value: "muda"}, &String{Value: info.ModTime().String()})

	return fileInfo
}

// isClosed checks if the file is closed
func (f *File) isClosed(args []Object) Object {
	if len(args) != 0 {
		return createError("Samahani, isFungwa inahitaji Hoja 0, wewe umeweka %d", len(args))
	}

	return &Boolean{Value: !f.IsOpen}
}

// OpenFile opens a file with the specified mode and returns a File object
func OpenFile(filename string, mode string) (*File, error) {
	var flag int
	var perm os.FileMode = 0644

	// Parse mode string to flags
	switch mode {
	case "r":
		flag = os.O_RDONLY
	case "r+":
		flag = os.O_RDWR
	case "w":
		flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	case "w+":
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	case "a":
		flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	case "a+":
		flag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	default:
		return nil, fmt.Errorf("mode isiyotambuliwa: %s", mode)
	}

	file, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}

	// Read file content if opening for reading
	var content string
	if flag&os.O_RDONLY != 0 || flag&os.O_RDWR != 0 {
		data, err := ioutil.ReadAll(file)
		if err != nil {
			file.Close()
			return nil, err
		}
		content = string(data)

		// Reset file pointer to beginning for future reads
		file.Seek(0, 0)
	}

	// Create a Nuru File object
	fileObj := &File{
		Filename:    filename,
		Content:     content,
		FileDesc:    file,
		Mode:        mode,
		IsOpen:      true,
		Permissions: perm,
	}

	return fileObj, nil
}

// IsDirectory checks if a path is a directory
func IsDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return info.IsDir(), nil
}

// IsFile checks if a path is a regular file
func IsFile(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return !info.IsDir(), nil
}

// ReadDir reads a directory and returns a slice of file info objects
func ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

// createError is a helper function to create error objects
func createError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
