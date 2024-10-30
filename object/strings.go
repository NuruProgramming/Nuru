package object

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

type String struct {
	Value  string
	offset int
}

func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
func (s *String) Next() (Object, Object) {
	offset := s.offset
	if len(s.Value) > offset {
		s.offset = offset + 1
		return &Integer{Value: int64(offset)}, &String{Value: string(s.Value[offset])}
	}
	return nil, nil
}
func (s *String) Reset() {
	s.offset = 0
}
func (s *String) Method(method string, args []Object) Object {
	switch method {
	case "idadi":
		return s.len(args)
	case "herufikubwa":
		return s.upper(args)
	case "herufindogo":
		return s.lower(args)
	case "gawa":
		return s.split(args)
	case "badilisha":
		return s.format(args)
	default:
		return newError("Samahani, kiendesha hiki hakitumiki na tungo (Neno)")
	}
}

func (s *String) len(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	return &Integer{Value: int64(len(s.Value))}
}

func (s *String) upper(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	return &String{Value: strings.ToUpper(s.Value)}
}

func (s *String) lower(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	return &String{Value: strings.ToLower(s.Value)}
}

func (s *String) split(args []Object) Object {
	if len(args) > 1 {
		return newError("Samahani, tunahitaji Hoja 1 au 0, wewe umeweka %d", len(args))
	}
	sep := " "
	if len(args) == 1 {
		sep = args[0].(*String).Value
	}
	parts := strings.Split(s.Value, sep)
	length := len(parts)
	elements := make([]Object, length)
	for k, v := range parts {
		elements[k] = &String{Value: v}
	}
	return &Array{Elements: elements}
}

func (s *String) format(args []Object) Object {
	value, err := formatStr(s.Value, args)

	if err != nil {
		return newError(err.Error())
	}

	return &String{Value: value}
}

// Format the string like "Hello {0}", jina ...
func formatStr(format string, options []Object) (string, error) {
	var str strings.Builder
	var val strings.Builder
	var check_val bool
	var opts_len int = len(options)

	// This is to enable escaping the {} braces if you want them included
	var escapeChar bool

	type optM struct {
		val bool
		obj Object
	}

	var optionsMap = make(map[int]optM, opts_len)

	// convert the options into a map (may not be the most efficient but we are not going fast)
	for i, optm := range options {
		optionsMap[i] = optM{val: false, obj: optm}
	}

	// Now go through the format string and do the replacement(s)
	// this has approx time complexity of O(n) (bestest case)
	for _, opt := range format {

		if !escapeChar && opt == '\\' {
			escapeChar = true
			continue
		}

		if opt == '{' && !escapeChar {
			check_val = true
			continue
		}

		if escapeChar {
			if opt != '{' && opt != '}' {
				str.WriteRune('\\')
			}
			escapeChar = false
		}

		if check_val && opt == '}' {
			vstr := strings.TrimSpace(val.String()) // remove accidental spaces

			// check the value and try to convert to int
			arrv, err := strconv.Atoi(vstr)
			if err != nil {
				// return non-cryptic go errors
				return "", fmt.Errorf(fmt.Sprintf("Ulichopeana si NAMBA, jaribu tena: `%s'", vstr))
			}

			oVal, exists := optionsMap[arrv]

			if !exists {
				return "", fmt.Errorf(fmt.Sprintf("Nambari ya chaguo unalolitaka %d ni kubwa kuliko ulizopeana (%d)", arrv, opts_len))
			}

			str.WriteString(oVal.obj.Inspect())
			// may be innefficient
			optionsMap[arrv] = optM{val: true, obj: oVal.obj}

			check_val = false
			val.Reset()
			continue
		}

		if check_val {
			val.WriteRune(opt)
			continue
		}

		str.WriteRune(opt)
	}

	// A check if they never closed the formatting braces e.g '{0'
	if check_val {
		return "", fmt.Errorf(fmt.Sprintf("Haukufunga '{', tuliokota kabla ya kufika mwisho `%s'", val.String()))
	}

	// Another innefficient loop
	for _, v := range optionsMap {
		if !v.val {
			return "", fmt.Errorf(fmt.Sprintf("Ulipeana hili chaguo (%s) {%s} lakini haukutumia", v.obj.Inspect(), v.obj.Type()))
		}
	}

	// !start
	// Here we can talk about wtf just happened
	// 3 loops to do just formatting, formatting for codes sake.
	// it can be done in 2 loops, the last one is just to confirm that you didn't forget anything.
	// this is not required but a nice syntatic sugar, we are not here for speed, so ergonomics matter instead of speed
	// finally we can say we are slower than python!, what an achievement
	// done!

	return str.String(), nil
}
