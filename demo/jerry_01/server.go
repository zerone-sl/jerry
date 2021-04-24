package main

import (
	"fmt"
	jnet "github.com/zerone/jerry/server"
)

func mainn() {

	go func() {
		fmt.Println("server started")

		jServer := jnet.New("jerry")

		router := jnet.NewRouter()
		jServer.AddRouter(&router)

		jServer.Serve()
	}()

}
