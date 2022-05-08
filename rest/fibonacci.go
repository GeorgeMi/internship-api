package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful/v3"
)

type GetFibonacciRequest struct {
	Index *int
	Size  int
}

type GetFibonacciResponse struct {
	Response []int `json:"response"`
}

func next_fibonacci(a, b *int) (response int) {
	response = *a + *b
	*a = *b
	*b = response

	return
}

func (r *Service) checkListAndAppend(list []int, val int) []int {
	for _, value := range r.blackList {
		if val == value {
			return list
		}
	}
	list = append(list, val)
	return list
}

func (r *Service) GetFibonacciRequest(request *restful.Request, response *restful.Response) {
	var responseJson GetFibonacciResponse
	//responseJson.Response = []int{1, 3, 5}

	requestQuery, err := readGetFibonacciRequest(request)
	if err != nil {
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("error:%s", err))

		return
	}

	fmt.Printf("index: %v", requestQuery.Index)
	fmt.Printf("size: %v", requestQuery.Size)

	err = validateGetFibonacciRequest(requestQuery)
	if err != nil {
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("error:%s", err))

		return
	}

	s := make([]int, 0) // uint64 poti pune cea mai mare valoare de timp int pozitiv
	s = r.checkListAndAppend(s, 0)
	s = r.checkListAndAppend(s, 1)
	x := 0
	y := 1

	for len(s) < *requestQuery.Index+requestQuery.Size {
		s = r.checkListAndAppend(s, next_fibonacci(&x, &y))
	}

	responseJson.Response = s[*requestQuery.Index : *requestQuery.Index+requestQuery.Size]

	response.WriteAsJson(responseJson)
}

func readGetFibonacciRequest(request *restful.Request) (requestQuery GetFibonacciRequest, err error) {

	requestQuery.Index = new(int)
	if len(request.QueryParameter("index")) > 0 {

		*requestQuery.Index, err = strconv.Atoi(request.QueryParameter("index"))
	}

	if len(request.QueryParameter("size")) > 0 {

		requestQuery.Size, err = strconv.Atoi(request.QueryParameter("size"))
	} else {
		requestQuery.Size = 1
	}

	return
}

func validateGetFibonacciRequest(requestQuery GetFibonacciRequest) error {
	if requestQuery.Index == nil {

		return fmt.Errorf("invalid index: %v", *requestQuery.Index)
	}

	if *requestQuery.Index < 0 {

		return fmt.Errorf("invalid index: %v", *requestQuery.Index)
	}
	if requestQuery.Size <= 0 {

		return fmt.Errorf("invalid ißize: %v", requestQuery.Size)
	}

	return nil //fmt.Errorf("Index invalid")
}
