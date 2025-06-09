package module

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/NuruProgramming/Nuru/object"
)

// FSFunctions contains all the file system related functions
var FSFunctions = map[string]object.ModuleFunction{}

func init() {
	// File operations
	FSFunctions["soma"] = readFile     // readFile
	FSFunctions["andika"] = writeFile  // writeFile
	FSFunctions["ongeza"] = appendFile // appendFile
	FSFunctions["fungua"] = openFile   // open
	FSFunctions["futa"] = removeFile   // unlink/remove
	FSFunctions["fanya"] = createFile  // create
	FSFunctions["ipo"] = exists        // exists

	// Directory operations
	FSFunctions["orodha"] = readdir        // readdir
	FSFunctions["tengenezaSarafa"] = mkdir // mkdir
	FSFunctions["futaSarafa"] = rmdir      // rmdir

	// Path operations
	FSFunctions["hali"] = stat            // stat
	FSFunctions["niSarafa"] = isDirectory // isDirectory
	FSFunctions["niFaili"] = isFile       // isFile

	// Permissions and attributes
	FSFunctions["ruhusu"] = chmod     // chmod
	FSFunctions["mmiliki"] = chown    // chown
	FSFunctions["badilisha"] = rename // rename

	// Symbolic links
	FSFunctions["kiungo"] = symlink      // symlink
	FSFunctions["somaKiungo"] = readlink // readlink

	// File descriptors
	FSFunctions["funga"] = closeFile // close
}

// Helper function to convert Nuru objects to Go types
func convertToGoType(obj object.Object) interface{} {
	switch obj := obj.(type) {
	case *object.String:
		return obj.Value
	case *object.Integer:
		return int(obj.Value)
	case *object.Boolean:
		return obj.Value
	case *object.Array:
		goSlice := make([]interface{}, len(obj.Elements))
		for i, elem := range obj.Elements {
			goSlice[i] = convertToGoType(elem)
		}
		return goSlice
	case *object.Dict:
		goMap := make(map[string]interface{})
		for _, pair := range obj.Pairs {
			hashKey := pair.Key.Inspect()
			goMap[hashKey] = convertToGoType(pair.Value)
		}
		return goMap
	default:
		return nil
	}
}

// Helper function to convert Go types to Nuru objects
func convertToNuruObject(val interface{}) object.Object {
	if val == nil {
		return &object.Null{}
	}

	switch v := val.(type) {
	case string:
		return &object.String{Value: v}
	case int:
		return &object.Integer{Value: int64(v)}
	case int64:
		return &object.Integer{Value: v}
	case bool:
		return &object.Boolean{Value: v}
	case []interface{}:
		elements := make([]object.Object, len(v))
		for i, elem := range v {
			elements[i] = convertToNuruObject(elem)
		}
		return &object.Array{Elements: elements}
	case map[string]interface{}:
		dict := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
		for k, v := range v {
			key := &object.String{Value: k}
			value := convertToNuruObject(v)
			dict.Set(key, value)
		}
		return dict
	case os.FileInfo:
		dict := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
		dict.Set(&object.String{Value: "jina"}, &object.String{Value: v.Name()})
		dict.Set(&object.String{Value: "ukubwa"}, &object.Integer{Value: v.Size()})
		dict.Set(&object.String{Value: "ruhusu"}, &object.Integer{Value: int64(v.Mode())})
		dict.Set(&object.String{Value: "muda"}, &object.String{Value: v.ModTime().String()})
		dict.Set(&object.String{Value: "ni_sarafa"}, &object.Boolean{Value: v.IsDir()})
		return dict
	default:
		return &object.Null{}
	}
}

// newError returns a new error object with formatted message
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// readFile reads the entire contents of a file
func readFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("soma: inahitaji hoja 1 (jina la faili), ulitoa %d", len(args))
	}

	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("soma: hoja inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	// Create a temporary File object and use its read method
	file := &object.File{
		Filename: filename.Value,
	}

	data, err := ioutil.ReadFile(filename.Value)
	if err != nil {
		return newError("soma: %s", err.Error())
	}

	file.Content = string(data)
	return file.Method("soma", []object.Object{})
}

// writeFile writes data to a file, replacing the file if it already exists
func writeFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return newError("andika: inahitaji hoja 2 (jina la faili, data), ulitoa %d", len(args))
	}

	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("andika: hoja ya kwanza inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	// Create a file object and use its write method
	file := &object.File{
		Filename: filename.Value,
	}

	return file.Method("andika", []object.Object{args[1]})
}

