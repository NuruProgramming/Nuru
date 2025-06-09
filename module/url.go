package module

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/NuruProgramming/Nuru/object"
)

var URLFunctions = map[string]object.ModuleFunction{}

func init() {
	URLFunctions["changanua"] = parseURL                 // Parse URL strings
	URLFunctions["URL"] = newURL                         // Constructor for URL objects
	URLFunctions["tengeneza"] = formatURL                // Format URL objects to strings
	URLFunctions["tatua"] = resolveURL                   // Resolve URLs
	URLFunctions["URLSearchParams"] = newURLSearchParams // For query string manipulation
}

// parseURL parses a URL string and returns a URL object
func parseURL(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Changanua inahitaji hoja ya kwanza (URL)"}
	}

	urlString, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la URL"}
	}

	// Parse query parameters?
	parseQuery := false
	if len(args) > 1 {
		if queryArg, ok := args[1].(*object.Boolean); ok {
			parseQuery = queryArg.Value
		}
	}

	// Parse the URL
	parsedURL, err := url.Parse(urlString.Value)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Hitilafu kuchanganua URL: %s", err)}
	}

	// Create a URL object
	return createURLObject(parsedURL, parseQuery)
}

// newURL creates a new URL object (constructor)
func newURL(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "URL inahitaji hoja ya kwanza (URL)"}
	}

	input, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la URL"}
	}

	var baseURL string
	if len(args) > 1 {
		if baseArg, ok := args[1].(*object.String); ok {
			baseURL = baseArg.Value
		}
	}

	var parsedURL *url.URL
	var err error

	if baseURL != "" {
		// If base URL is provided, resolve against it
		base, err := url.Parse(baseURL)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Hitilafu kuchanganua URL msingi: %s", err)}
		}
		parsedURL, err = base.Parse(input.Value)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Hitilafu kuchanganua URL: %s", err)}
		}
	} else {
		// Otherwise parse as absolute URL
		parsedURL, err = url.Parse(input.Value)
	}

	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Hitilafu kuchanganua URL: %s", err)}
	}

	// Create a URL object
	return createURLObject(parsedURL, false)
}

// formatURL formats a URL object back to a string
func formatURL(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Tengeneza inahitaji hoja ya kwanza (URL object)"}
	}

	// Accept either a URL object (Dict) or a parsed URL object
	var urlString string

	switch arg := args[0].(type) {
	case *object.Dict:
		// Extract components from Dict and build URL
		protocol := ""
		if protocolObj, ok := arg.Get(&object.String{Value: "protocol"}); ok {
			if protocolStr, ok := protocolObj.(*object.String); ok {
				protocol = protocolStr.Value
				// Make sure protocol ends with ":"
				if !strings.HasSuffix(protocol, ":") {
					protocol += ":"
				}
			}
		}

		auth := ""
		if usernameObj, ok := arg.Get(&object.String{Value: "username"}); ok {
			if username, ok := usernameObj.(*object.String); ok && username.Value != "" {
				auth = username.Value
				if passwordObj, ok := arg.Get(&object.String{Value: "password"}); ok {
					if password, ok := passwordObj.(*object.String); ok && password.Value != "" {
						auth += ":" + password.Value
					}
				}
				auth += "@"
			}
		}

		host := ""
		if hostObj, ok := arg.Get(&object.String{Value: "host"}); ok {
			if hostStr, ok := hostObj.(*object.String); ok {
				host = hostStr.Value
			}
		}

		pathname := ""
		if pathnameObj, ok := arg.Get(&object.String{Value: "pathname"}); ok {
			if pathnameStr, ok := pathnameObj.(*object.String); ok {
				pathname = pathnameStr.Value
				// Ensure pathname starts with "/"
				if pathname != "" && !strings.HasPrefix(pathname, "/") {
					pathname = "/" + pathname
				}
			}
		}

		search := ""
		if searchObj, ok := arg.Get(&object.String{Value: "search"}); ok {
			if searchStr, ok := searchObj.(*object.String); ok {
				search = searchStr.Value
				// Ensure search starts with "?"
				if search != "" && !strings.HasPrefix(search, "?") {
					search = "?" + search
				}
			}
		}

		hash := ""
		if hashObj, ok := arg.Get(&object.String{Value: "hash"}); ok {
			if hashStr, ok := hashObj.(*object.String); ok {
				hash = hashStr.Value
				// Ensure hash starts with "#"
				if hash != "" && !strings.HasPrefix(hash, "#") {
					hash = "#" + hash
				}
			}
		}

		// Combine components
		urlString = protocol + "//"
		if auth != "" {
			urlString += auth
		}
		urlString += host + pathname + search + hash
	default:
		return &object.Error{Message: "Tengeneza inahitaji URL object"}
	}

	return &object.String{Value: urlString}
}

