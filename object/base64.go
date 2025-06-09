package object

import (
	"encoding/base64"
	"fmt"
)

// Define the Base64 object type constant in object.go
// const BASE64_OBJ = "BASE64"

// Base64 represents a base64 encoded data in the Nuru language
type Base64 struct {
	// The raw data (decoded)
	Data []byte
	// The encoded string (Base64)
	Encoded string
}

// Type returns the type of the object
func (b *Base64) Type() ObjectType { return BASE64_OBJ }

// Inspect returns a string representation of the object
func (b *Base64) Inspect() string { return b.Encoded }

// Method implements method calls on Base64 objects
func (b *Base64) Method(method string, args []Object) Object {
	switch method {
	case "kukata": // decode
		return b.decode(args)
	case "kutoka": // from data
		return b.fromData(args)
	case "data": // get raw data
		return b.getRawData(args)
	default:
		return newError("Samahani, kiendesha '%s' hakitumiki na Base64", method)
	}
}

// decode is a method that decodes a Base64 string
func (b *Base64) decode(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	
	return &String{Value: string(b.Data)}
}

// fromData is a method that creates a Base64 from raw data
func (b *Base64) fromData(args []Object) Object {
	if len(args) != 1 {
		return newError("Samahani, tunahitaji Hoja 1, wewe umeweka %d", len(args))
	}
	
	var data []byte
	
	switch arg := args[0].(type) {
	case *String:
		data = []byte(arg.Value)
	case *Byte:
		data = arg.Value
	default:
		return newError("Samahani, tunahitaji Neno au Byte, wewe umeweka %s", args[0].Type())
	}
	
	encoded := base64.StdEncoding.EncodeToString(data)
	
	return &Base64{
		Data:    data,
		Encoded: encoded,
	}
}

// getRawData returns the raw data as a string object
func (b *Base64) getRawData(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	
	return &String{Value: string(b.Data)}
}

// EncodeToBase64 encodes data to Base64
func EncodeToBase64(data []byte) (*Base64, error) {
	encoded := base64.StdEncoding.EncodeToString(data)
	return &Base64{
		Data:    data,
		Encoded: encoded,
	}, nil
}

// DecodeFromBase64 decodes a Base64 string
func DecodeFromBase64(encoded string) (*Base64, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %v", err)
	}
	
	return &Base64{
		Data:    data,
		Encoded: encoded,
	}, nil
}
