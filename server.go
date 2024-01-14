package main

import (
	"fmt"
	"iso8583_server/global"
	"iso8583_server/service"
)

func main() {
	// close listener
	defer global.Listen.Close()
	for {
		fmt.Println("listening")
		ser_conn, err := global.Listen.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go service.HandleRequest(ser_conn)
	}
}
