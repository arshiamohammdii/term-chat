package room

import "sync"

type Room struct {
	Name    string
	Users   map[*User]bool
	History []Message

	Broadcast chan Message
	Sync      sync.Mutex
}

func (r *Room) run() {
	for {
		select {
		case msg := <-r.Broadcast:
			for user := range r.Users {
				user.Term.Write([]byte(msg.Body))
			}
		}
	}
}

func NewRoom(name string) *Room {
	r := &Room{
		Name:  name,
		Users: make(map[*User]bool),
		// Join:      make(chan *User),
		// Leave:     make(chan *User),
		Broadcast: make(chan Message),
	}

	go r.run()
	return r
}

func (r *Room) Send(msg Message) {
	r.Broadcast <- msg
}
