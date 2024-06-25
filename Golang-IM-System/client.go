package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //Current mode of client
}

func NewClient(serverIp string, serverPort int) *Client {
	// Create a client object
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	// Connect server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn
	// return object
	return client
}

// The message returned by the server is processed and displayed directly to the standard output
func (client *Client) DealResponse() {
	// Once client.conn has data, it copies it directly to stdout standard output, blocking the listening permanently
	io.Copy(os.Stdout, client.conn)

	// equivalent code
	/*for {
		buf := make([]byte, 4096)
		client.conn.Read(buf)
		fmt.Println(buf)
	}*/
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.Public chat mode")
	fmt.Println("2.Private chat mode")
	fmt.Println("3.Modifying the User Name")
	fmt.Println("0.Exit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>Please enter a number within the legal range<<<<<")
		return false
	}
}

// Query Online Users
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}
}

// Private chat mode
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println(">>>>>Please enter the chat object [user name], 'exit' exit:")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>Please enter the chat object [user name], 'exit' exit:")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// The message is sent if it is not empty
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>>>Please enter the message content, 'exit' exit:")
			fmt.Scanln(&chatMsg)
		}

		client.SelectUsers()
		fmt.Println(">>>>>Please enter the chat object [user name], 'exit' exit:")
		fmt.Scanln(&remoteName)
	}
}

func (client *Client) PublicChat() {
	// Prompts the user for a message
	var chatMsg string
	fmt.Println(">>>>>Please enter the chat content, 'exit' exit.")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// Send to server

		// The message is sent if it is not empty
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>>Please enter the chat content, 'exit' exit.")
		fmt.Scanln(&chatMsg)
	}
	// Send to server
}

func (client *Client) updateName() bool {
	fmt.Println(">>>>>Please enter your username:")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))

	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}

		// Different businesses are processed according to different patterns
		switch client.flag {
		case 1:
			// Public chat mode
			client.PublicChat()
			break
		case 2:
			// Private chat mode
			client.PrivateChat()
			break
		case 3:
			// Update user name
			client.updateName()
			break
		}
	}
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1(default) -port 8888(default)

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "Set the server ip address（default is:127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "Setting the server port（default is:8888）")
}

func main() {
	// Command line parsing
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>Failed to connect to server...")
		return
	}

	// Open a separate goroutine to handle server receipt messages
	go client.DealResponse()

	fmt.Println(">>>>>Successfully connecting to the server...")
	// Start services on the client
	// select {}
	client.Run()
}
