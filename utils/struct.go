package utils

import (
	"net"
)

var (
	Server_conn net.Conn
)

type ResponseCode struct {
	Resposecode  string `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
}

// request for echo_check
type Req_echo struct {
	Stan string `json:"stan"`
}

// response from echo_check
type Res_echo struct {
	F0  string
	F39 string
}

type RequestCode struct {
	Id           string `json:"id"`
	AadhaarRefId string `json:"aadhaarRefId"`
	Amount       string `json:"amount"`
	Stan         string `json:"stan"`
	Rrn          string `json:"rrn"`
	TerminalId   string `json:"terminalId"`
	MerchantId   string `json:"merchantId"`
	Location     string `json:"location"`
	UserId       string `json:"userId"`
	Type         string `json:"type"`
	Aadhartoken  string `json:"aadharToken"`
	ReqOrgId     string `json:"reqOrgId"`
	Dd           string `json:"dd"`
}

type ResponseData struct {
	F0   string
	F2   string
	F3   string
	F4   string
	F11  string
	F12  string
	F14  string
	F17  string
	F18  string
	F22  string
	F24  string
	F32  string
	F37  string
	F38  string
	F39  string
	F41  string
	F42  string
	F43  string
	F48  string
	F49  string
	F50  string
	F60  string
	F61  string
	F63  string
	F102 string
	F123 string
	F125 string
	F126 string
	F127 string
}

type CbsService struct {
	Ltd                string
	Id                 string
	Mti                string
	AadhaarRefId       string
	Amount             string
	Stan               string
	Rrn                string
	TerminalId         string
	MerchantId         string
	Location           string
	UserId             string
	Type               string
	Aadhartoken        string
	Reserved_national  string //60
	Reserved_private   string //61
	Mac                string
	ReqOrgId           string
	ReqIn_time         string
	ResOut_time        string
	Status             string
	ResponseCode       string
	Controller_id      string
	ProcessingCode     string
	Reserved_Iso       string //56 reversal
	Is_Reversal        bool
	Status_Description string
}
