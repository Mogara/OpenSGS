package helper

import (
	restful "github.com/emicklei/go-restful/v3"
)

const (
	ApiRootPath = "/api"
)

func NewWebService(groupVersion string) *restful.WebService {
	webService := restful.WebService{}
	webService.Path(ApiRootPath + "/" + groupVersion).Produces(restful.MIME_JSON).Consumes(restful.MIME_JSON)
	return &webService
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
