package module

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/NuruProgramming/Nuru/object"
)

var HTTPFunctions = map[string]object.ModuleFunction{}

// The Agent class equivalent
type Agent struct {
	maxSockets      int
	keepAlive       bool
	keepAliveMsecs  int
	maxFreeSockets  int
	sockets         map[string][]*http.Client
	freeSockets     map[string][]*http.Client
	requests        map[string][]*http.Request
	defaultProtocol string
}

// Initialize the module functions
func init() {
	HTTPFunctions["tengenezaServer"] = createServer
	HTTPFunctions["tengenezaOmbi"] = createRequest
	HTTPFunctions["pata"] = get
	HTTPFunctions["STATUS_CODES"] = getStatusCodes
	HTTPFunctions["METHODS"] = getMethods

	// Global agent
	HTTPFunctions["agentiDunia"] = getGlobalAgent
}

// createServer creates an HTTP server
// tengenezaServer = unda(mshikaji) { ... }
func createServer(args []object.Object, defs map[string]object.Object) object.Object {
	// Check for callback
	if len(args) < 1 {
		return &object.Error{Message: "Lazima upe function ya kutunza maombi"}
	}

	// Check if the argument is a function
	_, ok := args[0].(*object.Function)
	if !ok {
		return &object.Error{Message: "Hoja lazima iwe function"}
	}

	// Create a new server with methods
	serverObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Add listen method to start the server
	serverObj.Set(&object.String{Value: "sikiliza"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			port := 8080
			host := "localhost"

			// Process the arguments
			if len(args) >= 1 {
				switch arg := args[0].(type) {
				case *object.Integer:
					port = int(arg.Value)
				case *object.Dict:
					// Check for port
					if portObj, ok := arg.Get(&object.String{Value: "port"}); ok {
						if portInt, ok := portObj.(*object.Integer); ok {
							port = int(portInt.Value)
						}
					}
					// Check for host
					if hostObj, ok := arg.Get(&object.String{Value: "mwenyeji"}); ok {
						if hostStr, ok := hostObj.(*object.String); ok {
							host = hostStr.Value
						}
					}
				}
			}

			// Add port and host to the server object
			serverObj.Set(&object.String{Value: "port"}, &object.Integer{Value: int64(port)})
			serverObj.Set(&object.String{Value: "mwenyeji"}, &object.String{Value: host})

			// Start the HTTP server in a goroutine
			go func() {
				addr := fmt.Sprintf("%s:%d", host, port)

				// Create an HTTP server
				server := &http.Server{
					Addr: addr,
					Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						// Create request and response objects for Nuru
						reqObj := createRequestObject(r)
						resObj := createResponseObject(w)

						// Since we can't call evaluator directly (import cycle),
						// just store the objects in environment for now
						// In a real implementation, we would pass these to an evaluation system
						env := object.NewEnvironment()
						env.Set("ombi", reqObj)
						env.Set("jibu", resObj)

						// Note: Actual evaluation would happen here,
						// but can't directly call evaluator due to import cycle
						// For demonstration, we'll just send a simple response
						w.Header().Set("Content-Type", "text/html")
						w.WriteHeader(200)
						fmt.Fprint(w, "<h1>Karibu kwenye Nuru HTTP Server!</h1>")
					}),
				}

				// Set the server object to listen state
				serverObj.Set(&object.String{Value: "anasikiliza"}, &object.Boolean{Value: true})

				// Start the server
				err := server.ListenAndServe()
				if err != nil {
					fmt.Printf("Hitilafu kwenye server: %s\n", err)
					serverObj.Set(&object.String{Value: "anasikiliza"}, &object.Boolean{Value: false})
				}
			}()

			return serverObj
		},
	})

	// Add close method to stop the server
	serverObj.Set(&object.String{Value: "funga"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// In a real implementation, you'd need to keep a reference to the server to stop it
			serverObj.Set(&object.String{Value: "anasikiliza"}, &object.Boolean{Value: false})
			return serverObj
		},
	})

	// Add default properties
	serverObj.Set(&object.String{Value: "anasikiliza"}, &object.Boolean{Value: false})
	serverObj.Set(&object.String{Value: "timeout"}, &object.Integer{Value: 120000}) // 2 minutes default timeout

	return serverObj
}

