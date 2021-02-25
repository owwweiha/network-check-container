package main

import (
	"fmt"
	"time"
	"net"
        "os"
)

func main() {
  server := os.Args[1]
  port := os.Args[2]
  raw_connect(server, port)
}

func raw_connect(host string, port string) {
  timeout := time.Second
  conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
  if err != nil {
    fmt.Println("Connecting failed:", err)
  }
  if conn != nil {
    defer conn.Close()
    fmt.Println("Connection succesful:", net.JoinHostPort(host, port))
  }
}
