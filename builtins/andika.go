package builtins

import (
	"fmt"
	"github.com/AvicennaJr/Nuru/object"
	"strings"
)

func Andika(args ...object.Object) object.Object {
	if len(args) == 0 {
		fmt.Println("")
		return nil
	}

	var arr []string
	var types []string
	for _, arg := range args {
		if arg == nil {
			return newError("Hauwezi kufanya operesheni hii")
		}
		arr = append(arr, arg.Inspect())           // Append the arguments (strings) into an array
		types = append(types, Aina(arg).Inspect()) // Append the types of the above args (string format) to an array
	}

	passedType := types[0]
	if passedType == "NENO" {
		fmts := arr[1:] // Choose from index 1, this is to skip the 'format' string
		types = types[1:]
		drrr, err := formatString(arr[0], fmts, types) // call the formatter to do its thing :)
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
	str := strings.Join(arr, " ")
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
// The vals array contains the array with the actual values to be displayed.
// The types array contains the types that will be matched before being
//
//	displayed. If a mismatch happens, then it will be used to return an error
//
// Two strings are returned:
//   - The first string is the actual formatted string
//   - The second string returns an 'error' string.
//   - It should have been an actual error but the context usage won't allow it.
func formatString(fmts string, vals []string, types []string) (string, string) {
	var newStr strings.Builder
	pos := 0 // The position in the passed arrays
	lv := len(vals)

	// Check if the vals and types array match
	if lv != len(types) {
		return "", "Orodha ya majibu hailingani, ripoti hii shida tafadhali"
	}
	// Define the error messages
	srf := "Safu mbadala ni chini ya safu halisi"
	frf := "Siwezi kupata fomati inayofaa"
	mrf := "Mismatch ya Fomati"

	// Loop through the string passed.
	for i := 0; i < len(fmts); i++ {
		if fmts[i] == '%' && i+1 < len(fmts) {
			switch fmts[i+1] {
			case 't':
				if pos < lv {
					if types[pos] == "NENO" { // Check if it is what is says
						newStr.WriteString(vals[pos]) // Pull the value from the array
						pos++                         // Move the position forwards
						i++                           // Move the string forwards
						continue                      // Skip the format specifier
					} else { // Check if there was a mismatch in types
						fmrf := fmt.Sprintf("%s: %s si NENO", mrf, types[pos])
						return "", fmrf
					}
				} else {
					return "", srf // Check if there was a mismatch in length
				}
			case 'n':
				if pos < lv {
					if types[pos] == "NAMBA" {
						newStr.WriteString(vals[pos])
						pos++
						i++
						continue
					} else {
						fmrf := fmt.Sprintf("%s: %s si NAMBARI", mrf, types[pos])
						return "", fmrf
					}
				} else {
					return "", srf
				}
			case 's':
				if pos < lv {
					if types[pos] == "ORODHA" {
						newStr.WriteString(vals[pos])
						pos++
						i++
						continue
					} else {
						fmrf := fmt.Sprintf("%s: %s si ORODHA", mrf, types[pos])
						return "", fmrf
					}
				} else {
					return "", srf
				}
			case 'b':
				if pos < lv {
					if types[pos] == "BOOLEAN" {
						newStr.WriteString(vals[pos])
						pos++
						i++
						continue
					} else {
						fmrf := fmt.Sprintf("%s: %s si BOOLEAN", mrf, types[pos])
						return "", fmrf
					}
				} else {
					return "", srf
				}
			default:
				return "", frf // Check for mismatch in format specifier
			}
		} else {
			newStr.WriteByte(fmts[i])
		}
	}
	return newStr.String(), ""
}
