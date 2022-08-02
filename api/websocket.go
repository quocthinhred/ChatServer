package api

import (
	"encoding/json"
	"fmt"
	"gitlab.com/thuocsi.vn-sdk/go-sdk/sdk/websocket"
)

type messageInterface struct {
	Name    string `json:"Name,omitempty" bson:"Name, omitempty"`
	Message string `json:"Message,omitempty" bson:"Message, omitempty"`
}

var conns []*websocket.Connection

func OnWSConnected(conn *websocket.Connection) {
	fmt.Println("Connected!")
	conns = append(conns, conn)
}

func OnWSMessage(conn *websocket.Connection, message string) {
	var msg messageInterface
	fmt.Println(message)
	err := json.Unmarshal([]byte(message), &msg)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(msg)
	for _, i := range conns {
		if i != conn {
			temp, _ := json.Marshal(msg)
			i.Send(string(temp))
		}
	}
}

func OnWSClose(conn *websocket.Connection, err error) {

}
