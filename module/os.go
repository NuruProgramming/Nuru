package module

import (
	"os"
	"os/exec"
	"strings"

	"github.com/NuruProgramming/Nuru/object"
)

var OsFunctions = map[string]object.ModuleFunction{}

func init() {
	OsFunctions["toka"] = exit
	OsFunctions["kimbiza"] = run
}

func exit(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 1 {
		return &object.Error{Message: "Hoja sii sahihi"}
	}

	if len(args) == 1 {
		status, ok := args[0].(*object.Integer)
		if !ok {
			return &object.Error{Message: "Hoja sii namba"}
		}
		os.Exit(int(status.Value))
		return nil
	}

	os.Exit(0)

	return nil
}

func run(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Idadi ya hoja sii sahihi"}
	}

	cmd, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja lazima iwe neno"}
	}
	cmdMain := cmd.Value
	cmdArgs := strings.Split(cmdMain, " ")
	cmdArgs = cmdArgs[1:]

	out, err := exec.Command(cmdMain, cmdArgs...).Output()
	if err != nil {
		return &object.Error{Message: "Tumeshindwa kukimbiza komandi"}
	}

	return &object.String{Value: string(out)}
}
