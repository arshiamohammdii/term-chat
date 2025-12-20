package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
	"term-chat.com/room"
)

var availableRooms = []*room.Room{
	room.NewRoom("hacking"),
	room.NewRoom("studying"),
}

func handleCommand(line string, user *room.User) {
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
				room.Sync.Lock()
				user.Room = room
				availableRooms[i].Users[user] = true
				room.Sync.Unlock()
				user.Term.Write(fmt.Appendf(nil, "You joined room %s", room.Name))
				return
			}
		}
		user.Term.Write([]byte("Room Not Found"))
	case "rooms":
		for _, room := range availableRooms {
			user.Term.Write(fmt.Appendf(nil, "%s\n", room.Name))
		}
	}

}

func main() {
	ssh.Handle(func(s ssh.Session) {
		terminal := term.NewTerminal(s, "\n> ")
		user := &room.User{SSH: s, TimeJoined: time.Now(), Term: terminal}

		terminal.Write([]byte("Welcome to the SSH chat server\n"))
		for {
			line, err := terminal.ReadLine()
			if strings.HasPrefix(line, "/") {
				handleCommand(line, user)
			} else {
				if user.Room == nil {
					user.Term.Write([]byte("Join a room first (/join <room>)\n"))
					continue
				}
				user.Room.Send(room.Message{
					From: user,
					Body: line,
				})
			}
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
		}
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil, ssh.HostKeyFile("/Users/arshiamohammadi/.ssh/host_key")))
}