// resolveURL resolves a URL against a base URL
func resolveURL(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "Tatua inahitaji hoja mbili: URL msingi na URL ya kutatua"}
	}

	from, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe neno la URL msingi"}
	}

	to, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya pili lazima iwe neno la URL ya kutatua"}
	}

	// Parse base URL
	baseURL, err := url.Parse(from.Value)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Hitilafu kuchanganua URL msingi: %s", err)}
	}

	// Resolve relative URL
	resolvedURL, err := baseURL.Parse(to.Value)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Hitilafu kutatua URL: %s", err)}
	}

	// Return as string
	return &object.String{Value: resolvedURL.String()}
}

// createURLObject creates a URL dictionary object from a parsed URL
func createURLObject(parsedURL *url.URL, parseQuery bool) *object.Dict {
	urlObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Set URL components
	protocol := parsedURL.Scheme
	if protocol != "" && !strings.HasSuffix(protocol, ":") {
		protocol += ":"
	}
	urlObj.Set(&object.String{Value: "protocol"}, &object.String{Value: protocol})

	urlObj.Set(&object.String{Value: "host"}, &object.String{Value: parsedURL.Host})
	urlObj.Set(&object.String{Value: "hostname"}, &object.String{Value: parsedURL.Hostname()})
	urlObj.Set(&object.String{Value: "port"}, &object.String{Value: parsedURL.Port()})
	urlObj.Set(&object.String{Value: "pathname"}, &object.String{Value: parsedURL.Path})
	urlObj.Set(&object.String{Value: "search"}, &object.String{Value: parsedURL.RawQuery})
	urlObj.Set(&object.String{Value: "hash"}, &object.String{Value: parsedURL.Fragment})
	urlObj.Set(&object.String{Value: "href"}, &object.String{Value: parsedURL.String()})

	// Handle auth
	if parsedURL.User != nil {
		urlObj.Set(&object.String{Value: "username"}, &object.String{Value: parsedURL.User.Username()})
		password, exists := parsedURL.User.Password()
		if exists {
			urlObj.Set(&object.String{Value: "password"}, &object.String{Value: password})
		} else {
			urlObj.Set(&object.String{Value: "password"}, &object.String{Value: ""})
		}
	} else {
		urlObj.Set(&object.String{Value: "username"}, &object.String{Value: ""})
		urlObj.Set(&object.String{Value: "password"}, &object.String{Value: ""})
	}

	// Origin (protocol + host)
	origin := ""
	if parsedURL.Scheme != "" {
		origin = parsedURL.Scheme + "://" + parsedURL.Host
	}
	urlObj.Set(&object.String{Value: "origin"}, &object.String{Value: origin})

	// Parse query if requested
	if parseQuery && parsedURL.RawQuery != "" {
		query := parsedURL.Query()
		queryObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

		for key, values := range query {
			if len(values) == 1 {
				queryObj.Set(&object.String{Value: key}, &object.String{Value: values[0]})
			} else if len(values) > 1 {
				// Create array for multiple values
				valueArray := &object.Array{Elements: []object.Object{}}
				for _, v := range values {
					valueArray.Elements = append(valueArray.Elements, &object.String{Value: v})
				}
				queryObj.Set(&object.String{Value: key}, valueArray)
			}
		}

		urlObj.Set(&object.String{Value: "query"}, queryObj)
	}

	// Add searchParams method that returns a URLSearchParams object
	urlObj.Set(&object.String{Value: "searchParams"}, createURLSearchParams(parsedURL.RawQuery))

	// Add toString method
	urlObj.Set(&object.String{Value: "toString"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.String{Value: parsedURL.String()}
		},
	})

	// Add toJSON method (same as toString in Node.js)
	urlObj.Set(&object.String{Value: "toJSON"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.String{Value: parsedURL.String()}
		},
	})

	return urlObj
}

