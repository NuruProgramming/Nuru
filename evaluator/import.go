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
	var scope *object.Environment
	scope, err := evaluateFile(filename, env)
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

	scope := object.NewEnvironment()
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
	env.Set(name, value)
	return NULL
}
