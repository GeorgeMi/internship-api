package rest

import (
	"github.com/emicklei/go-restful/v3"
)

type EndpointErrorResponse struct {
	responseErrorStatus `json:"ResponseStatus,omitempty"`
}

type responseErrorStatus struct {
	ErrorStatusCode int    `json:"ErrorStatusCode,omitempty"`
	Message         string `json:"Message,omitempty"`
}

func buildEndPointErrorResponse(resp *restful.Response, httpStatusCode int, errorMsg string) {
	var errResp = EndpointErrorResponse{
		responseErrorStatus{
			ErrorStatusCode: httpStatusCode,
			Message:         errorMsg,
		},
	}
	_ = resp.WriteHeaderAndJson(httpStatusCode, errResp, restful.MIME_JSON)
}
