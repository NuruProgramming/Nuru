package builtins

import (
	"fmt"
	"github.com/AvicennaJr/Nuru/object"
	"strings"
	"unicode"
)

type fmtSt struct {
	Thamani string
	Aina    string
}

type fmtSpec struct {
	Specifier string
	Handler   func(flags string, value fmtSt) (string, string)
}

func Andika(args ...object.Object) object.Object {
	if len(args) == 0 {
		fmt.Println("")
		return nil
	}

	var fmtv []fmtSt
	for _, arg := range args {
		if arg == nil {
			return newError("Hauwezi kufanya operesheni hii")
		}
		fmtv = append(fmtv, fmtSt{Thamani: arg.Inspect(), Aina: Aina(arg).Inspect()}) // Append the values (type and value)
	}

	passedType := fmtv[0].Aina
	if passedType == "NENO" {
		fmts := fmtv[1:]                                 // Choose from index 1, this is to skip the 'format' string
		drrr, err := formatString(fmtv[0].Thamani, fmts) // call the formatter to do its thing :)
		// Handle errors, they are returned in string format so that I can call the newError
		// which takes in a string as the first argument.
		// This also allows for a central point of handling (or failure in some cases)
		if err != "" {
			return newError(err)
		}
		fmt.Println(drrr) // print the formatted string
		return nil
	}

	// Incase the first argument wasn't a string, then do what was already the default.
	// This is to preserve some kind of 'backward compatibility'
	str := fmtv[0].Thamani // Temporary solution
	fmt.Println(str)
	return nil
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
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
	srf := "Safu mbadala ni chini ya safu halisi"
	frf := "Siwezi kupata fomati inayofaa"
    var fmtsp = []fmtSpec{
        {"n", handleNamba},
        {"t", handleTungo},
		{"b", handleBuliani},
		{"s", handleSafu},
    }
    for _, char := range fmts {
        if char == '%' {
            insideFmt = true
            continue
        }
        if insideFmt {
            if unicode.IsLetter(char) {
				if pos >= len(fmtv) {
						return "", srf
				}
                for _, spec := range fmtsp {
                    if char == rune(spec.Specifier[0]) {
						val, err := spec.Handler(flags.String(), fmtv[pos])
						if err != "" {
								return "", err
						}
                        newStr.WriteString(val)
                        pos++
                        insideFmt = false
                        mismatch = false
                        break
                    }
                }
                if insideFmt { // If no match was found, it's a mismatch
                    mismatch = true
                }
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
    return newStr.String(), ""
}

func handleNamba(flags string, value fmtSt) (string, string) {
		mrf := "Mismatch ya Fomati"
		aina := "NAMBA"
		// First check if the type is correct
		if value.Aina != aina {
				return "", fmt.Sprintf("%s: %s si %s", mrf, value.Aina, aina)
		}

		// Check if the flags is empty, if so return the value directly
		if flags == "" {
				return value.Thamani, ""
		}

		return "", ""
}

func handleTungo(flags string, value fmtSt) (string, string) {
		mrf := "Mismatch ya Fomati"
		aina := "NENO"
		// First check if the type is correct
		if value.Aina != aina {
				return "", fmt.Sprintf("%s: %s si %s", mrf, value.Aina, aina)
		}

		// Check if the flags is empty, if so return the value directly
		if flags == "" {
				return value.Thamani, ""
		}


		return "",""
}
func handleBuliani(flags string, value fmtSt) (string, string) {
		mrf := "Mismatch ya Fomati"
		aina := "BOOLEAN"
		// First check if the type is correct
		if value.Aina != aina {
				return "", fmt.Sprintf("%s: %s si %s", mrf, value.Aina, aina)
		}

		// Check if the flags is empty, if so return the value directly
		if flags == "" {
				return value.Thamani, ""
		}


		return "",""
}

func handleSafu(flags string, value fmtSt) (string, string) {
		mrf := "Mismatch ya Fomati"
		aina := "ORODHA"
		// First check if the type is correct
		if value.Aina != aina {
				return "", fmt.Sprintf("%s: %s si %s", mrf, value.Aina, aina)
		}

		// Check if the flags is empty, if so return the value directly
		if flags == "" {
				return value.Thamani, ""
		}


		return "",""
}
