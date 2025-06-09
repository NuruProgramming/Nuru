package module

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NuruProgramming/Nuru/object"
)

var PathFunctions = map[string]object.ModuleFunction{}

func init() {
	PathFunctions["jina"] = basename       // path.basename
	PathFunctions["kigawaji"] = delimiter  // path.delimiter
	PathFunctions["sarafa"] = dirname      // path.dirname
	PathFunctions["ext"] = extname         // path.extname
	PathFunctions["umbiza"] = format       // path.format
	PathFunctions["niKamili"] = isAbsolute // path.isAbsolute
	PathFunctions["unganisha"] = join      // path.join
	PathFunctions["sawazisha"] = normalize // path.normalize
	PathFunctions["changanua"] = parse     // path.parse
	PathFunctions["husika"] = relative     // path.relative
	PathFunctions["tatua"] = resolve       // path.resolve
	PathFunctions["kitenga"] = separator   // path.sep
	PathFunctions["posix"] = posixObject   // path.posix - returns a path object using POSIX rules
	PathFunctions["win32"] = win32Object   // path.win32 - returns a path object using Windows rules
}

// basename returns the last part of a path, similar to the Unix basename command
func basename(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Jina inahitaji hoja ya kwanza (njia)"}
	}

	pathStr, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la njia"}
	}

	// Optional extension to remove
	ext := ""
	if len(args) > 1 {
		if extObj, ok := args[1].(*object.String); ok {
			ext = extObj.Value
		}
	}

	// Use the filepath package to get the base name
	base := filepath.Base(pathStr.Value)

	// If extension is provided and matches the end of the base, remove it
	if ext != "" && strings.HasSuffix(base, ext) {
		base = base[:len(base)-len(ext)]
	}

	return &object.String{Value: base}
}

// delimiter returns the platform-specific path delimiter
func delimiter(args []object.Object, defs map[string]object.Object) object.Object {
	// In Unix/Linux/MacOS, the delimiter is ":", in Windows it's ";"
	delimiter := string(os.PathListSeparator)
	return &object.String{Value: delimiter}
}

// dirname returns the directory part of a path
func dirname(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Sarafa inahitaji hoja ya kwanza (njia)"}
	}

	pathStr, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la njia"}
	}

	// Use the filepath package to get the directory
	dir := filepath.Dir(pathStr.Value)
	return &object.String{Value: dir}
}

// extname returns the extension of a path
func extname(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Ext inahitaji hoja ya kwanza (njia)"}
	}

	pathStr, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la njia"}
	}

	// Use the filepath package to get the extension
	ext := filepath.Ext(pathStr.Value)
	return &object.String{Value: ext}
}

// format returns a path string from an object
func format(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Umbiza inahitaji hoja ya kwanza (object ya njia)"}
	}

	pathObj, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe kamusi ya njia"}
	}

	// Extract components
	dir := ""
	root := ""
	base := ""
	name := ""
	ext := ""

	// Get dir
	if dirObj, ok := pathObj.Get(&object.String{Value: "dir"}); ok {
		if dirStr, ok := dirObj.(*object.String); ok {
			dir = dirStr.Value
		}
	}

	// Get root
	if rootObj, ok := pathObj.Get(&object.String{Value: "root"}); ok {
		if rootStr, ok := rootObj.(*object.String); ok {
			root = rootStr.Value
		}
	}

	// Get base
	if baseObj, ok := pathObj.Get(&object.String{Value: "base"}); ok {
		if baseStr, ok := baseObj.(*object.String); ok {
			base = baseStr.Value
		}
	}

	// Get name
	if nameObj, ok := pathObj.Get(&object.String{Value: "name"}); ok {
		if nameStr, ok := nameObj.(*object.String); ok {
			name = nameStr.Value
		}
	}

	// Get ext
	if extObj, ok := pathObj.Get(&object.String{Value: "ext"}); ok {
		if extStr, ok := extObj.(*object.String); ok {
			ext = extStr.Value
		}
	}

	// Format the path according to Node.js rules
	result := ""

	// If dir is provided, use it + separator + base
	if dir != "" {
		if base != "" {
			result = dir + string(filepath.Separator) + base
		} else {
			// If base is not provided, use name + ext
			result = dir + string(filepath.Separator) + name + ext
		}
	} else if root != "" {
		// If root is provided but not dir
		if base != "" {
			result = root + base
		} else {
			result = root + name + ext
		}
	} else {
		// Neither dir nor root is provided
		if base != "" {
			result = base
		} else {
			result = name + ext
		}
	}

	return &object.String{Value: result}
}

// isAbsolute determines if the path is an absolute path
func isAbsolute(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "NiKamili inahitaji hoja ya kwanza (njia)"}
	}

	pathStr, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la njia"}
	}

	// Use the filepath package to check if the path is absolute
	isAbs := filepath.IsAbs(pathStr.Value)
	return &object.Boolean{Value: isAbs}
}

// join joins all path segments together
func join(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Unganisha inahitaji hoja angalau moja"}
	}

	// Convert all arguments to strings
	pathParts := []string{}
	for _, arg := range args {
		if pathStr, ok := arg.(*object.String); ok {
			pathParts = append(pathParts, pathStr.Value)
		} else {
			return &object.Error{Message: "Hoja zote lazima ziwe maneno ya njia"}
		}
	}

	// Use the filepath package to join the paths
	joined := filepath.Join(pathParts...)
	return &object.String{Value: joined}
}

