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

	//fmt.Println(response.WriteAsJson(responseJson)) <-----1
	response.WriteAsJson(responseJson)
}

func readDeleteNumberRequest(request *restful.Request) (requestQuery DeleteNumberRequest, err error) {

	requestQuery.Number = new(int)

	if len(request.PathParameter(BlackListParameter)) > 0 {
		*requestQuery.Number, err = strconv.Atoi(request.PathParameter(BlackListParameter))
	}

	return
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

func printBlackList(list []int) {

	fmt.Println(list[:])

}
