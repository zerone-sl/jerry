package main

import (
	"fmt"
	jnet "github.com/zerone/jerry/server"
)

func main() {
	fmt.Println("server started")

	jServer := jnet.New("jerry")

	jServer.Serve()
}
