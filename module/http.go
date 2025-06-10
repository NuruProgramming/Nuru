package module

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/NuruProgramming/Nuru/object"
)

// --- MODULE EXPORTS ---

var HTTPFunctions = map[string]object.ModuleFunction{}

// --- USAGE EXAMPLES ---
//
// // Client (default):
// jibu = pata("https://example.com")
//
// // Custom client:
// mteja = undaMteja({mudaWaMwisho: 5000})
// jibu = mteja.pata("https://example.com")
//
// // Custom request:
// ombi = undaOmbi("POST", "https://example.com/api", "{json}")
// jibu = mteja.tekeleza(ombi)
//
// // Server:
// sajiliNjia("GET", "/salamu", fun(jibu, ombi) {
//   jibu.andika("Habari dunia!")
// })
// seva = undaServer()
// seva.sikiliza(":8080")
//
// // See documentation for more details.
//
// --- END OF SWAHILI GO-IDIOMATIC HTTP API ---

func init() {
	HTTPFunctions["undaMteja"] = undaMteja
	HTTPFunctions["undaOmbi"] = undaOmbi
	HTTPFunctions["sajiliNjia"] = sajiliNjia
	HTTPFunctions["undaServer"] = undaServer

	// Register default client helpers at top-level for http.pata, etc.
	defaultClient := undaMteja(nil, nil).(*object.Dict)
	HTTPFunctions["pata"] = func(args []object.Object, defs map[string]object.Object) object.Object {
		if fnObj, ok := defaultClient.Get(&object.String{Value: "pata"}); ok {
			return fnObj.(*object.Builtin).Fn(args...)
		}
		return &object.Error{Message: "'pata' method not found on default client"}
	}
	HTTPFunctions["kichwa"] = func(args []object.Object, defs map[string]object.Object) object.Object {
		if fnObj, ok := defaultClient.Get(&object.String{Value: "kichwa"}); ok {
			return fnObj.(*object.Builtin).Fn(args...)
		}
		return &object.Error{Message: "'kichwa' method not found on default client"}
	}
	HTTPFunctions["tuma"] = func(args []object.Object, defs map[string]object.Object) object.Object {
		if fnObj, ok := defaultClient.Get(&object.String{Value: "tuma"}); ok {
			return fnObj.(*object.Builtin).Fn(args...)
		}
		return &object.Error{Message: "'tuma' method not found on default client"}
	}
	HTTPFunctions["weka"] = func(args []object.Object, defs map[string]object.Object) object.Object {
		if fnObj, ok := defaultClient.Get(&object.String{Value: "weka"}); ok {
			return fnObj.(*object.Builtin).Fn(args...)
		}
		return &object.Error{Message: "'weka' method not found on default client"}
	}
	HTTPFunctions["futa"] = func(args []object.Object, defs map[string]object.Object) object.Object {
		if fnObj, ok := defaultClient.Get(&object.String{Value: "futa"}); ok {
			return fnObj.(*object.Builtin).Fn(args...)
		}
		return &object.Error{Message: "'futa' method not found on default client"}
	}
	HTTPFunctions["bandika"] = func(args []object.Object, defs map[string]object.Object) object.Object {
		if fnObj, ok := defaultClient.Get(&object.String{Value: "bandika"}); ok {
			return fnObj.(*object.Builtin).Fn(args...)
		}
		return &object.Error{Message: "'bandika' method not found on default client"}
	}
	HTTPFunctions["chaguzi"] = func(args []object.Object, defs map[string]object.Object) object.Object {
		if fnObj, ok := defaultClient.Get(&object.String{Value: "chaguzi"}); ok {
			return fnObj.(*object.Builtin).Fn(args...)
		}
		return &object.Error{Message: "'chaguzi' method not found on default client"}
	}
	// Optionally add tumaFomu if implemented
}

// --- SWAHILI GO-IDIOMATIC HTTP CLIENT/SERVER API ---
//
// High-level helpers (default client): pata, kichwa, tuma, weka, futa, bandika, chaguzi, tumaFomu
// Custom client: undaMteja({chaguo})
// Request: undaOmbi(njia, url, chaguo?)
// Server: undaServer({chaguo}), sajiliNjia(njia, mpangilio, mshikaji), sikiliza(anuani)
//
// All objects and methods use Swahili names and Go idioms.

// --- CLIENT ---

type GoBox struct {
	Value interface{}
}

func (g *GoBox) Type() object.ObjectType { return "GOBOX" }
func (g *GoBox) Inspect() string        { return "<GoBox>" }