// Create a request object from an HTTP request
func createRequestObject(r *http.Request) *object.Dict {
	reqObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Add basic properties
	reqObj.Set(&object.String{Value: "njia"}, &object.String{Value: r.Method})

	// Create URL object using our URL module (if available)
	if URLFunctions != nil {
		// Create arguments for URL function
		args := []object.Object{&object.String{Value: r.URL.String()}}
		defs := make(map[string]object.Object)

		// Call URL constructor
		if urlFunction, ok := URLFunctions["changanua"]; ok {
			urlObj := urlFunction(args, defs)
			reqObj.Set(&object.String{Value: "url"}, urlObj)
		} else {
			// Fallback if URL module not available
			reqObj.Set(&object.String{Value: "url"}, &object.String{Value: r.URL.String()})
		}
	} else {
		// Fallback if URL module not available
		reqObj.Set(&object.String{Value: "url"}, &object.String{Value: r.URL.String()})
	}

	// Still set the path for backward compatibility
	reqObj.Set(&object.String{Value: "njia"}, &object.String{Value: r.URL.Path})

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
	headersSent := false

	// Add write method
	resObj.Set(&object.String{Value: "andika"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "Andika inahitaji angalau hoja moja"}
			}

			data := args[0].Inspect()
			if !headersSent {
				w.Header().Set("Content-Type", "text/plain")
				headersSent = true
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
			// statusMsg variable is intentionally unused since WriteHeader doesn't take a message
			// but we want to keep the Node.js-like API that accepts it
			var _ string
			headers := make(map[string]string)

			// Get status code
			if statusCodeObj, ok := args[0].(*object.Integer); ok {
				statusCode = int(statusCodeObj.Value)
			}

			// Get status message if provided (intentionally ignored)
			if len(args) > 1 {
				if statusMsgObj, ok := args[1].(*object.String); ok {
					_ = statusMsgObj.Value
				}
			}

			// Get headers if provided
			if len(args) > 2 {
				if headersObj, ok := args[2].(*object.Dict); ok {
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
			headersSent = true

			return &object.Null{}
		},
	})

	// Add end method
	resObj.Set(&object.String{Value: "mwisho"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 0 {
				data := args[0].Inspect()
				if !headersSent {
					w.Header().Set("Content-Type", "text/plain")
					headersSent = true
				}
				fmt.Fprint(w, data)
			}

			// In a real implementation, this would need to actually close the response
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

// createRequest creates an HTTP request (similar to http.request in Node.js)
func createRequest(args []object.Object, defs map[string]object.Object) object.Object {
	options := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Process arguments
	if len(args) > 0 {
		// First argument could be options, URL string, or URL object
		switch arg := args[0].(type) {
		case *object.Dict:
			// Check if this is a URL object by looking for URL properties
			if _, ok := arg.Get(&object.String{Value: "protocol"}); ok {
				// This is likely a URL object
				if urlStr, ok := arg.Get(&object.String{Value: "href"}); ok {
					options.Set(&object.String{Value: "url"}, urlStr)
				}
			} else {
				// Regular options dictionary
				options = arg
			}
		case *object.String:
			options.Set(&object.String{Value: "url"}, arg)
		}
	}

	// Process options from named parameters (defs)
	for k, v := range defs {
		key := &object.String{Value: k}
		options.Set(key, v)
	}

	// Create a request object that mimics Node.js's ClientRequest
	reqObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Extract options
	method := "GET"
	url := ""
	headers := make(map[string]string)
	body := ""

	// Get method
	if methodObj, ok := options.Get(&object.String{Value: "njia"}); ok {
		if methodStr, ok := methodObj.(*object.String); ok {
			method = methodStr.Value
		}
	}

	// Get URL
	if urlObj, ok := options.Get(&object.String{Value: "url"}); ok {
		url = urlObj.Inspect()
	}

	// Get headers
	if headersObj, ok := options.Get(&object.String{Value: "vichwa"}); ok {
		if headerDict, ok := headersObj.(*object.Dict); ok {
			for _, pair := range headerDict.Pairs {
				if key, ok := pair.Key.(*object.String); ok {
					headers[key.Value] = pair.Value.Inspect()
				}
			}
		}
	}

	// Add write method to the request
	reqObj.Set(&object.String{Value: "andika"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "Andika inahitaji data"}
			}

			// Append to body
			body += args[0].Inspect()
			return reqObj
		},
	})

	// Add end method to the request
	reqObj.Set(&object.String{Value: "mwisho"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Create and send the actual HTTP request
			var httpReq *http.Request
			var err error

			if len(body) > 0 {
				httpReq, err = http.NewRequest(method, url, bytes.NewBufferString(body))
			} else {
				httpReq, err = http.NewRequest(method, url, nil)
			}

			if err != nil {
				return &object.Error{Message: fmt.Sprintf("Hitilafu kutengeneza ombi: %s", err)}
			}

			// Set headers
			for key, value := range headers {
				httpReq.Header.Set(key, value)
			}

			// Create a client and send the request
			client := &http.Client{}
			resp, err := client.Do(httpReq)

			if err != nil {
				// Emit error event
				if errEventObj, ok := reqObj.Get(&object.String{Value: "onError"}); ok {
					// Check if the event handler is a function, even though we can't call it directly
					if _, ok := errEventObj.(*object.Function); ok {
						// Store error in environment but can't directly call evaluator
						env := object.NewEnvironment()
						env.Set("kosa", &object.String{Value: err.Error()})

						// In a real implementation, would call evaluator with the environment
						fmt.Printf("Hitilafu: %s\n", err.Error())
					}
				}
				return &object.Error{Message: fmt.Sprintf("Hitilafu kutuma ombi: %s", err)}
			}

			// Create a response object
			resObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

			// Add basic properties
			resObj.Set(&object.String{Value: "msimboWaHali"}, &object.Integer{Value: int64(resp.StatusCode)})
			resObj.Set(&object.String{Value: "ujumbeWaHali"}, &object.String{Value: resp.Status})

			// Add headers
			respHeaders := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
			for key, values := range resp.Header {
				respHeaders.Set(&object.String{Value: key}, &object.String{Value: strings.Join(values, ", ")})
			}
			resObj.Set(&object.String{Value: "vichwa"}, respHeaders)

			// Create URL object for the response URL (if URL module available)
			if URLFunctions != nil {
				if urlFunction, ok := URLFunctions["changanua"]; ok {
					urlObj := urlFunction([]object.Object{&object.String{Value: resp.Request.URL.String()}}, nil)
					resObj.Set(&object.String{Value: "url"}, urlObj)
				} else {
					resObj.Set(&object.String{Value: "url"}, &object.String{Value: resp.Request.URL.String()})
				}
			} else {
				resObj.Set(&object.String{Value: "url"}, &object.String{Value: resp.Request.URL.String()})
			}

			// Add on data event
			resObj.Set(&object.String{Value: "kwenyeData"}, &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) < 1 {
						return &object.Error{Message: "kwenyeData inahitaji function"}
					}

					// In a simplified implementation without direct evaluator access
					if resp.Body != nil {
						body, err := ioutil.ReadAll(resp.Body)
						resp.Body.Close()

						if err == nil {
							fmt.Printf("Data received: %s\n", string(body))
						}
					}

					return resObj
				},
			})

			// Add on end event
			resObj.Set(&object.String{Value: "kwenyeMwisho"}, &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) < 1 {
						return &object.Error{Message: "kwenyeMwisho inahitaji function"}
					}

					// In a simplified implementation, just note this was called
					fmt.Println("Request completed")

					return resObj
				},
			})

			// If callback was provided in the second argument
			if len(args) > 1 {
				if _, ok := args[1].(*object.Function); ok {
					// In a simplified implementation, would call callback
					// but can't directly evaluate due to import cycle
					fmt.Println("Callback would be called here with response")
				}
			}

			// Emit response event
			if respEventObj, ok := reqObj.Get(&object.String{Value: "onResponse"}); ok {
				// Just check the type to avoid unused variable warning
				if _, ok := respEventObj.(*object.Function); ok {
					// In a simplified implementation, would call response handler
					// but can't directly evaluate due to import cycle
					fmt.Println("Response handler would be called here")
				}
			}

			return &object.Null{}
		},
	})

	// Add on event handlers
	reqObj.Set(&object.String{Value: "kwenyeJibu"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "kwenyeJibu inahitaji function"}
			}

			handler, ok := args[0].(*object.Function)
			if !ok {
				return &object.Error{Message: "kwenyeJibu inahitaji function"}
			}

			reqObj.Set(&object.String{Value: "onResponse"}, handler)
			return reqObj
		},
	})

	reqObj.Set(&object.String{Value: "kwenyeHitilafu"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "kwenyeHitilafu inahitaji function"}
			}

			handler, ok := args[0].(*object.Function)
			if !ok {
				return &object.Error{Message: "kwenyeHitilafu inahitaji function"}
			}

			reqObj.Set(&object.String{Value: "onError"}, handler)
			return reqObj
		},
	})

	return reqObj
}

