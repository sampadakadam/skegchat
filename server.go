package skegchat

import (
  "fmt"
  "net"
)

func handleConnection(conn net.Conn) {
  fmt.Fprintf(conn, "Connected\r\n")
  conn.Close()
}

func Serve() {
  listener, err := net.Listen("tcp", ":8080")
  if err != nil {
    print(err)
  }
  for {
    conn, err := listener.Accept()
    if err != nil {
      print(err)
    }
    go handleConnection(conn)
  }
}
