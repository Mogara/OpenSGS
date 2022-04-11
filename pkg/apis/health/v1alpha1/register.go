package v1alpha1

import (
	"net/http"

	restful "github.com/emicklei/go-restful/v3"

	"github.com/Mogara/OpenSGS/pkg/apiserver/helper"
)

const (
	GroupName = "health"
	Version   = "v1alpha1"
)

var GroupVersion = GroupName + "/" + Version

func AddToContainer(c *restful.Container) error {
	ws := helper.NewWebService(GroupVersion)
	ws.Route(ws.GET("/ping").To(handlePing).Returns(http.StatusOK, "ok", nil))
	c.Add(ws)
	return nil
}