type Mteja struct {
	Client *http.Client
}

// undaMteja creates a custom HTTP client
func undaMteja(args []object.Object, defs map[string]object.Object) object.Object {
	var chaguo map[string]interface{} = map[string]interface{}{}
	if len(args) > 0 {
		if dict, ok := args[0].(*object.Dict); ok {
			chaguo = convertObjectToWhatever(dict).(map[string]interface{})
		}
	}
	client := &http.Client{}
	if timeout, ok := chaguo["mudaWaMwisho"]; ok {
		if t, ok := timeout.(int64); ok {
			client.Timeout = time.Duration(t) * time.Millisecond
		}
	}
	mteja := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	// Store Go client in _go field as GoBox
	mteja.Set(&object.String{Value: "_go"}, &GoBox{Value: &Mteja{Client: client}})
	addClientMethods(mteja)
	return mteja
}

func addClientMethods(mteja *object.Dict) {
	mteja.Set(&object.String{Value: "pata"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return clientHelperMethod(mteja, "GET", args)
	}})
	mteja.Set(&object.String{Value: "kichwa"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return clientHelperMethod(mteja, "HEAD", args)
	}})
	mteja.Set(&object.String{Value: "tuma"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return clientHelperMethod(mteja, "POST", args)
	}})
	mteja.Set(&object.String{Value: "weka"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return clientHelperMethod(mteja, "PUT", args)
	}})
	mteja.Set(&object.String{Value: "futa"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return clientHelperMethod(mteja, "DELETE", args)
	}})
	mteja.Set(&object.String{Value: "bandika"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return clientHelperMethod(mteja, "PATCH", args)
	}})
	mteja.Set(&object.String{Value: "chaguzi"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return clientHelperMethod(mteja, "OPTIONS", args)
	}})
	mteja.Set(&object.String{Value: "tekeleza"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) < 1 {
			return &object.Error{Message: "tekeleza inahitaji Ombi"}
		}
		ombi, ok := args[0].(*object.Dict)
		if !ok {
			return &object.Error{Message: "Ombi lazima iwe dict"}
		}
		goVal, ok := mteja.Get(&object.String{Value: "_go"})
		if !ok {
			return &object.Error{Message: "Hakuna mteja wa Go"}
		}
		gobox, ok := goVal.(*GoBox)
		if !ok {
			return &object.Error{Message: "Hakuna mteja wa Go"}
		}
		mt, ok := gobox.Value.(*Mteja)
		if !ok {
			return &object.Error{Message: "Hakuna mteja wa Go"}
		}
		req, err := buildGoRequestFromOmbi(ombi)
		if err != nil {
			return &object.Error{Message: err.Error()}
		}
		resp, err := mt.Client.Do(req)
		if err != nil {
			return &object.Error{Message: err.Error()}
		}
		defer resp.Body.Close()
		return createClientResponseObject(resp)
	}})
}

func clientHelperMethod(mteja *object.Dict, method string, args []object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: method + " inahitaji URL"}
	}
	url, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "URL lazima iwe neno"}
	}
	var body io.Reader
	if method == "POST" || method == "PUT" || method == "PATCH" {
		if len(args) < 2 {
			return &object.Error{Message: method + " inahitaji mwili"}
		}
		mwili, ok := args[1].(*object.String)
		if !ok {
			return &object.Error{Message: "Mwili lazima iwe neno"}
		}
		body = strings.NewReader(mwili.Value)
	}
	goVal, ok := mteja.Get(&object.String{Value: "_go"})
	if !ok {
		return &object.Error{Message: "Hakuna mteja wa Go"}
	}
	gobox, ok := goVal.(*GoBox)
	if !ok {
		return &object.Error{Message: "Hakuna mteja wa Go"}
	}
	mt, ok := gobox.Value.(*Mteja)
	if !ok {
		return &object.Error{Message: "Hakuna mteja wa Go"}
	}
	req, err := http.NewRequest(method, url.Value, body)
	if err != nil {
		return &object.Error{Message: err.Error()}
	}
	resp, err := mt.Client.Do(req)
	if err != nil {
		return &object.Error{Message: err.Error()}
	}
	defer resp.Body.Close()
	return createClientResponseObject(resp)
}