// get is a shorthand for request with GET method
func get(args []object.Object, defs map[string]object.Object) object.Object {
	options := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Process arguments
	if len(args) > 0 {
		// First argument could be options, URL string, or URL object
		switch arg := args[0].(type) {
		case *object.Dict:
			// Check if this is a URL object by looking for URL properties
			if _, ok := arg.Get(&object.String{Value: "protocol"}); ok {
				// This is likely a URL object
				if urlStr, ok := arg.Get(&object.String{Value: "href"}); ok {
					options.Set(&object.String{Value: "url"}, urlStr)
				}
			} else {
				// Regular options dictionary
				options = arg
			}
		case *object.String:
			options.Set(&object.String{Value: "url"}, arg)
		}
	}

	// Set method to GET
	options.Set(&object.String{Value: "njia"}, &object.String{Value: "GET"})

	// Call createRequest with the options
	return createRequest([]object.Object{options}, defs)
}

// getStatusCodes returns HTTP status codes as a dict
func getStatusCodes(args []object.Object, defs map[string]object.Object) object.Object {
	codes := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Add common HTTP status codes
	statusCodes := map[string]string{
		"100": "Continue",
		"101": "Switching Protocols",
		"200": "OK",
		"201": "Created",
		"202": "Accepted",
		"204": "No Content",
		"206": "Partial Content",
		"300": "Multiple Choices",
		"301": "Moved Permanently",
		"302": "Found",
		"304": "Not Modified",
		"307": "Temporary Redirect",
		"308": "Permanent Redirect",
		"400": "Bad Request",
		"401": "Unauthorized",
		"403": "Forbidden",
		"404": "Not Found",
		"405": "Method Not Allowed",
		"406": "Not Acceptable",
		"409": "Conflict",
		"410": "Gone",
		"413": "Payload Too Large",
		"414": "URI Too Long",
		"415": "Unsupported Media Type",
		"429": "Too Many Requests",
		"500": "Internal Server Error",
		"501": "Not Implemented",
		"502": "Bad Gateway",
		"503": "Service Unavailable",
		"504": "Gateway Timeout",
	}

	for code, message := range statusCodes {
		key := &object.String{Value: code}
		value := &object.String{Value: message}
		codes.Set(key, value)
	}

	return codes
}

