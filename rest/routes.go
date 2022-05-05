package rest

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

const (
	BlackListParameter = "BlackListValue"
)

type Service struct {
	container *restful.Container
	blackList []int
}

func NewService() (*Service, error) {
	r := &Service{
		container: restful.NewContainer(),
		blackList: make([]int, 0),
	}

	r.container.Add(r.buildRoutes())

	return r, nil
}

func (r *Service) Container() *restful.Container {
	return r.container
}

func (r *Service) buildRoutes() *restful.WebService {
	ws := new(restful.WebService)

	ws.Path("/internship-api").
		Route(ws.GET("/health").
			Operation("health check").
			To(func(request *restful.Request, response *restful.Response) {
				_ = response.WriteErrorString(http.StatusOK, "internship-api is up and running!")
			}))

	ws.Route(ws.GET("/fibonacci").
		Param(restful.QueryParameter("index", "pozitie").DataType("int").Required(true)).
		Param(restful.QueryParameter("size", "pozitie").DataType("int").Required(false)).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetFibonacciResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.GetFibonacciRequest).
		Writes(map[string]string{}))

	ws.Route(ws.PUT("/blackList/{"+BlackListParameter+"}").
		Param(restful.QueryParameter("val", "numar").DataType("int").Required(true)).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetBlacklistResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.AddBlacklistElement).
		Writes(map[string]string{}))

	ws.Route(ws.DELETE("/blackList/{"+BlackListParameter+"}").
		Param(restful.QueryParameter("number", "valoare").DataType("int").Required(true)).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetFibonacciResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.DeleteNumberRequest).
		Writes(map[string]string{}))

	ws.Route(ws.GET("/blackList").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetFibonacciResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.printBlacklist).
		Writes(map[string]string{}))

	return ws
}
