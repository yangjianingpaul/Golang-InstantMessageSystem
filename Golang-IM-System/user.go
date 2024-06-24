package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// Create a user API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}
	// Start a goroutine that listens for messages on the current user channel
	go user.ListenMessage()
	return user
}

// User on-line function
func (this *User) Online() {
	// The user goes online and is added to onlineMap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// Broadcast the message that the current user goes online
	this.server.BroadCast(this, "Go online")
}

// User logout function
func (this *User) Offline() {
	//The user is logged out. The user is deleted from onlineMap
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// The message that the current user is offline is broadcast
	this.server.BroadCast(this, "Be offline")
}

// Sends a message to the current User client
func (this *User) sendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// Users process messages
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// Example Query the current online users
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "on line...\n"
			this.sendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// message format：Message format: rename|new name
		newName := strings.Split(msg, "|")[1]
		// Check whether the name exists
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.sendMsg("The current user name is used\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.sendMsg("You have updated your username：" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// Message format: to| recipient name | Message content
		// 1) Get the user name of the other party
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.sendMsg("The message format is incorrect. Please use \"to| recipient name | Hello")
			return
		}
		// 2）Get the User object of the other party based on the user name
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.sendMsg("The user name does not exist\n")
			return
		}
		// 3）Get the message content and send the message content through the User object of the other party
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.sendMsg("No message, please resend\n")
			return
		}
		remoteUser.sendMsg(this.Name + "Say to you:" + content)

	} else {
		this.server.BroadCast(this, msg)
	}

}

// A method of listening to the current User channel and sending messages directly to the peer client once there is a message
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
