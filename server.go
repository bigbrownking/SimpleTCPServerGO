package main

import (
	"Ex1_Week1/constants"
	"bufio"
	"net"
	"os"
	"sync"

	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

var (
	clients = make(map[net.Conn]bool)
	mu      sync.Mutex
)

func main() {
	log.AddHandler(console.New(true), log.AllLevels...)

	listen, err := net.Listen(constants.TYPE, constants.HOST+":"+constants.PORT)
	if err != nil {
		log.WithError(err).Error("error starting server")
		os.Exit(1)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.WithError(err).Error("error accepting connection")
			os.Exit(1)
		}
		mu.Lock()
		clients[conn] = true
		mu.Unlock()
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.WithError(err).Error("error reading from client")
		return
	}

	log.Infof("New client connected: %s", username)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Infof("%s disconnected.", username)
			break
		}

		log.Infof("Received message from %s: %s", username, message)

		_, err = conn.Write([]byte("Thank you for sending message\n"))
		if err != nil {
			log.WithError(err).Error("Error sending confirmation message to client")
			break
		}
	}
}
