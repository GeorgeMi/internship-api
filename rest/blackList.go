package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful/v3"
)

type DeleteNumberRequest struct {
	Number *int
}

type DeleteNumberResponse struct {
	Response []int `json:"response"`
}

type GetBlacklistRequest struct {
	Val *int
}

type GetBlacklistResponse struct {
	Response []int `json:"response"`
}

func (r *Service) DeleteNumberRequest(request *restful.Request, response *restful.Response) {
	var responseJson DeleteNumberResponse

	requestQuery, err := readDeleteNumberRequest(request)
	if err != nil {
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("error:%s", err))

		return
	}

	fmt.Printf("number: %v", requestQuery.Number)

	err = validateDeleteNumberRequest(requestQuery)
	if err != nil {
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("error:%s", err))

		return
	}

	for i, w := range r.blackList {
		if *requestQuery.Number == w {
			//r.blackList = append(r.blackList[:i], r.blackList[i+1:])
			responseJson.Response = append(r.blackList[:i], r.blackList[i+1:]...)

		} else {
			buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("number doesn't exist in array"))

			return
		}
	}

	response.WriteAsJson(responseJson)
}

func readAddBlacklistRequest(request *restful.Request) (requestQuery GetBlacklistRequest, err error) {

	requestQuery.Val = new(int)
	request.PathParameter("BlacklistParameter")

	if len(request.PathParameter("BacklistParameter")) > 0 {
		*requestQuery.Val, err = strconv.Atoi(request.PathParameter("BlacklistParameter"))
	}
	return
}

func (r *Service) AddBlacklistElement(request *restful.Request, response *restful.Response) {

	requestQuery, err := readAddBlacklistRequest(request)
	if err != nil {
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("error:%s", err))
		return
	}

	r.blackList = append(r.blackList, *requestQuery.Val)
	return

}

func readDeleteNumberRequest(request *restful.Request) (requestQuery DeleteNumberRequest, err error) {

	requestQuery.Number = new(int)

	if len(request.PathParameter(BlackListParameter)) > 0 {
		*requestQuery.Number, err = strconv.Atoi(request.PathParameter(BlackListParameter))
	}

	return
}

func duplicateInBlacklist(arr []int, el int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == el {
			return true
		}
	}

	return false
}

func validateDeleteNumberRequest(requestQuery DeleteNumberRequest) error {
	if requestQuery.Number == nil {
		return fmt.Errorf("Invalid Number: %v", *requestQuery.Number)

	}

	if *requestQuery.Number < 0 {
		return fmt.Errorf("Invalid Number: %v", *requestQuery.Number)

	}

	return nil
}

func (r *Service) validateGetBlacklistRequest(requestQuery GetBlacklistRequest) error {
	if requestQuery.Val == nil {

		return fmt.Errorf("Invalid Number: %v", *requestQuery.Val)
	}

	if *requestQuery.Val < 0 {

		return fmt.Errorf("Invalid Number: %v", requestQuery.Val)
	}

	if duplicateInBlacklist(r.blackList, *requestQuery.Val) == true {
		return fmt.Errorf("%v already exists!", requestQuery.Val)
	}

	return nil
}