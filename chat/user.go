package chat

import (
	"time"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type User struct {
	Name       string
	SSH        ssh.Session
	TimeJoined time.Time
	Term       *term.Terminal
	Room       *Room
}
