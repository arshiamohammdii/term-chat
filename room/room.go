package room

import (
	"fmt"
	"sync"
)

type Room struct {
	Name    string
	Users   map[*User]bool
	History []Message

	Join      chan *User
	Broadcast chan Message
	Sync      sync.Mutex
}

func (r *Room) run() {
	for {
		select {
		case msg := <-r.Broadcast:
			for user := range r.Users {
				// skip the sender
				if msg.From != nil && user == msg.From {
					continue
				}
				fmt.Fprintf(user.Term, "%s: %s\n", msg.From.Name, msg.Body)
			}
		case joined := <-r.Join:
			r.Users[joined] = true
			for user := range r.Users {
				fmt.Fprintf(user.Term, "%s joined the room\n", joined.Name)
			}

		}
	}
}

func NewRoom(name string) *Room {
	r := &Room{
		Name:  name,
		Users: make(map[*User]bool),
		Join:  make(chan *User),
		// Leave:     make(chan *User),
		Broadcast: make(chan Message),
	}

	go r.run()
	return r
}

func (r *Room) Send(msg Message) {
	r.Broadcast <- msg
}
