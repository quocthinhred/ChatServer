package main

import (
	"ChatServer/api"
	"gitlab.com/thuocsi.vn-sdk/go-sdk/sdk"
)

var app *sdk.App

func main() {

	app = sdk.NewApp("My Chat Project")

	wss := app.SetupWSServer("chat-wsocket")
	wss.Expose(8000)
	wsRoute := wss.NewRoute("/")
	wsRoute.OnConnected = api.OnWSConnected
	wsRoute.OnMessage = api.OnWSMessage
	wsRoute.OnClose = api.OnWSClose

	app.Launch()
}
