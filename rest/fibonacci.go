package rest

import "github.com/emicklei/go-restful/v3"

type GetFibonacciRequest struct {
	Index int
}

type GetFibonacciResponse struct {
	Response []int `json:"response"`
}

func (r *Service) GetFibonacciRequest(request *restful.Request, response *restful.Response) {
	var responseJson GetFibonacciResponse
	responseJson.Response = []int{1, 3, 5}

	response.WriteAsJson(responseJson)
}
