package websocket

type OnWSConnectedHandler = func(conn *Connection)
type OnWSMessageHandler = func(conn *Connection, message string)
type OnWSCloseHandler = func(conn *Connection, err error)

type wsRoute struct {
	OnConnected OnWSConnectedHandler
	OnMessage   OnWSMessageHandler
	OnClose     OnWSCloseHandler

	conMap      map[int]*Connection
	payloadSize int
}

// default construction
func newWSRoute() *wsRoute {
	return &wsRoute{
		conMap:      map[int]*Connection{},
		payloadSize: 512, // default 512
	}
}

// transferring keep chatting with client
func (wsr *wsRoute) transferring(con *Connection) {

	for {
		if con.rootCon.MaxPayloadBytes != wsr.payloadSize {
			con.rootCon.MaxPayloadBytes = wsr.payloadSize
		}
		payload, err := con.Read()
		if err != nil {
			con.Deactive()
			delete(wsr.conMap, con.Id)
			if wsr.OnClose != nil {
				wsr.OnClose(con, err)
			}
			return
		} else {
			if wsr.OnMessage != nil {
				wsr.OnMessage(con, payload)
			}
		}
	}
}

func (wsr *wsRoute) addCon(con *Connection) {
	wsr.conMap[con.Id] = con
}

func (wsr *wsRoute) GetConnectionMap() map[int]*Connection {
	return wsr.conMap
}

func (wsr *wsRoute) GetConnection(id int) *Connection {
	return wsr.conMap[id]
}

func (wsr *wsRoute) SetPayloadSize(size int) {
	wsr.payloadSize = size

}
