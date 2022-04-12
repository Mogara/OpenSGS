package apiserver

import (
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	log "github.com/sirupsen/logrus"

	healthv1alpha1 "github.com/Mogara/OpenSGS/pkg/apis/health/v1alpha1"
	"github.com/Mogara/OpenSGS/pkg/apiserver/helper"
)

type APIServer struct {
	Server    *http.Server
	container *restful.Container
}

func NewAPIServer(host string, port int) *APIServer {
	s := &APIServer{}
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}
	s.Server = server
	return s
}

func (s *APIServer) PrepareRun() error {
	container := restful.NewContainer()
	container.Router(restful.CurlyRouter{})
	container.RecoverHandler(recoverHandler)
	container.DoNotRecover(false)
	s.container = container

	s.installAPIs()

	for _, ws := range s.container.RegisteredWebServices() {
		for _, route := range ws.Routes() {
			log.Infof("%s %s --> %s", route.Method, route.Path, nameOfFunction(route.Function))
		}
	}

	s.Server.Handler = s.container
	return nil
}

func (s *APIServer) Run() error {
	return s.Server.ListenAndServe()
}

func (s *APIServer) installAPIs() {
	helper.Must(healthv1alpha1.AddToContainer(s.container))
}