// undaOmbi creates a new HTTP request object
func undaOmbi(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "undaOmbi inahitaji njia na URL"}
	}
	njia, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Njia lazima iwe neno"}
	}
	url, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "URL lazima iwe neno"}
	}
	var mwili io.Reader
	if len(args) > 2 {
		if s, ok := args[2].(*object.String); ok {
			mwili = strings.NewReader(s.Value)
		}
	}
	req, err := http.NewRequest(njia.Value, url.Value, mwili)
	if err != nil {
		return &object.Error{Message: err.Error()}
	}
	return createRequestObject(req)
}

func buildGoRequestFromOmbi(ombi *object.Dict) (*http.Request, error) {
	njiaObj, ok := ombi.Get(&object.String{Value: "njia"})
	if !ok {
		return nil, fmt.Errorf("ombi haina njia")
	}
	njia, ok := njiaObj.(*object.String)
	if !ok {
		return nil, fmt.Errorf("njia lazima iwe neno")
	}
	urlObj, ok := ombi.Get(&object.String{Value: "url"})
	if !ok {
		return nil, fmt.Errorf("ombi haina url")
	}
	url, ok := urlObj.(*object.String)
	if !ok {
		return nil, fmt.Errorf("URL lazima iwe neno")
	}
	var mwili io.Reader
	if mwiliObj, ok := ombi.Get(&object.String{Value: "mwili"}); ok {
		if s, ok := mwiliObj.(*object.String); ok {
			mwili = strings.NewReader(s.Value)
		}
	}
	req, err := http.NewRequest(njia.Value, url.Value, mwili)
	if err != nil {
		return nil, err
	}
	if vichwaObj, ok := ombi.Get(&object.String{Value: "vichwa"}); ok {
		if dict, ok := vichwaObj.(*object.Dict); ok {
			for _, pair := range dict.Pairs {
				if k, ok := pair.Key.(*object.String); ok {
					req.Header.Set(k.Value, pair.Value.Inspect())
				}
			}
		}
	}
	return req, nil
}

// --- SERVER ---

// NOTE: mux is global and not safe for concurrent handler registration. TODO: allow per-server mux.
var mux = http.NewServeMux()

// Refactor sajiliNjia to register by path only and check method inside handler
func sajiliNjia(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 3 {
		return &object.Error{Message: "sajiliNjia inahitaji njia, mpangilio, mshikaji"}
	}
	njia, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Njia lazima iwe neno"}
	}
	mpangilio, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Mpangilio lazima iwe neno"}
	}
	mshikaji, ok := args[2].(*object.Function)
	if !ok {
		return &object.Error{Message: "Mshikaji lazima iwe function"}
	}
	fmt.Println("DEBUG: Registering route", mpangilio.Value, "for method", njia.Value)
	mux.HandleFunc(mpangilio.Value, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("DEBUG: Handler entered for", r.Method, r.URL.Path)
		if r.Method != njia.Value {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Njia haikubaliki: %s", r.Method)
			return
		}
		jibu := createResponseObject(w)
		ombi := createRequestObject(r)
		// Call mshikaji(jibu, ombi) in Nuru evaluator
		callFunction(mshikaji, []object.Object{jibu, ombi})
	})
	return &object.Null{}
}

// undaServer creates a custom HTTP server
func undaServer(args []object.Object, defs map[string]object.Object) object.Object {
	seva := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	seva.Set(&object.String{Value: "sikiliza"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		anuani := ":8080"
		if len(args) > 0 {
			if s, ok := args[0].(*object.String); ok {
				anuani = s.Value
			}
		}
		server := &http.Server{
			Addr:    anuani,
			Handler: mux,
		}
		// BLOCKING: Run ListenAndServe in the main thread
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Hitilafu kwenye server: %s\n", err)
		}
		return seva
	}})
	seva.Set(&object.String{Value: "funga"}, &object.Builtin{Fn: func(args ...object.Object) object.Object {
		// Not implemented: need to keep reference to server
		return &object.Null{}
	}})
	return seva
}

// Create a request object from an HTTP request
func createRequestObject(r *http.Request) *object.Dict {
	reqObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Add HTTP method
	reqObj.Set(&object.String{Value: "njia"}, &object.String{Value: r.Method})

	// Add URL as string (always)
	reqObj.Set(&object.String{Value: "url"}, &object.String{Value: r.URL.String()})

	// Optionally add path as 'mpangilio' (route pattern)
	reqObj.Set(&object.String{Value: "mpangilio"}, &object.String{Value: r.URL.Path})

	// Add headers
	headers := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	for key, values := range r.Header {
		headers.Set(&object.String{Value: key}, &object.String{Value: strings.Join(values, ", ")})
	}
	reqObj.Set(&object.String{Value: "vichwa"}, headers)

	// Read body
	if r.Body != nil {
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		// Create a new reader for the body so it can be read again
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		if err == nil {
			reqObj.Set(&object.String{Value: "mwili"}, &object.String{Value: string(body)})
		}
	}

	return reqObj
}

