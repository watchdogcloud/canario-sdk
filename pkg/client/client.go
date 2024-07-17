package client

import (
	"net/http"
	"time"

	"github.com/watchdogcloud/canario/internal/conf"
	"github.com/watchdogcloud/canario/internal/requests"
	"github.com/watchdogcloud/canario/internal/resources"
	"github.com/watchdogcloud/canario/pkg/versioning"
)

type Client struct {
	RecvData resources.RecvData
}

var Request *requests.Request

func CreateNewClient(key string, secret string) *Client {
	auth := requests.Auth{
		K: key,
		V: secret,
	}

	httpClient := http.Client{
		Timeout: requests.TIMEOUT_S * time.Second,
	}

	Request = &requests.Request{
		Auth:       auth,
		HTTPClient: &httpClient,
		Version:    versioning.GetSDKVersion(),
		SDKName:    versioning.GetSDKVersion(),
		BASE_URI:   conf.CreateNewConf().API.BaseURI,
	}

	recv_data := resources.RecvData{Req: Request}

	client := &Client{
		RecvData: recv_data,
	}

	return client
}

func (client *Client) AddHeaders(headers map[string]string) {
	Request.AppendHeaders(headers)
}

func (client *Client) SetTimeout(timeout int16) {
	Request.SetTimeout(timeout)
}
