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

type getBlacklistElements struct {
	Response []int `json:"response"`
}

/*
######################################################################################################
######################################  ADD TO BLACKLIST  ############################################
######################################################################################################
*/

func readAddBlacklistRequest(request *restful.Request) (requestQuery GetBlacklistRequest, err error) {

	requestQuery.Val = new(int)

	if len(request.PathParameter(BlackListParameter)) > 0 {
		*requestQuery.Val, err = strconv.Atoi(request.PathParameter(BlackListParameter))
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

/*
######################################################################################################
######################################  DELETE FROM BLACKLIST  #######################################
######################################################################################################
*/

func readDeleteNumberRequest(request *restful.Request) (requestQuery DeleteNumberRequest, err error) {

	requestQuery.Number = new(int)

	if len(request.PathParameter(BlackListParameter)) > 0 {
		*requestQuery.Number, err = strconv.Atoi(request.PathParameter(BlackListParameter))
	}

	return
}

func (r *Service) DeleteNumberRequest(request *restful.Request, response *restful.Response) {

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
	r.blackList = removeFromArray(*requestQuery.Number, r.blackList)

}

func removeFromArray(val int, array []int) []int {
	for i, w := range array {
		if val == w {
			//r.blackList = append(r.blackList[:i], r.blackList[i+1:])
			//responseJson.Response = append(r.blackList[:i], r.blackList[i+1:]...)
			array[i] = array[len(array)-1]
			return array[:len(array)-1] // asa se sterge o valoare din array

		}
	}
	return array
}

/*
######################################################################################################
########################################  PRINT BLACKLIST  ###########################################
######################################################################################################
*/

func (r *Service) printBlacklist(request *restful.Request, response *restful.Response) {
	var responseJson getBlacklistElements

	responseJson.Response = r.blackList[:]

	response.WriteAsJson(responseJson)

	return
}

func validateDeleteNumberRequest(requestQuery DeleteNumberRequest) error {
	if requestQuery.Number == nil {
		return fmt.Errorf("invalid number: %v", *requestQuery.Number)

	}

	if *requestQuery.Number < 0 {
		return fmt.Errorf("invalid number: %v", *requestQuery.Number)

	}

	return nil
}

func (r *Service) validateGetBlacklistRequest(requestQuery GetBlacklistRequest) error {
	if requestQuery.Val == nil {

		return fmt.Errorf("invalid number: %v", *requestQuery.Val)
	}

	if *requestQuery.Val < 0 {

		return fmt.Errorf("invalid number: %v", requestQuery.Val)
	}

	if duplicateInBlacklist(r.blackList, *requestQuery.Val) {
		return fmt.Errorf("%v already exists!", requestQuery.Val)
	}

	return nil
}

/*
######################################################################################################
##################################  VERIFY DUPLICATES IN BLACKLIST ###################################
######################################################################################################
*/

func duplicateInBlacklist(arr []int, el int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == el {
			return true
		}
	}

	return false
}
