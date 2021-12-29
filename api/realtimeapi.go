package api

import (
	"sync"

	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var connMap sync.Map

func (obj *API) setup(c *fiber.Ctx) error {
	// IsWebSocketUpgrade returns true if the client
	// requested upgrade to the WebSocket protocol.
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (obj *API) listen(c *websocket.Conn) {
	// c.Locals is added to the *websocket.Conn
	println(c.Locals("allowed")) // true
	println(c.Params("id"))      // 123

	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			println("read:", err)
			break
		}
		println("recv:", msg)

		if err = c.WriteMessage(mt, msg); err != nil {
			println("write:", err)
			break
		}
	}
}

func (obj *API) SendMessage(groupID uint, message model.MessageOutDTO) {
	// Send meesage to all websockets for given group
}
