package module

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/NuruProgramming/Nuru/object"
)

var NetFunctions = map[string]object.ModuleFunction{}

func init() {
	NetFunctions["peruzi"] = getRequest
	NetFunctions["tuma"] = postRequest
}

func getRequest(args []object.Object, defs map[string]object.Object) object.Object {

	if len(defs) != 0 {
		var url *object.String
		var headers, params *object.Dict
		for k, v := range defs {
			switch k {
			case "yuareli":
				strUrl, ok := v.(*object.String)
				if !ok {
					return &object.Error{Message: "Yuareli iwe neno"}
				}
				url = strUrl
			case "vichwa":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Vichwa lazima viwe kamusi"}
				}
				headers = dictHead
			case "mwili":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Mwili lazima iwe kamusi"}
				}
				params = dictHead
			default:
				return &object.Error{Message: "Hoja si sahihi. Tumia yuareli na vichwa."}
			}
		}
		if url.Value == "" {
			return &object.Error{Message: "Yuareli ni lazima"}
		}

		var responseBody *bytes.Buffer
		if params != nil {
			booty := convertObjectToWhatever(params)

			jsonBody, err := json.Marshal(booty)

			if err != nil {
				return &object.Error{Message: "Huku format query yako vizuri."}
			}

			responseBody = bytes.NewBuffer(jsonBody)
		}

		var req *http.Request
		var err error
		if responseBody != nil {
			req, err = http.NewRequest("GET", url.Value, responseBody)
		} else {
			req, err = http.NewRequest("GET", url.Value, nil)
		}
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kufanya request"}
		}

		if headers != nil {
			for _, val := range headers.Pairs {
				req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
			}
		}
		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Tumeshindwa kutuma request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kusoma majibu."}
		}

		return &object.String{Value: string(respBody)}

	}

	if len(args) == 1 {
		url, ok := args[0].(*object.String)
		if !ok {
			return &object.Error{Message: "Yuareli lazima iwe neno"}
		}
		req, err := http.NewRequest("GET", url.Value, nil)
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kufanya request"}
		}

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Tumeshindwa kutuma request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kusoma majibu."}
		}

		return &object.String{Value: string(respBody)}
	}
	return &object.Error{Message: "Hoja si sahihi. Tumia yuareli na vichwa."}
}

func postRequest(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		var url *object.String
		var headers, params *object.Dict
		for k, v := range defs {
			switch k {
			case "yuareli":
				strUrl, ok := v.(*object.String)
				if !ok {
					return &object.Error{Message: "Yuareli iwe neno"}
				}
				url = strUrl
			case "vichwa":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Vichwa lazima viwe kamusi"}
				}
				headers = dictHead
			case "mwili":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Mwili lazima iwe kamusi"}
				}
				params = dictHead
			default:
				return &object.Error{Message: "Hoja si sahihi. Tumia yuareli na vichwa."}
			}
		}
		if url.Value == "" {
			return &object.Error{Message: "Yuareli ni lazima"}
		}
		var responseBody *bytes.Buffer
		if params != nil {
			booty := convertObjectToWhatever(params)

			jsonBody, err := json.Marshal(booty)

			if err != nil {
				return &object.Error{Message: "Huku format query yako vizuri."}
			}

			responseBody = bytes.NewBuffer(jsonBody)
		}
		var req *http.Request
		var err error
		if responseBody != nil {
			req, err = http.NewRequest("POST", url.Value, responseBody)
		} else {
			req, err = http.NewRequest("POST", url.Value, nil)
		}
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kufanya request"}
		}
		if headers != nil {
			for _, val := range headers.Pairs {
				req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
			}
		}
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Tumeshindwa kutuma request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kusoma majibu."}
		}
		return &object.String{Value: string(respBody)}
	}
	return &object.Error{Message: "Hoja si sahihi. Tumia yuareli na vichwa."}
}
