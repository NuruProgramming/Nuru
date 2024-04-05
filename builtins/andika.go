package builtins

import (
	"fmt"
	"github.com/AvicennaJr/Nuru/object"
	"strconv"
	"strings"
	"unicode"
)

// Struct for the format 'format' that is passed around to the 'formatter'.
type fmtSt struct {
	Thamani string // The value of the 'replacement' e.g 123
	Aina    string // The type of the 'replacement' e.g NAMBA
}

// Flags passed for easier handling of relevant flags
// Acceptable flags: [ ., +, !] ('.' has to have proceding integers)
type fmtFlags struct {
	DotFlag  bool
	DotValue int
	PlusFlag bool
	EXCFlag  bool // Not really a good flag
}

// Define the common error messages

var (
	// 'Standard' error messages
	srf = "Safu mbadala ni chini ya safu halisi"
	frf = "Siwezi kupata fomati inayofaa"
	hrf = "Safu halisi ni chini ya safu mbadala"
	mrf = "Mismatch ya Fomati"
)

// Andika - Write string to the `stdout` (whatever that is).
// If an error is encountered, it is returned as early as possible.
// For the `object` object, check the relevant file(s)
func Andika(args ...object.Object) object.Object {
	if len(args) == 0 {
		fmt.Println("")
		return nil
	}

	var fmtv []fmtSt
	// Loop over the arguments passed in and manipulate them accordingly.
	for _, arg := range args {
		if arg == nil {
			return newError("Hauwezi kufanya operesheni hii")
		}
		// Append the values (type and value)
		fmtv = append(fmtv, fmtSt{Thamani: arg.Inspect(), Aina: Aina(arg).Inspect()})
	}

	// Check the first part of the argument passed what type it is.
	// If it is a string, then attempt to do string manipulation as it
	// may contain the relevant 'strings'. Else join the value passed.
	// The `fmtv` is assumed that it will always have at least one value.
	// This can be slightly guranteed by the entry check but more checks should
	// be put in place just in case.
	if fmtv[0].Aina == "NENO" {
		fmts := fmtv[1:]                                 // Choose from index 1, this is to skip the 'format' string
		drrr, err := formatString(fmtv[0].Thamani, fmts) // call the formatter to do its thing :)
		// Handle errors, they are returned in string format so that I can call the newError
		// which takes in a string as the first argument.
		// This also allows for a central point of handling (or failure in some cases)
		if err != "" {
			return newError(err)
		}
		fmt.Printf(drrr) // print the formatted string
		return nil
	}

	// Incase the first argument wasn't a string, then do what was already the default.
	// This is to preserve some kind of 'backward compatibility'
	var strd strings.Builder
	for i := 0; i < len(fmtv); i++ {
		strd.WriteString(fmtv[i].Thamani)
		strd.WriteString(" ")
	}
	fmt.Println(strd.String())
	return nil
}

// Format and then subsequntly return an error message
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// Define a map where keys are format specifiers and values are handler functions.
var formatHandlers = map[rune]func(flags fmtFlags, value fmtSt) (string, string){
	'n': handleNamba,
	't': handleTungo,
	'b': handleBuliani,
	's': handleSafu,
	'd': handleDesimali,
}

// Format the string according to predefined characters
// Takes in a string -> This contains the format that the partern will follow.
//
//	This follows the already established convention (the format part not what is used)
//	To accommodate the swahili speakers that this is geered to, the following formats are used:
//		- %n -> Nambari
//		- %t -> Tungo (ama neno)
//		- %b -> Buliani (kweli, sikweli)
//		- %s -> Safu (ama orodha)
//		- %d -> Desimali
//	This is what can safely be handled by now.
//
//	displayed. If a mismatch happens, then it will be used to return an error
//
// Two strings are returned:
//   - The first string is the actual formatted string
//   - The second string returns an 'error' string.
//   - It should have been an actual error but the context usage won't allow it.
func formatString(fmts string, fmtv []fmtSt) (string, string) {
	var newStr strings.Builder
	var flags strings.Builder
	var pos int
	var insideFmt bool
	var mismatch bool

	lf := len(fmtv)

	for _, char := range fmts {
		if char == '%' {
			insideFmt = true
			mismatch = true
			continue
		}
		if insideFmt {
			if unicode.IsLetter(char) {
				if pos >= lf {
					return "", srf // Check for mismatch in the number of format specifiers and values.
				}
				handler, ok := formatHandlers[char]
				if !ok {
					return "", frf // Format specifier not found.
				}
				ff := setFlags(flags.String())
				val, err := handler(ff, fmtv[pos])
				if err != "" {
					return "", err
				}
				newStr.WriteString(val) // Append the formatted value.
				pos++
				insideFmt = false
				mismatch = false
			} else {
				flags.WriteRune(char)
			}
		} else {
			if mismatch {
				return "", frf
			}
			newStr.WriteRune(char)
		}
	}
	if mismatch {
		return "", frf
	}
	if pos < lf {
		return "", hrf // Check if there are more values than format specifiers.
	}
	return newStr.String(), ""
}

