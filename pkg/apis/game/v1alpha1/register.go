package v1alpha1

import (
	restful "github.com/emicklei/go-restful/v3"

	"github.com/Mogara/OpenSGS/pkg/apiserver/helper"
)

const (
	GroupName = "game"
	Version   = "v1alpha1"
)

var GroupVersion = GroupName + "/" + Version

func AddToContainer(c *restful.Container) error {
	ws := helper.NewWebService(GroupVersion)

	gamer := newGameHandler()

	ws.Route(ws.GET("/game/{room}").
		To(gamer.handleSession).
		Param(ws.PathParameter("room", "room id")).
		Doc("handle game room session").
		Writes(nil))

	c.Add(ws)
	return nil
}
