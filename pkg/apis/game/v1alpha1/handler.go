package v1alpha1

import (
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/Mogara/OpenSGS/pkg/models/game"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type gameHandler struct {
	gamer game.Interface
}

func newGameHandler() *gameHandler {
	return &gameHandler{
		gamer: game.NewGame(),
	}
}

func (g *gameHandler) handleSession(request *restful.Request, response *restful.Response) {
	roomId := request.PathParameter("room")
	conn, err := upgrader.Upgrade(response.ResponseWriter, request.Request, nil)
	if err != nil {
		log.WithField("url", request.Request.URL).WithError(err).Warn("upgrade to websocket failed")
		return
	}
	g.gamer.HandleSession(roomId, conn)
}
