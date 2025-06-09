package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/lexer"
	"github.com/NuruProgramming/Nuru/module"
	"github.com/NuruProgramming/Nuru/object"
	"github.com/NuruProgramming/Nuru/parser"
)

var searchPaths []string

// Track modules that are currently being imported to detect circular imports
var importInProgress = make(map[string]bool)

func evalImport(node *ast.Import, env *object.Environment) object.Object {
	for k, v := range node.Identifiers {
		if mod, ok := module.Mapper[v.Value]; ok {
			env.Set(k, mod)
		} else {
			return evalImportFile(k, v, env)
		}
	}
	return NULL
}

func evalImportFile(name string, ident *ast.Identifier, env *object.Environment) object.Object {
	addSearchPath("")
	filename := findFile(name)
	if filename == "" {
		return newError("Moduli %s haipo", name)
	}

	// Check if this module is already being imported
	if importInProgress[filename] {
		fmt.Printf("Circular import detected for module: %s\n", name)
		// Return a placeholder or partially initialized module to break the cycle
		// For now, we'll just return NULL, but in a more sophisticated implementation,
		// you might want to return a partially initialized module
		return NULL
	}

	// Mark this module as being imported
	importInProgress[filename] = true

	var scope *object.Environment
	scope, err := evaluateFile(filename, env)

	// Clear the in-progress flag for this module
	delete(importInProgress, filename)

	if err != nil {
		return err
	}
	return importFile(name, ident, env, scope)
}

func addSearchPath(path string) {
	searchPaths = append(searchPaths, path)
}

func findFile(name string) string {
	basename := fmt.Sprintf("%s.nr", name)
	for _, path := range searchPaths {
		file := filepath.Join(path, basename)
		if fileExists(file) {
			return file
		}
	}
	return ""
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func evaluateFile(file string, env *object.Environment) (*object.Environment, object.Object) {
	source, err := os.ReadFile(file)
	if err != nil {
		return nil, &object.Error{Message: fmt.Sprintf("Tumeshindwa kufungua pakeji: %s", file)}
	}
	l := lexer.New(string(source))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return nil, &object.Error{Message: fmt.Sprintf("Pakeji %s ina makosa yafuatayo:\n%s", file, strings.Join(p.Errors(), "\n"))}
	}

	// Create a new environment for the imported module
	scope := object.NewEnvironment()

	// Add imported module name to scope for self-reference (if needed)
	// This allows the module to reference itself in its own definition
	moduleName := filepath.Base(file)
	if strings.HasSuffix(moduleName, ".nr") {
		moduleName = moduleName[:len(moduleName)-3]
	}

	result := Eval(program, scope)
	if isError(result) {
		return nil, result
	}
	return scope, nil
}

func importFile(name string, ident *ast.Identifier, env *object.Environment, scope *object.Environment) object.Object {
	value, ok := scope.Get(ident.Value)
	if !ok {
		return newError("%s sio pakeji", name)
	}

	// Handle potential circular references by using weak references for parent-child relationships
	// This is just a simple example and would need to be adapted to the specific needs of your system

	// For simplicity, let's just scan through all pairs in the dictionary and convert any value that
	// refers back to the parent environment to a weak reference
	if dict, ok := value.(*object.Dict); ok {
		for hashKey, pair := range dict.Pairs {
			// Check if the value is another dictionary or array that might contain circular references
			if _, isDict := pair.Value.(*object.Dict); isDict {
				// Replace with weak reference to break potential circular references
				dict.Pairs[hashKey] = object.DictPair{Key: pair.Key, Value: object.NewWeakReference(pair.Value)}
			}
		}
	}

	env.Set(name, value)
	return NULL
}