// Create a response object for an HTTP response writer
func createResponseObject(w http.ResponseWriter) *object.Dict {
	resObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	// Use a pointer to bool for headersSent to ensure per-response state
	headersSent := new(bool)

	// Add write method
	resObj.Set(&object.String{Value: "andika"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "Andika inahitaji angalau hoja moja"}
			}
			data := args[0].Inspect()
			if !*headersSent {
				w.Header().Set("Content-Type", "text/plain")
				*headersSent = true
			}
			fmt.Fprint(w, data)
			return &object.Null{}
		},
	})

	// Add writeHead method
	resObj.Set(&object.String{Value: "andikaKichwa"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "andikaKichwa inahitaji msimbo wa hali"}
			}
			statusCode := 200
			headers := make(map[string]string)
			// Get status code
			if statusCodeObj, ok := args[0].(*object.Integer); ok {
				statusCode = int(statusCodeObj.Value)
			}
			// Get headers if provided
			if len(args) > 1 {
				if headersObj, ok := args[1].(*object.Dict); ok {
					for _, pair := range headersObj.Pairs {
						if key, ok := pair.Key.(*object.String); ok {
							headers[key.Value] = pair.Value.Inspect()
						}
					}
				}
			}
			// Set the headers
			for key, value := range headers {
				w.Header().Set(key, value)
			}
			// Write the status code
			w.WriteHeader(statusCode)
			*headersSent = true
			return &object.Null{}
		},
	})

	// Add end method
	resObj.Set(&object.String{Value: "mwisho"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			fmt.Println("DEBUG: mwisho called with args:", args)
			if len(args) > 0 {
				data := args[0].Inspect()
				fmt.Println("DEBUG: mwisho writing data:", data)
				if !*headersSent {
					w.Header().Set("Content-Type", "text/plain")
					w.WriteHeader(200)
					*headersSent = true
				}
				n, err := w.Write([]byte(data))
				fmt.Println("DEBUG: w.Write returned:", n, err)
			}
			return &object.Null{}
		},
	})

	// Add setHeader method
	resObj.Set(&object.String{Value: "wekaKichwa"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return &object.Error{Message: "wekaKichwa inahitaji jina na thamani"}
			}
			name, ok := args[0].(*object.String)
			if !ok {
				return &object.Error{Message: "Jina la kichwa lazima liwe neno"}
			}
			value, ok := args[1].(*object.String)
			if !ok {
				return &object.Error{Message: "Thamani ya kichwa lazima liwe neno"}
			}
			w.Header().Set(name.Value, value.Value)
			return &object.Null{}
		},
	})

	// Add status code property
	resObj.Set(&object.String{Value: "msimboWaHali"}, &object.Integer{Value: 200})

	return resObj
}

// Create a response object for an HTTP client response
func createClientResponseObject(resp *http.Response) *object.Dict {
	resObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Add basic properties
	resObj.Set(&object.String{Value: "msimboWaHali"}, &object.Integer{Value: int64(resp.StatusCode)})
	resObj.Set(&object.String{Value: "ujumbeWaHali"}, &object.String{Value: resp.Status})

	// Add headers
	headers := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	for key, values := range resp.Header {
		headers.Set(&object.String{Value: key}, &object.String{Value: strings.Join(values, ", ")})
	}
	resObj.Set(&object.String{Value: "vichwa"}, headers)

	// Add body
	var mwiliStr string
	if resp.Body != nil {
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err == nil {
			mwiliStr = string(body)
		} else {
			mwiliStr = ""
		}
	} else {
		mwiliStr = ""
	}

	resObj.Set(&object.String{Value: "mwili"}, &object.String{Value: mwiliStr})
	// DEBUG: Print all keys in resObj when setting 'mwili'
	fmt.Println("DEBUG: Setting 'mwili' in response object. Current keys:")
	for k := range resObj.Pairs {
		fmt.Println("  key:", k.Type, ", value:", resObj.Pairs[k].Key.Inspect(), ", type:", resObj.Pairs[k].Value.Type())
	}

	// Add URL
	if resp.Request != nil && resp.Request.URL != nil {
		resObj.Set(&object.String{Value: "url"}, &object.String{Value: resp.Request.URL.String()})
	}

	return resObj
}

