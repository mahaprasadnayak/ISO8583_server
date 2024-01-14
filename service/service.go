package service

import (
	"encoding/hex"
	"fmt"
	"iso8583_server/utils"
	"log"
	"net"
	"os"

	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/network"
)

func HandleRequest(conn net.Conn) {
	fmt.Println("Acquiring New Connection ... ")
	for {

		header := network.NewBinary2BytesHeader()
		header_data, err := header.ReadFrom(conn)
		fmt.Println("Incoming Header Data :: ", header_data)
		if err != nil {
			fmt.Println("error while reading header from conn: ", err)
			fmt.Println("Client closing Connection ... ")
			return
		}

		// Make a buffer to hold msg
		buf := make([]byte, 1024)
		// Read the incoming msg(maindata) into the buffer.

		Connection_data, err := conn.Read(buf)
		fmt.Println("Incomming Connection Data: ", Connection_data)
		if err != nil {
			fmt.Println("error while reading msg", err)
		}
		msg := string(buf[:Connection_data])

		fmt.Println("msg :", msg)

		log.Printf("\n%v", hex.Dump([]byte(msg)))

		required_data := msg

		fmt.Println(required_data)

		var isomessage = iso8583.NewMessage(utils.Spec)
		err = isomessage.Unpack([]byte(required_data))
		if err != nil {
			fmt.Println("error while unpacking msg", err)
		}

		dataMessage := utils.ResponseData{}
		dataMessage.F0, err = isomessage.GetString(0)
		if err != nil {
			fmt.Println("error in getting string 0", err)
		}
		dataMessage.F2, err = isomessage.GetString(2)
		if err != nil {
			fmt.Println("error in getting string 2", err)
		}
		dataMessage.F3, err = isomessage.GetString(3)
		if err != nil {
			fmt.Println("error in getting string 3", err)
		}
		dataMessage.F4, err = isomessage.GetString(4)
		if err != nil {
			fmt.Println("error in getting string 4", err)
		}
		dataMessage.F11, err = isomessage.GetString(11)
		if err != nil {
			fmt.Println("error in getting string 11", err)
		}
		dataMessage.F12, err = isomessage.GetString(12)
		if err != nil {
			fmt.Println("error in getting string 12", err)
		}
		dataMessage.F17, err = isomessage.GetString(17)
		if err != nil {
			fmt.Println("error in getting string 17", err)
		}
		dataMessage.F24, err = isomessage.GetString(24)
		if err != nil {
			fmt.Println("error in getting string 24", err)
		}
		dataMessage.F32, err = isomessage.GetString(32)
		if err != nil {
			fmt.Println("error in getting string 32", err)
		}
		dataMessage.F37, err = isomessage.GetString(37)
		if err != nil {
			fmt.Println("error in getting string 37", err)
		}
		dataMessage.F38, err = isomessage.GetString(38)
		if err != nil {
			fmt.Println("error in getting string 38", err)
		}
		dataMessage.F39, err = isomessage.GetString(39)
		if err != nil {
			fmt.Println("error in getting string 39", err)
		}
		dataMessage.F41, err = isomessage.GetString(41)
		if err != nil {
			fmt.Println("error in getting string 41", err)
		}
		dataMessage.F42, err = isomessage.GetString(42)
		if err != nil {
			fmt.Println("error in getting string 42", err)
		}
		dataMessage.F48, err = isomessage.GetString(48)
		if err != nil {
			fmt.Println("error in getting string 48", err)
		}
		dataMessage.F49, err = isomessage.GetString(49)
		if err != nil {
			fmt.Println("error in getting string 49", err)
		}
		dataMessage.F50, err = isomessage.GetString(50)
		if err != nil {
			fmt.Println("error in getting string 50", err)
		}
		dataMessage.F123, err = isomessage.GetString(123)
		if err != nil {
			fmt.Println("error in getting string 123", err)
		}
		dataMessage.F125, err = isomessage.GetString(125)
		if err != nil {
			fmt.Println("error in getting string 125", err)
		}
		dataMessage.F126, err = isomessage.GetString(126)
		if err != nil {
			fmt.Println("error in getting string 126", err)
		}
		dataMessage.F127, err = isomessage.GetString(127)
		if err != nil {
			fmt.Println("error in getting string 127", err)
		}
		fmt.Printf("ISO Response data :: %+v\n", dataMessage)

		var message = iso8583.NewMessage(utils.Spec)

		if dataMessage.F0 == "" {
			return
		}

		if dataMessage.F0 == "1804" {
			message.MTI("1814")
		} else if dataMessage.F0 == "1420" {
			message.MTI("1430")
		} else if dataMessage.F0 == "1200" {
			message.MTI("1210")
		}

		if dataMessage.F2 != "" {
			message.Field(2, dataMessage.F2)
		}
		if dataMessage.F3 != "" {
			message.Field(3, dataMessage.F3)
		}
		if dataMessage.F4 != "" {
			message.Field(4, dataMessage.F4)
		}
		if dataMessage.F11 != "" {
			message.Field(11, dataMessage.F11)
		}
		if dataMessage.F12 != "" {
			message.Field(12, dataMessage.F12)
		}
		if dataMessage.F14 != "" {
			message.Field(14, dataMessage.F14) //expiry date
		}

		if dataMessage.F17 != "" {
			message.Field(17, dataMessage.F17)
		}
		if dataMessage.F18 != "" {
			message.Field(18, dataMessage.F18)
		}
		if dataMessage.F22 != "" {
			message.Field(22, dataMessage.F22)
		}
		if dataMessage.F24 != "" {
			message.Field(24, dataMessage.F24)
		}
		if dataMessage.F32 != "" {
			message.Field(32, dataMessage.F32)
		}
		if dataMessage.F37 != "" {
			message.Field(37, dataMessage.F37)
		}

		if dataMessage.F0 == "1804" {
			message.Field(39, "800")
		} else {
			message.Field(39, "000")
		}

		if dataMessage.F41 != "" {
			message.Field(41, dataMessage.F41)
		}
		if dataMessage.F42 != "" {
			message.Field(42, dataMessage.F42)
		}
		if dataMessage.F43 != "" {
			message.Field(43, dataMessage.F43)
		}
		if dataMessage.F49 != "" {
			message.Field(49, dataMessage.F49)
		}
		if dataMessage.F50 != "" {
			message.Field(50, dataMessage.F50)
		}
		if dataMessage.F60 != "" {
			message.Field(60, dataMessage.F60)
		}
		if dataMessage.F61 != "" {
			message.Field(61, dataMessage.F61)
		}
		if dataMessage.F102 != "" {
			message.Field(102, dataMessage.F102)
		}
		if dataMessage.F123 != "" {
			message.Field(123, dataMessage.F123)
		}
		if dataMessage.F125 != "" {
			message.Field(125, dataMessage.F125)
		}
		if dataMessage.F126 != "" {
			message.Field(126, dataMessage.F126)
		}

		iso8583.Describe(message, os.Stdout)
		packed, err := message.Pack()
		if err != nil {
			fmt.Println("error while packing msg", err.Error())
		}
		maindata := string(packed)
		log.Printf("\n%v", hex.Dump([]byte(maindata)))
		Write_data, err := conn.Write([]byte(packed))
		if err != nil {
			fmt.Println("error while writing mesaage to connection: ", err)
		}

		fmt.Println("Write_data :", Write_data)

	}
}
