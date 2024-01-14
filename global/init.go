package global

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

var (
	HOST, PORT string
	Listen     net.Listener
	err        error
)

func init() {
	load_err := godotenv.Load("./app/config.env")
	if load_err != nil {
		fmt.Println("error in loading config", load_err.Error())
		log.Fatal("error in loading config", load_err.Error())
	}
	HOST = os.ExpandEnv("$host")
	PORT = os.ExpandEnv("$port")
	fmt.Println("Application port:: ", PORT, " Host:: ", HOST)
	Listen, err = net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("error in listen: ", err)
		return
	}

}
