package rest

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

type Service struct {
	container *restful.Container
}

func NewService() (*Service, error) {
	r := &Service{
		container: restful.NewContainer(),
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
		Route(ws.GET("/health_test1").
			Operation("health check").
			To(func(request *restful.Request, response *restful.Response) {
				_ = response.WriteErrorString(http.StatusOK, "internship-api is up and running!")
			}))

	ws.Route(ws.GET("/fibonacci").
		Param(restful.QueryParameter("index", "pozitie").DataType("int").Required(true)).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetFibonacciResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.GetFibonacciRequest).
		Writes(map[string]string{}))

	return ws
}