// Helper function to create a URLSearchParams object
func createURLSearchParams(queryString string) *object.Dict {
	// First, ensure queryString doesn't have a leading ?
	queryString = strings.TrimPrefix(queryString, "?")

	return newURLSearchParamsFromString(queryString)
}

// newURLSearchParams creates a new URLSearchParams object
func newURLSearchParams(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) == 0 {
		// No args - create empty params
		return newURLSearchParamsFromString("")
	}

	switch arg := args[0].(type) {
	case *object.String:
		// From string
		queryString := arg.Value
		// Remove leading ? if present
		queryString = strings.TrimPrefix(queryString, "?")
		return newURLSearchParamsFromString(queryString)

	case *object.Dict:
		// From dictionary
		return newURLSearchParamsFromDict(arg)

	case *object.Array:
		// From array of [key, value] pairs
		return newURLSearchParamsFromArray(arg)

	default:
		return &object.Error{Message: "URLSearchParams inakubali neno, kamusi, au safu tu"}
	}
}

// Create URLSearchParams from string
func newURLSearchParamsFromString(queryString string) *object.Dict {
	paramsObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	params := make(map[string][]string)

	// Parse the query string
	if queryString != "" {
		for _, pair := range strings.Split(queryString, "&") {
			kv := strings.SplitN(pair, "=", 2)
			key := kv[0]
			value := ""
			if len(kv) > 1 {
				value = kv[1]
			}

			// URL-decode key and value
			var err error
			key, err = url.QueryUnescape(key)
			if err != nil {
				// Skip invalid encoding
				continue
			}

			if len(kv) > 1 {
				value, err = url.QueryUnescape(value)
				if err != nil {
					// Skip invalid encoding
					continue
				}
			}

			params[key] = append(params[key], value)
		}
	}

	// Store the parsed params
	paramsObj.Set(&object.String{Value: "_params"}, &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)})
	paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

	for key, values := range params {
		valuesArray := &object.Array{Elements: []object.Object{}}
		for _, value := range values {
			valuesArray.Elements = append(valuesArray.Elements, &object.String{Value: value})
		}
		paramsDict.Set(&object.String{Value: key}, valuesArray)
	}

	// Add methods
	addURLSearchParamsMethods(paramsObj)

	return paramsObj
}

// Create URLSearchParams from dictionary
func newURLSearchParamsFromDict(dict *object.Dict) *object.Dict {
	paramsObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	paramsObj.Set(&object.String{Value: "_params"}, &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)})
	paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

	// Convert each key-value pair
	for _, pair := range dict.Pairs {
		if keyStr, ok := pair.Key.(*object.String); ok {
			// Create array of values (even if just one)
			valuesArray := &object.Array{Elements: []object.Object{}}

			// Add the value
			switch val := pair.Value.(type) {
			case *object.Array:
				// Array of values
				valuesArray.Elements = val.Elements
			default:
				// Single value
				valuesArray.Elements = append(valuesArray.Elements, &object.String{Value: val.Inspect()})
			}

			paramsDict.Set(keyStr, valuesArray)
		}
	}

	// Add methods
	addURLSearchParamsMethods(paramsObj)

	return paramsObj
}

