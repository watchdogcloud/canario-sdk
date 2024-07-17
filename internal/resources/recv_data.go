package resources

import (
	"fmt"

	"github.com/watchdogcloud/canario/internal/constants"
	"github.com/watchdogcloud/canario/internal/requests"
)

type RecvData struct {
	Req *requests.Request
}

func (recvData *RecvData) PushMetricsToServer(payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.API_VERSION, constants.METRIC_COLLECTION_DEFAULT_ENDPOINT)
	return recvData.Req.Post(url, payload, extraHeaders)
}
