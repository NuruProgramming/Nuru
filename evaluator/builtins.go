package evaluator

import "github.com/AvicennaJr/Nuru/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Hoja hazilingani, tunahitaji=1, tumepewa=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("Samahani, hii function haitumiki na %s", args[0].Type())
			}
		},
	},
}
