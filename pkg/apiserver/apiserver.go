package apiserver

import (
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	gamev1alpha1 "github.com/Mogara/OpenSGS/pkg/apis/game/v1alpha1"
	healthv1alpha1 "github.com/Mogara/OpenSGS/pkg/apis/health/v1alpha1"
	"github.com/Mogara/OpenSGS/pkg/apiserver/helper"
)

type APIServer struct {
	Server         *http.Server
	container      *restful.Container
	allowedOrigins []string
	debug          bool
}

func NewAPIServer(host string, port int, allowedOrigins []string, debug bool) *APIServer {
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}
	return &APIServer{
		Server:         server,
		allowedOrigins: allowedOrigins,
		debug:          debug,
	}
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
	c := cors.New(cors.Options{
		Debug:            s.debug,
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowOriginFunc:  allowOriginFunc(s.allowedOrigins),
	})

	s.Server.Handler = c.Handler(s.container)
	return nil
}

func (s *APIServer) Run() error {
	return s.Server.ListenAndServe()
}

func (s *APIServer) installAPIs() {
	helper.Must(healthv1alpha1.AddToContainer(s.container))
	helper.Must(gamev1alpha1.AddToContainer(s.container))
}
