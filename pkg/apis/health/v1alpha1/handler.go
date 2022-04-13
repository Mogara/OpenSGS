package v1alpha1

import restful "github.com/emicklei/go-restful/v3"

type status struct {
	Ok bool `json:"ok"`
}

func handlePing(request *restful.Request, response *restful.Response) {
	_ = response.WriteAsJson(status{Ok: true})
}
