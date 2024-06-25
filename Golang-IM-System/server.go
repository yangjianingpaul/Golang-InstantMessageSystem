package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// List of online users
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	// A channel for message broadcasting
	Message chan string
}

// An interface of create the server
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// A goroutine that listens to the Message broadcast message channel and sends messages to all online users as soon as they arrive
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//msg is sent to all online users
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	// Currently connected services
	// fmt.Println("Connection established successfully")

	user := NewUser(conn, this)

	user.Online()

	// A channel that listens for whether the user is active
	isLive := make(chan bool)

	// Accepts the message sent by the client
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			// Extract the user's message (remove '\n')
			msg := string(buf[:n-1])
			// The user processes messages against msg
			user.DoMessage(msg)
			// User Any message indicates that the current user is active
			isLive <- true
		}
	}()

	// Current handler blocking
	for {
		select {
		case <-isLive:
			// The current user is active and the timer should be reset
			// Without doing anything, in order to activate select, update the timer below
		case <-time.After(time.Second * 300):
			// Have timed out
			// Forcibly disable the current User
			user.sendMsg("You are forced to quit!")

			// Destroy user resources
			close(user.C)
			// close connection
			conn.Close()
			// Exit the current handler
			// runtime.Goexit()
			return
		}
	}
}

// Interface for starting the server
func (this *Server) Start() {
	// socket lister
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}

	defer listener.Close()

	// Start a goroutine that listens for messages
	go this.ListenMessager()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}

	// close listen socket
}
