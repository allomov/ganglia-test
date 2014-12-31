package main

// import "sync"

type InfusionPump struct {
	// mutex sync.RWMutex
	works bool

}

// type connection struct {
// 	// The websocket connection.
// 	ws *websocket.Conn

// 	// Buffered channel of outbound messages.
// 	send chan []byte
// }

// type hub struct {
// 	connections map[*connection]bool
// 	broadcast chan []byte
// 	register chan *connection
// 	unregister chan *connection
// }

// var h = hub {
// 	broadcast:   make(chan []byte),
// 	register:    make(chan *connection),
// 	unregister:  make(chan *connection),
// 	connections: make(map[*connection]bool),
// }

func NewInfusionPump() (infusionPump InfusionPump) {
	infusionPump = InfusionPump{true}
	return
}


func (ivp *InfusionPump) Set(value int) {
	// ivp.mutex.Lock()

	// ivp.mutex.Unlock()
}

