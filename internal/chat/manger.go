package chat

/*
연결된 클라이언트들을 추적
클라이언트들의 등록
불필요한 클라이언트 객체 관리 (Destoryed or Removed)
*/
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}
