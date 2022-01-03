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

		conns, ok := connMap.Load(gid)
		if ok {
			conns := conns.([]*websocket.Conn)
			conns = append(conns, c)
			connMap.Store(gid, conns)
		} else {
			conns = []*websocket.Conn{c}
			connMap.Store(gid, conns)
		}
	})
}

func (obj *API) SendMessage(groupID uint, message model.MessageOutDTO) {
	// Send meesage to all websockets for given group
	conns, ok := connMap.Load(groupID)
	if ok {
		for _, conn := range conns.([]*websocket.Conn) {
			conn.WriteJSON(message)
		}
	}
}
