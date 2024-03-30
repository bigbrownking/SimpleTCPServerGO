# TCP Chat Application

This is a simple TCP chat application written in Go. It consists of a server and a client that can send and receive messages over a TCP connection.

## Components

### 1. Client

The client is responsible for establishing a TCP connection to the server, sending messages, and receiving responses. It also keeps a history of sent messages.

### 2. Server

The server listens for incoming TCP connections and handles each connection in a separate goroutine. When a message is received from a client, the server logs the message and sends a confirmation response back to the client.

### 3. Constants

This package contains constants used by both the client and the server, such as the host, port, and connection type.

## Usage

To run the server:

```bash
go run server.go