// getMethods returns HTTP methods as an array
func getMethods(args []object.Object, defs map[string]object.Object) object.Object {
	methods := []object.Object{
		&object.String{Value: "GET"},
		&object.String{Value: "HEAD"},
		&object.String{Value: "POST"},
		&object.String{Value: "PUT"},
		&object.String{Value: "DELETE"},
		&object.String{Value: "CONNECT"},
		&object.String{Value: "OPTIONS"},
		&object.String{Value: "TRACE"},
		&object.String{Value: "PATCH"},
	}

	return &object.Array{Elements: methods}
}

// getGlobalAgent returns a global agent instance
func getGlobalAgent(args []object.Object, defs map[string]object.Object) object.Object {
	agent := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Create an Agent struct to use its fields
	agentStruct := &Agent{
		maxSockets:      5,
		keepAlive:       true,
		keepAliveMsecs:  1000,
		maxFreeSockets:  10,
		sockets:         make(map[string][]*http.Client),
		freeSockets:     make(map[string][]*http.Client),
		requests:        make(map[string][]*http.Request),
		defaultProtocol: "http:",
	}

	// Set default properties using the struct fields
	agent.Set(&object.String{Value: "maxSockets"}, &object.Integer{Value: int64(agentStruct.maxSockets)})
	agent.Set(&object.String{Value: "keepAlive"}, &object.Boolean{Value: agentStruct.keepAlive})
	agent.Set(&object.String{Value: "keepAliveMsecs"}, &object.Integer{Value: int64(agentStruct.keepAliveMsecs)})
	agent.Set(&object.String{Value: "maxFreeSockets"}, &object.Integer{Value: int64(agentStruct.maxFreeSockets)})

	return agent
}

// Helper function to convert Nuru objects to Go types
func httpConvertObjectToWhatever(obj object.Object) interface{} {
	switch obj := obj.(type) {
	case *object.String:
		return obj.Value
	case *object.Integer:
		return obj.Value
	case *object.Float:
		return obj.Value
	case *object.Boolean:
		return obj.Value
	case *object.Array:
		result := make([]interface{}, len(obj.Elements))
		for i, elem := range obj.Elements {
			result[i] = httpConvertObjectToWhatever(elem)
		}
		return result
	case *object.Dict:
		result := make(map[string]interface{})
		for _, pair := range obj.Pairs {
			key := pair.Key.Inspect()
			result[key] = httpConvertObjectToWhatever(pair.Value)
		}
		return result
	default:
		return obj.Inspect()
	}
}
