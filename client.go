package main

import (
	"Ex1_Week1/constants"
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

func main() {
	log.AddHandler(console.New(true), log.AllLevels...)

	tcpServer, err := net.ResolveTCPAddr(constants.TYPE, constants.HOST+":"+constants.PORT)
	if err != nil {
		log.WithError(err).Error("ResolveTCPAddr failed")
		os.Exit(1)
	}

	conn, err := net.DialTCP(constants.TYPE, nil, tcpServer)
	if err != nil {
		log.WithError(err).Error("Dial failed")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')

	_, err = conn.Write([]byte(username))
	if err != nil {
		log.WithError(err).Error("Write data failed")
		os.Exit(1)
	}

	for {
		log.Info("Enter message: ")
		text, _ := reader.ReadString('\n')

		_, err = conn.Write([]byte(text))
		if err != nil {
			log.WithError(err).Error("Write data failed")
			os.Exit(1)
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.WithError(err).Error("Read response failed")
			os.Exit(1)
		}
		fmt.Print(response)

		time.Sleep(1 * time.Second)
	}
}