// Only parse flags in logical order, any screw ups
// on the part of the caller should be silently ignored.
// This means that most of errors will lead to undefined behaviour
func setFlags(flags string) fmtFlags {
	var ff fmtFlags
	var dvh bool
	var dv strings.Builder

	for _, val := range flags {
		if val == '.' {
			dvh = true
			ff.DotFlag = true
			continue
		}
		if !unicode.IsNumber(val) {
			dvh = false
		}
		if !dvh {
			switch val {
			case '+':
				ff.PlusFlag = true
			case '!':
				ff.EXCFlag = true
			}
		} else {
			dv.WriteRune(val)
		}
	}

	if dv.Len() > 0 {
		ff.DotValue, _ = strconv.Atoi(dv.String())
	}
	return ff
}

func handleNamba(flags fmtFlags, value fmtSt) (string, string) {
	rets := value.Thamani
	aina := "NAMBA"
	// First check if the type is correct
	if value.Aina != aina {
		return "", fmt.Sprintf("%s: %s si %s", mrf, value.Aina, aina)
	}

	if flags.PlusFlag {
		rets = handlePlus(rets)
	}

	return rets, ""
}

func handleDesimali(flags fmtFlags, value fmtSt) (string, string) {
	rets := value.Thamani
	aina := "DESIMALI"
	// First check if the type is correct
	if value.Aina != aina {
		return "", fmt.Sprintf("%s: %s si %s", mrf, value.Aina, aina)
	}

	if flags.PlusFlag {
		rets = handlePlus(rets)
	}

	if flags.DotFlag {
		ads := strings.Split(rets, ".")[1]
		bds := strings.Split(rets, ".")[0]
		rets = bds + "." + handleDot(flags.DotValue, ads)
	}

	return rets, ""
}

func handleTungo(flags fmtFlags, value fmtSt) (string, string) {
	rets := value.Thamani
	aina := "NENO"
	// First check if the type is correct
	if value.Aina != aina {
		return "", fmt.Sprintf("%s: %s si %s", mrf, value.Aina, aina)
	}

	if flags.EXCFlag {
		rets = handleExc(rets)
	}

	return rets, ""
}
func handleBuliani(flags fmtFlags, value fmtSt) (string, string) {
	aina := "BOOLEAN"
	// First check if the type is correct
	if value.Aina != aina {
		return "", fmt.Sprintf("%s: %s si %s", mrf, value.Aina, aina)
	}

	// Since bool doesn't accept any flags, just ignore them
	return value.Thamani, ""
}

func handleSafu(flags fmtFlags, value fmtSt) (string, string) {
	aina := "ORODHA"
	// First check if the type is correct
	if value.Aina != aina {
		return "", fmt.Sprintf("%s: %s si %s", mrf, value.Aina, aina)
	}

	return value.Thamani, ""
}

func handlePlus(value string) string {
	// Check if it already contains a '+' or negative
	if !(strings.Contains(value, "+") || strings.Contains(value, "-")) {
		value = fmt.Sprintf("+%s", value)
	}

	return value
}

func handleDot(precision int, value string) string {
	lv := len(value)

	if lv < precision {
		for i := lv; i < precision; i++ {
			value += "0"
		}
	}

	return value
}

func handleExc(val string) string {
	return strings.ToUpper(val)
}
