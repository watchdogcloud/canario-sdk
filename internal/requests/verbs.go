package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (req *Request) Get(path string, query map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s%s", req.BASE_URI, path)

	url = BuildURLWithParams(url, query)

	_req, _ := http.NewRequest("GET", url, nil)

	internalHeader := make(map[string]string)
	internalHeader[req.Auth.K] = req.Auth.V
	req.AddRequestHeadersInternal(_req, internalHeader)

	req.AddRequestHeaders(_req, extraHeaders)

	return req.APICall(_req)
}

func (req *Request) Post(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	jsonBytes, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s%s", req.BASE_URI, path)

	_req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))

	internalHeader := make(map[string]string)
	internalHeader[req.Auth.K] = req.Auth.V
	req.AddRequestHeadersInternal(_req, extraHeaders)
	req.AddRequestHeaders(_req, extraHeaders)

	return req.APICall(_req)
}

func (req *Request) Patch(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	jsonBytes, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s%s", req.BASE_URI, path)

	_req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBytes))

	internalHeader := make(map[string]string)
	internalHeader[req.Auth.K] = req.Auth.V
	req.AddRequestHeadersInternal(_req, extraHeaders)

	return req.APICall(_req)
}

func (req *Request) Put(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	jsonBytes, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s%s", req.BASE_URI, path)

	_req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytes))

	internalHeader := make(map[string]string)
	internalHeader[req.Auth.K] = req.Auth.V
	req.AddRequestHeadersInternal(_req, extraHeaders)

	req.AddRequestHeaders(_req, extraHeaders)

	return req.APICall(_req)
}

func (req *Request) Delete(path string, query map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s%s", req.BASE_URI, path)
	url = BuildURLWithParams(url, query)

	_req, _ := http.NewRequest("PATCH", url, nil)

	internalHeader := make(map[string]string)
	internalHeader[req.Auth.K] = req.Auth.V
	req.AddRequestHeadersInternal(_req, extraHeaders)

	req.AddRequestHeaders(_req, extraHeaders)

	return req.APICall(_req)
}
