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
		gid := obj.getGroupIDFromToken(c)
		c.Locals(id, gid)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (obj *API) websocket() func(*fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		id := c.Locals(id).(float64)
		gid := uint(id)

		newConn := make(chan model.MessageOutDTO)
		conns, ok := connMap.Load(gid)
		if ok {
			conns := conns.([]chan model.MessageOutDTO)
			conns = append(conns, newConn)
			connMap.Store(gid, conns)
		} else {
			conns = []chan model.MessageOutDTO{newConn}
			connMap.Store(gid, conns)
		}

		for {
			message := <-newConn
			c.WriteJSON(message)
		}
	})
}

func (obj *API) SendMessage(groupID uint, message model.MessageOutDTO) {
	// Send meesage to all websockets for given group
	conns, ok := connMap.Load(groupID)
	if ok {
		for _, conn := range conns.([]chan model.MessageOutDTO) {
			conn <- message
		}
	}
}
