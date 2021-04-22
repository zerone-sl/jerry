package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	mainn()
	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:6069")
	if err != nil {
		fmt.Println("conn to tcp server err: ", err)
		return
	}
	fmt.Println("client started")
	go func() {
		idx := 0
		for {
			msg := fmt.Sprintf("[client] hello server, i am client %d", idx)
			fmt.Println(msg)
			conn.Write([]byte(msg))

			idx++
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			buf := make([]byte, 512)
			cnt, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read from server err: ", err)
				break
			}
			fmt.Println("[server] ", string(buf[:cnt]))
		}

	}()

	time.Sleep(60 * time.Second)
}
