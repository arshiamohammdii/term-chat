package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/arshiamohammdii/term-chat/chat"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type Server struct {
	Rooms   []*chat.Room
	HostKey string
	Addr    string
}

func NewServer(addr string, hostkey string) *Server {
	return &Server{
		HostKey: hostkey,
		Addr:    addr,
		Rooms: []*chat.Room{
			chat.NewRoom("hacking"),
			chat.NewRoom("studying"),
		},
	}
}

func (srv *Server) ListenSSH() error {
	ssh.Handle(func(sess ssh.Session) {
		// theme := chat.InitTheme()

		terminal := term.NewTerminal(sess, fmt.Sprintf("%s> ", sess.User()))
		user := &chat.User{
			Name:       sess.User(),
			SSH:        sess,
			TimeJoined: time.Now(),
			Term:       terminal,
		}

		//clear the screen
		fmt.Fprint(terminal, "\x1b[2J\x1b[H")

		terminal.Write([]byte("Welcome to the SSH chat server\n"))
		for {
			line, err := terminal.ReadLine()
			if err != nil {
				return
			}

			if strings.HasPrefix(line, "/") {
				srv.handleCommand(line, user)
				continue
			}

			if user.Room == nil {
				user.Term.Write([]byte("Join a room first (/join <room>)\n"))
				continue
			}

			user.Room.Send(chat.Message{
				From: user,
				Body: line,
			})
		}
	})

	return ssh.ListenAndServe(srv.Addr, nil, ssh.HostKeyFile(srv.HostKey))
}

func (s *Server) handleCommand(line string, user *chat.User) {
	parts := strings.Fields(line)
	cmd := strings.TrimPrefix(parts[0], "/")

	switch cmd {
	case "join":
		if len(parts) < 2 {
			user.Term.Write([]byte("Usage /join <room>"))
			return
		}

		targetRoom := parts[1]
		for _, r := range s.Rooms {
			if r.Name == targetRoom {
				user.Room = r
				r.Join <- user
				return
			}
		}
		user.Term.Write([]byte("Room Not Found"))
	case "rooms":
		for _, room := range s.Rooms {
			user.Term.Write(fmt.Appendf(nil, "%s\n", room.Name))
		}
	}

}