// appendFile appends data to a file, creating the file if it does not exist
func appendFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return newError("ongeza: inahitaji hoja 2 (jina la faili, data), ulitoa %d", len(args))
	}

	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("ongeza: hoja ya kwanza inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	// Create a file object and use its append method
	file := &object.File{
		Filename: filename.Value,
	}

	return file.Method("ongeza", []object.Object{args[1]})
}

// openFile opens a file for reading or writing
func openFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("fungua: inahitaji hoja 1-2 (jina la faili, [mode]), ulitoa %d", len(args))
	}

	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("fungua: hoja ya kwanza inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	// Default mode is read-only
	mode := "r"
	if len(args) == 2 {
		modeArg, ok := args[1].(*object.String)
		if !ok {
			return newError("fungua: hoja ya pili inapaswa kuwa neno (string), ulitoa %s", args[1].Type())
		}
		mode = modeArg.Value
	}

	// Use the enhanced OpenFile function from object package
	fileObj, err := object.OpenFile(filename.Value, mode)
	if err != nil {
		return newError("fungua: %s", err.Error())
	}

	return fileObj
}

// closeFile closes a file descriptor
func closeFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("funga: inahitaji hoja 1 (file descriptor), ulitoa %d", len(args))
	}

	fileObj, ok := args[0].(*object.File)
	if !ok {
		return newError("funga: hoja inapaswa kuwa file object, ulitoa %s", args[0].Type())
	}

	// Use the file's close method
	return fileObj.Method("funga", []object.Object{})
}

// removeFile removes a file or directory
func removeFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("futa: inahitaji hoja 1 (jina la faili), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("futa: hoja inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	err := os.Remove(path.Value)
	if err != nil {
		return newError("futa: %s", err.Error())
	}

	return &object.Boolean{Value: true}
}

// createFile creates a file for writing
func createFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("fanya: inahitaji hoja 1 (jina la faili), ulitoa %d", len(args))
	}

	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("fanya: hoja inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	// Use the enhanced OpenFile function with "w" mode to create a new file
	fileObj, err := object.OpenFile(filename.Value, "w")
	if err != nil {
		return newError("fanya: %s", err.Error())
	}

	return fileObj
}

