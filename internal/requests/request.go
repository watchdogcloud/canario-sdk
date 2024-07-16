package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/zakhaev26/canario/internal/constants"
	"github.com/zakhaev26/canario/internal/errors"
)

type Auth struct {
	K string
	V string
}

const (
	TIMEOUT_S = 10
)

type Request struct {
	Auth       Auth
	HTTPClient *http.Client
	Headers    map[string]string
	Version    string
	SDKName    string
	AppDetails string
	BASE_URI   string
}

// Utility func that appends query params in the URL
// eg : http://127.0.0.1:3030?deez=nuts
// encoded query values, without '?'
func BuildURLWithParams(requestURL string, data map[string]interface{}) string {

	URL, err := url.Parse(requestURL)

	if err != nil {
		panic(err)
	}
	parameters := url.Values{}

	for k, v := range data {
		parameters.Add(k, fmt.Sprintf("%v", v))
	}

	URL.RawQuery = parameters.Encode()
	return URL.String()
}

func (request *Request) AppendHeaders(headers map[string]string) {
	for k, v := range headers {
		request.Headers[k] = v
	}
}

func (request *Request) AddRequestHeadersInternal(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		if key == "Content-Type" || key == "User-Agent" {
			continue
		}
		req.Header.Set(key, value)
	}
}

func (request *Request) AddRequestHeaders(req *http.Request, headers map[string]string, contentType ...string) {
	//Set the Defaults First in case unavailable
	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", request.SDKName, request.Version))

	if len(contentType) == 0 {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", contentType[0])
	}

	// Set the already added headers
	request.AddRequestHeadersInternal(req, request.Headers)
	request.AddRequestHeadersInternal(req, headers)
}

func (request *Request) SetTimeout(timeout int16) {
	timeoutSeconds := int64(timeout) * int64(time.Second)
	request.HTTPClient = &http.Client{Timeout: time.Duration(timeoutSeconds)}
}

func ProcessResponse(response *http.Response) (map[string]interface{}, error) {
	//avoid mem leak in client
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if len(body) == 0 {
		resp := make(map[string]interface{})
		return resp, nil
	}

	if err != nil {
		return nil, err
	}

	resp := make(map[string]interface{})

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (request *Request) APICall(req *http.Request) (map[string]interface{}, error) {

	client := request.HTTPClient
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= constants.HTTP_STATUS_OK &&
		resp.StatusCode < constants.HTTP_STATUS_REDIRECT {
		return ProcessResponse(resp)
	}

	// resp is errorfull...

	var jsonResp errors.CanarioErrorJSON
	json.NewDecoder(resp.Body).Decode(&jsonResp)

	errorData := jsonResp.ErrorData

	switch errorData.InternalErrorCode {
	case constants.SERVER_ERROR:
		return nil, &errors.ServerError{Message: errorData.Description}
	case constants.GATEWAY_ERROR:
		return nil, &errors.GatewayError{Message: errorData.Description}
	case constants.BAD_REQUEST_ERROR:
	default:
		return nil, &errors.BadRequestError{Message: errorData.Description}
	}
	return ProcessResponse(resp)
}
