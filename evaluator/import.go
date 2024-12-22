package evaluator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/lexer"

	"github.com/NuruProgramming/Nuru/object"
	"github.com/NuruProgramming/Nuru/parser"
)

var searchPaths []string

type maktaba int

const (
	MAKTABA_YA_NURU maktaba = iota
	MAKTABA_YA_HAPA
	MAKTABA_YA_NJE
)

func evalImport(node *ast.Import, env *object.Environment) object.Object {
	/*
		Maktaba yanafuata mkondo wa:
			mkt => ni maktaba ya kawaida (stdlib)
			hapa => ni maktaba haya, tunaanza na faili "nuru.toml" (ama faili ya kisawa)
			_ => majina mengine ni ya maktaba ya nje
	*/

	// Go through the Identifiers
	for k, v := range node.Identifiers {
		vs := strings.Split(v.Value, "/")
		if len(vs) <= 0 {
			break
		}

		switch vs[0] {
		case "mkt":
			loc := mahali_pa_maktaba(MAKTABA_YA_NURU)
			ss := strings.Split(v.Value, "/")[1:]

			fi__ := filepath.Join(loc)
			g__ := filepath.Join(ss...)
			fi := filepath.Join(fi__, g__)

			return evalImportFile(k, &ast.Identifier{Value: fi}, env)
		default:
			log.Printf("Maktaba ya nje hazijazaidiwa '%s'\n", k)
			return NULL
		}
	}

	return NULL
}

func mahali_pa_maktaba(mkt maktaba) string {
	switch mkt {
	case MAKTABA_YA_NURU:
		loc := os.Getenv("MAKTABA_YA_NURU")
		if len(loc) > 0 {
			return loc
		}

		var usr_lib string
		var lib_ string

		if runtime.GOOS == "windows" {
			uhd, _ := os.UserHomeDir()
			usr_lib = filepath.Join(filepath.Base(uhd), "nuru", "maktaba")
		} else {
			usr_lib = "/usr/lib/nuru/maktaba"
			lib_ = "/lib/nuru/maktaba"
		}
		if fileExists(usr_lib) {
			return usr_lib
		}
		if fileExists(lib_) {
			return lib_
		}

		return usr_lib
	default:
		return ""
	}
}

func evalImportFile(name string, ident *ast.Identifier, env *object.Environment) object.Object {
	addSearchPath("")
	filename := findFile(ident.Value)
	if filename == "" {
		return newError("Moduli %s haipo", name)
	}
	var scope *object.Environment
	scope, err := evaluateFile(name, filename, env)
	if err != nil {
		return err
	}
	return importFile(name, env, scope)
}

func addSearchPath(path string) {
	searchPaths = append(searchPaths, path)
}

func findFile(name string) string {
	basename := fmt.Sprintf("%s.nuru", name)
	fmt.Println(basename)
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

func evaluateFile(name, file string, env *object.Environment) (*object.Environment, object.Object) {
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

	pkg := &object.Package{
		Name:  &ast.Identifier{Value: name},
		Env:   env,
		Scope: object.NewEnclosedEnvironment(env),
	}

	result := Eval(program, pkg.Scope)

	env.Set(name, pkg)

	if isError(result) {
		return nil, result
	}
	return pkg.Env, nil
}

func importFile(name string, env *object.Environment, scope *object.Environment) object.Object {
	value, ok := scope.Get(name)
	if !ok {
		return newError("%s sio pakeji", name)
	}
	env.Set(name, value)
	return NULL
}