// exists checks if a file or directory exists
func exists(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("ipo: inahitaji hoja 1 (jina la faili), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("ipo: hoja inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	_, err := os.Stat(path.Value)
	exists := !os.IsNotExist(err)

	return &object.Boolean{Value: exists}
}

// readdir reads the contents of a directory
func readdir(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("orodha: inahitaji hoja 1 (jina la sarafa), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("orodha: hoja inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	files, err := ioutil.ReadDir(path.Value)
	if err != nil {
		return newError("orodha: %s", err.Error())
	}

	elements := make([]object.Object, len(files))
	for i, file := range files {
		elements[i] = convertToNuruObject(file)
	}

	return &object.Array{Elements: elements}
}

// mkdir creates a directory
func mkdir(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("tengenezaSarafa: inahitaji hoja 1-2 (jina la sarafa, [recursive]), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("tengenezaSarafa: hoja ya kwanza inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	// Default is non-recursive
	recursive := false
	if len(args) == 2 {
		recursiveArg, ok := args[1].(*object.Boolean)
		if !ok {
			return newError("tengenezaSarafa: hoja ya pili inapaswa kuwa boolean, ulitoa %s", args[1].Type())
		}
		recursive = recursiveArg.Value
	}

	var err error
	if recursive {
		err = os.MkdirAll(path.Value, 0755)
	} else {
		err = os.Mkdir(path.Value, 0755)
	}

	if err != nil {
		return newError("tengenezaSarafa: %s", err.Error())
	}

	return &object.Boolean{Value: true}
}

// rmdir removes a directory
func rmdir(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("futaSarafa: inahitaji hoja 1-2 (jina la sarafa, [recursive]), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("futaSarafa: hoja ya kwanza inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	// Default is non-recursive
	recursive := false
	if len(args) == 2 {
		recursiveArg, ok := args[1].(*object.Boolean)
		if !ok {
			return newError("futaSarafa: hoja ya pili inapaswa kuwa boolean, ulitoa %s", args[1].Type())
		}
		recursive = recursiveArg.Value
	}

	var err error
	if recursive {
		err = os.RemoveAll(path.Value)
	} else {
		err = os.Remove(path.Value)
	}

	if err != nil {
		return newError("futaSarafa: %s", err.Error())
	}

	return &object.Boolean{Value: true}
}

// stat returns information about a file
func stat(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("hali: inahitaji hoja 1 (jina la faili), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("hali: hoja inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	info, err := os.Stat(path.Value)
	if err != nil {
		return newError("hali: %s", err.Error())
	}

	return convertToNuruObject(info)
}

// isDirectory checks if a path is a directory
func isDirectory(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("niSarafa: inahitaji hoja 1 (jina la faili), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("niSarafa: hoja inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	info, err := os.Stat(path.Value)
	if err != nil {
		// If path doesn't exist, it's not a directory
		return &object.Boolean{Value: false}
	}

	return &object.Boolean{Value: info.IsDir()}
}

// isFile checks if a path is a file
func isFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("niFaili: inahitaji hoja 1 (jina la faili), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("niFaili: hoja inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	info, err := os.Stat(path.Value)
	if err != nil {
		// If path doesn't exist, it's not a file
		return &object.Boolean{Value: false}
	}

	return &object.Boolean{Value: !info.IsDir()}
}

// chmod changes file permissions
func chmod(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return newError("ruhusu: inahitaji hoja 2 (jina la faili, mode), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("ruhusu: hoja ya kwanza inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	mode, ok := args[1].(*object.Integer)
	if !ok {
		return newError("ruhusu: hoja ya pili inapaswa kuwa namba (integer), ulitoa %s", args[1].Type())
	}

	err := os.Chmod(path.Value, os.FileMode(mode.Value))
	if err != nil {
		return newError("ruhusu: %s", err.Error())
	}

	return &object.Boolean{Value: true}
}

// chown changes file owner
func chown(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return newError("mmiliki: inahitaji hoja 3 (jina la faili, uid, gid), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("mmiliki: hoja ya kwanza inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	uid, ok := args[1].(*object.Integer)
	if !ok {
		return newError("mmiliki: hoja ya pili inapaswa kuwa namba (integer), ulitoa %s", args[1].Type())
	}

	gid, ok := args[2].(*object.Integer)
	if !ok {
		return newError("mmiliki: hoja ya tatu inapaswa kuwa namba (integer), ulitoa %s", args[2].Type())
	}

	err := os.Chown(path.Value, int(uid.Value), int(gid.Value))
	if err != nil {
		return newError("mmiliki: %s", err.Error())
	}

	return &object.Boolean{Value: true}
}

// rename renames a file or directory
func rename(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return newError("badilisha: inahitaji hoja 2 (jina la zamani, jina jipya), ulitoa %d", len(args))
	}

	oldPath, ok := args[0].(*object.String)
	if !ok {
		return newError("badilisha: hoja ya kwanza inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	newPath, ok := args[1].(*object.String)
	if !ok {
		return newError("badilisha: hoja ya pili inapaswa kuwa neno (string), ulitoa %s", args[1].Type())
	}

	err := os.Rename(oldPath.Value, newPath.Value)
	if err != nil {
		return newError("badilisha: %s", err.Error())
	}

	return &object.Boolean{Value: true}
}

// symlink creates a symbolic link
func symlink(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return newError("kiungo: inahitaji hoja 2 (lengo, jina la kiungo), ulitoa %d", len(args))
	}

	target, ok := args[0].(*object.String)
	if !ok {
		return newError("kiungo: hoja ya kwanza inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	linkname, ok := args[1].(*object.String)
	if !ok {
		return newError("kiungo: hoja ya pili inapaswa kuwa neno (string), ulitoa %s", args[1].Type())
	}

	err := os.Symlink(target.Value, linkname.Value)
	if err != nil {
		return newError("kiungo: %s", err.Error())
	}

	return &object.Boolean{Value: true}
}

// readlink reads the destination of a symbolic link
func readlink(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("somaKiungo: inahitaji hoja 1 (jina la kiungo), ulitoa %d", len(args))
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return newError("somaKiungo: hoja inapaswa kuwa neno (string), ulitoa %s", args[0].Type())
	}

	linkDest, err := os.Readlink(path.Value)
	if err != nil {
		return newError("somaKiungo: %s", err.Error())
	}

	return &object.String{Value: linkDest}
}