// normalize normalizes the path
func normalize(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Sawazisha inahitaji hoja ya kwanza (njia)"}
	}

	pathStr, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la njia"}
	}

	// Use the filepath package to clean the path
	normalized := filepath.Clean(pathStr.Value)
	return &object.String{Value: normalized}
}

// parse returns an object with path components
func parse(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Changanua inahitaji hoja ya kwanza (njia)"}
	}

	pathStr, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la njia"}
	}

	// Use the filepath package to get the components
	dir := filepath.Dir(pathStr.Value)
	base := filepath.Base(pathStr.Value)
	ext := filepath.Ext(pathStr.Value)
	name := strings.TrimSuffix(base, ext)

	// Determine the root
	root := ""
	if filepath.IsAbs(pathStr.Value) {
		// For Windows, the root is the drive letter with separator (e.g., "C:\")
		// For POSIX, the root is "/"
		if len(pathStr.Value) >= 2 && pathStr.Value[1] == ':' {
			// Windows drive letter
			root = pathStr.Value[:2] + string(filepath.Separator)
		} else if strings.HasPrefix(pathStr.Value, string(filepath.Separator)) {
			// POSIX root
			root = string(filepath.Separator)
		}
	}

	// Create the result object
	result := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	result.Set(&object.String{Value: "root"}, &object.String{Value: root})
	result.Set(&object.String{Value: "dir"}, &object.String{Value: dir})
	result.Set(&object.String{Value: "base"}, &object.String{Value: base})
	result.Set(&object.String{Value: "name"}, &object.String{Value: name})
	result.Set(&object.String{Value: "ext"}, &object.String{Value: ext})

	return result
}

// relative returns the relative path from 'from' to 'to'
func relative(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "Husika inahitaji hoja mbili: kutoka na mpaka"}
	}

	fromStr, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la njia"}
	}

	toStr, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya pili lazima iwe neno la njia"}
	}

	// Use the filepath package to get the relative path
	rel, err := filepath.Rel(fromStr.Value, toStr.Value)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Hitilafu kupata njia ya husika: %s", err)}
	}

	return &object.String{Value: rel}
}

// resolve resolves all paths into an absolute path
func resolve(args []object.Object, defs map[string]object.Object) object.Object {
	// Prepare path segments
	pathSegments := []string{}

	// If no args, return current directory
	if len(args) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Hitilafu kupata sarafa ya sasa: %s", err)}
		}
		return &object.String{Value: cwd}
	}

	// Process path segments from right to left
	for i := len(args) - 1; i >= 0; i-- {
		if pathStr, ok := args[i].(*object.String); ok {
			// If this path is absolute, it becomes our base
			if filepath.IsAbs(pathStr.Value) {
				// Clean path and return it joined with any segments we've collected
				result := filepath.Join(pathStr.Value, filepath.Join(pathSegments...))
				return &object.String{Value: filepath.Clean(result)}
			}

			// Otherwise add it to our segments
			pathSegments = append([]string{pathStr.Value}, pathSegments...)
		} else {
			return &object.Error{Message: "Hoja zote lazima ziwe maneno ya njia"}
		}
	}

	// If we get here, we need to add the current directory
	cwd, err := os.Getwd()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Hitilafu kupata sarafa ya sasa: %s", err)}
	}

	result := filepath.Join(cwd, filepath.Join(pathSegments...))
	return &object.String{Value: filepath.Clean(result)}
}

// separator returns the platform-specific path separator
func separator(args []object.Object, defs map[string]object.Object) object.Object {
	// Just return the platform-specific path separator
	return &object.String{Value: string(filepath.Separator)}
}

// createPathModuleObject creates a path module object with POSIX or Windows specific behavior
func createPathModuleObject(usePosix bool) object.Object {
	pathModule := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Add all path functions with appropriate behavior
	pathModule.Set(&object.String{Value: "jina"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Implementation with posix or windows specific behavior
			return basename(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "kigawaji"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if usePosix {
				return &object.String{Value: ":"}
			}
			return &object.String{Value: ";"}
		},
	})

	pathModule.Set(&object.String{Value: "sarafa"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return dirname(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "ext"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return extname(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "umbiza"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return format(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "niKamili"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return isAbsolute(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "unganisha"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return join(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "sawazisha"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return normalize(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "changanua"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return parse(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "husika"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return relative(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "tatua"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return resolve(args, nil)
		},
	})

	pathModule.Set(&object.String{Value: "kitenga"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if usePosix {
				return &object.String{Value: "/"}
			}
			return &object.String{Value: "\\"}
		},
	})

	return pathModule
}

// posixObject returns a path module object with POSIX behavior
func posixObject(args []object.Object, defs map[string]object.Object) object.Object {
	return createPathModuleObject(true)
}

// win32Object returns a path module object with Windows behavior
func win32Object(args []object.Object, defs map[string]object.Object) object.Object {
	return createPathModuleObject(false)
}