// Create URLSearchParams from array of [key, value] pairs
func newURLSearchParamsFromArray(arr *object.Array) *object.Dict {
	paramsObj := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	paramsObj.Set(&object.String{Value: "_params"}, &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)})
	paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

	// Process each [key, value] pair
	for _, item := range arr.Elements {
		if pairArray, ok := item.(*object.Array); ok && len(pairArray.Elements) >= 2 {
			key := pairArray.Elements[0].Inspect()
			value := pairArray.Elements[1].Inspect()

			keyObj := &object.String{Value: key}

			// Get existing values array or create new one
			var valuesArray *object.Array
			if existing, ok := paramsDict.Get(keyObj); ok {
				valuesArray = existing.(*object.Array)
			} else {
				valuesArray = &object.Array{Elements: []object.Object{}}
				paramsDict.Set(keyObj, valuesArray)
			}

			// Add the value
			valuesArray.Elements = append(valuesArray.Elements, &object.String{Value: value})
		}
	}

	// Add methods
	addURLSearchParamsMethods(paramsObj)

	return paramsObj
}

// Add methods to URLSearchParams object
func addURLSearchParamsMethods(paramsObj *object.Dict) {
	// append method
	paramsObj.Set(&object.String{Value: "ongeza"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return &object.Error{Message: "Ongeza inahitaji jina na thamani"}
			}

			key, ok := args[0].(*object.String)
			if !ok {
				return &object.Error{Message: "Jina lazima liwe neno"}
			}

			value := args[1].Inspect()

			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Get existing values array or create new one
			var valuesArray *object.Array
			if existing, ok := paramsDict.Get(key); ok {
				valuesArray = existing.(*object.Array)
			} else {
				valuesArray = &object.Array{Elements: []object.Object{}}
				paramsDict.Set(key, valuesArray)
			}

			// Add the value
			valuesArray.Elements = append(valuesArray.Elements, &object.String{Value: value})

			return paramsObj
		},
	})

	// delete method
	paramsObj.Set(&object.String{Value: "futa"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "Futa inahitaji jina"}
			}

			key, ok := args[0].(*object.String)
			if !ok {
				return &object.Error{Message: "Jina lazima liwe neno"}
			}

			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Delete the key
			paramsDict.Delete(key)

			return paramsObj
		},
	})

	// get method
	paramsObj.Set(&object.String{Value: "pata"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "Pata inahitaji jina"}
			}

			key, ok := args[0].(*object.String)
			if !ok {
				return &object.Error{Message: "Jina lazima liwe neno"}
			}

			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Get the values array
			if values, ok := paramsDict.Get(key); ok {
				valuesArray := values.(*object.Array)
				if len(valuesArray.Elements) > 0 {
					return valuesArray.Elements[0]
				}
			}

			return &object.Null{}
		},
	})

	// getAll method
	paramsObj.Set(&object.String{Value: "pataZote"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "PataZote inahitaji jina"}
			}

			key, ok := args[0].(*object.String)
			if !ok {
				return &object.Error{Message: "Jina lazima liwe neno"}
			}

			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Get the values array
			if values, ok := paramsDict.Get(key); ok {
				return values.(*object.Array)
			}

			return &object.Array{Elements: []object.Object{}}
		},
	})

	// has method
	paramsObj.Set(&object.String{Value: "ina"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "Ina inahitaji jina"}
			}

			key, ok := args[0].(*object.String)
			if !ok {
				return &object.Error{Message: "Jina lazima liwe neno"}
			}

			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Check if the key exists
			_, exists := paramsDict.Get(key)
			return &object.Boolean{Value: exists}
		},
	})

	// set method
	paramsObj.Set(&object.String{Value: "weka"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return &object.Error{Message: "Weka inahitaji jina na thamani"}
			}

			key, ok := args[0].(*object.String)
			if !ok {
				return &object.Error{Message: "Jina lazima liwe neno"}
			}

			value := args[1].Inspect()

			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Create a new values array with just this value
			valuesArray := &object.Array{Elements: []object.Object{
				&object.String{Value: value},
			}}

			// Set the key to the new values array
			paramsDict.Set(key, valuesArray)

			return paramsObj
		},
	})

	// toString method
	paramsObj.Set(&object.String{Value: "kuNeno"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			var parts []string

			// Build query string parts
			for _, pair := range paramsDict.Pairs {
				key, ok := pair.Key.(*object.String)
				if !ok {
					continue
				}

				valuesArray, ok := pair.Value.(*object.Array)
				if !ok {
					continue
				}

				for _, valueObj := range valuesArray.Elements {
					value := valueObj.Inspect()

					// URL-encode key and value
					encodedKey := url.QueryEscape(key.Value)
					encodedValue := url.QueryEscape(value)

					parts = append(parts, encodedKey+"="+encodedValue)
				}
			}

			// Join with &
			return &object.String{Value: strings.Join(parts, "&")}
		},
	})

	// sort method
	paramsObj.Set(&object.String{Value: "panga"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Get all keys
			var keys []string
			for _, pair := range paramsDict.Pairs {
				if key, ok := pair.Key.(*object.String); ok {
					keys = append(keys, key.Value)
				}
			}

			// Sort keys
			sort.Strings(keys)

			// Create a new sorted dictionary
			sortedDict := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

			// Add keys in sorted order
			for _, key := range keys {
				keyObj := &object.String{Value: key}
				if values, ok := paramsDict.Get(keyObj); ok {
					sortedDict.Set(keyObj, values)
				}
			}

			// Replace the params dictionary
			paramsObj.Set(&object.String{Value: "_params"}, sortedDict)

			return paramsObj
		},
	})

	// keys method
	paramsObj.Set(&object.String{Value: "funguo"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Create an array of keys
			keys := []object.Object{}

			// Add each key (potentially multiple times for repeated keys)
			for _, pair := range paramsDict.Pairs {
				key, ok := pair.Key.(*object.String)
				if !ok {
					continue
				}

				valuesArray, ok := pair.Value.(*object.Array)
				if !ok {
					continue
				}

				// Add key once for each value
				for range valuesArray.Elements {
					keys = append(keys, key)
				}
			}

			return &object.Array{Elements: keys}
		},
	})

	// values method
	paramsObj.Set(&object.String{Value: "thamani"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Create an array of values
			values := []object.Object{}

			// Add each value
			for _, pair := range paramsDict.Pairs {
				valuesArray, ok := pair.Value.(*object.Array)
				if !ok {
					continue
				}

				values = append(values, valuesArray.Elements...)
			}

			return &object.Array{Elements: values}
		},
	})

	// entries method
	paramsObj.Set(&object.String{Value: "viingilio"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Get the params dictionary
			paramsDict := paramsObj.Pairs[(&object.String{Value: "_params"}).HashKey()].Value.(*object.Dict)

			// Create an array of [key, value] pairs
			entries := []object.Object{}

			// Add each entry
			for _, pair := range paramsDict.Pairs {
				key, ok := pair.Key.(*object.String)
				if !ok {
					continue
				}

				valuesArray, ok := pair.Value.(*object.Array)
				if !ok {
					continue
				}

				for _, valueObj := range valuesArray.Elements {
					entry := &object.Array{Elements: []object.Object{
						key,
						valueObj,
					}}
					entries = append(entries, entry)
				}
			}

			return &object.Array{Elements: entries}
		},
	})

	// forEach method
	paramsObj.Set(&object.String{Value: "kila"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "Kila inahitaji function"}
			}

			_, ok := args[0].(*object.Function)
			if !ok {
				return &object.Error{Message: "Hoja ya kwanza lazima iwe function"}
			}

			// This is a simplified version that can't actually execute the callback due to import cycle
			// In a real implementation, we would need to evaluate the callback for each entry

			return paramsObj
		},
	})
}
