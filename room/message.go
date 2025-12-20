package room

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"

	"time"
)

type User struct {
	Name       string
	SSH        ssh.Session
	TimeJoined time.Time
	Term       *term.Terminal
	Room       *Room
}

type Message struct {
	From *User
	Body string
}
