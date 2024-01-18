package router

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/novalagung/gubrak/v2"
	"log"
	"net/http"
	"strings"
)

type M map[string]interface{}

var connections = make([]*WebSocketConnection, 0)

type SocketPayload struct {
	Message string
}

func StartWebsocket(e *echo.Echo) {
	e.GET("/ws", func(c echo.Context) error {
		currentGorillaConn, err := websocket.Upgrade(c.Response(), c.Request(), c.Response().Header(), 1024, 1024)
		if err != nil {
			http.Error(c.Response(), "Could not open websocket connection", http.StatusBadRequest)
		}

		currentConn := WebSocketConnection{Conn: currentGorillaConn}
		connections = append(connections, &currentConn)

		go handleIO(&currentConn, connections)
		return nil
	})
}

type SocketResponse struct {
	From    string
	Type    string
	Message string
}

type WebSocketConnection struct {
	*websocket.Conn
}

func handleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	broadcastMessage(currentConn, "content_update")

	for {
		payload := SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, "content_update")
				ejectConnection(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		broadcastMessage(currentConn, "content_update")
	}
}

func ejectConnection(currentConn *WebSocketConnection) {
	filtered := gubrak.From(connections).Reject(func(each *WebSocketConnection) bool {
		return each == currentConn
	}).Result()
	connections = filtered.([]*WebSocketConnection)
}

func broadcastMessage(currentConn *WebSocketConnection, kind string) {
	log.Println("BROADCAST CALL", len(connections), kind)

	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		err := eachConn.WriteJSON(SocketResponse{
			Type: kind,
		})

		if err != nil {
			log.Println("ERROR ON BROADCAST", err.Error())
			return
		}
	}
}
