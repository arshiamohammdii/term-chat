package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type User struct {
	Name       string
	ssh        *ssh.Session
	TimeJoined time.Time
	Term       *term.Terminal
}

type Room struct {
	Name  string
	Users map[*User]bool
}

var availableRooms = []*Room{
	{Name: "hacking", Users: map[*User]bool{}},
	{Name: "studying", Users: map[*User]bool{}},
}

func handleCommand(line string, user *User) {
	parts := strings.Fields(line)
	cmd := strings.TrimPrefix(parts[0], "/")

	switch cmd {
	case "join":
		if len(parts) < 2 {
			user.Term.Write([]byte("Usage /join <room>"))
			return
		}

		targetRoom := parts[1]
		for i, room := range availableRooms {
			if room.Name == targetRoom {
				availableRooms[i].Users[user] = true
				user.Term.Write(fmt.Appendf(nil, "You joined room %s", room.Name))
			}
		}
	case "rooms":
		for _, room := range availableRooms {
			user.Term.Write([]byte(room.Name))
		}
	}
}

func main() {
	ssh.Handle(func(s ssh.Session) {
		terminal := term.NewTerminal(s, "\n> ")
		user := &User{ssh: &s, TimeJoined: time.Now(), Term: terminal}

		terminal.Write([]byte("Welcome to the SSH chat server\n"))
		for {
			line, err := terminal.ReadLine()
			// terminal.Write([]byte(s.User()))
			if strings.HasPrefix(line, "/") {
				handleCommand(line, user)
			}
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			// terminal.Write([]byte(line))
		}
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil, ssh.HostKeyFile("/Users/arshiamohammadi/.ssh/host_key")))
}
