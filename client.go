package skegchat

import (
  "bufio"
  "net"
)

func Listen() {
  conn, err := net.Dial("tcp", "localhost:8080")
  if err != nil {
    print(err)
  }

  result, err := bufio.NewReader(conn).ReadString('\n')
  print(result)
}
