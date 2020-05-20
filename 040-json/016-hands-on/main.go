//hands-on
// see instructions.md

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type httpStatusCodes []struct {
	Code        int    `json:"Code"`
	Description string `json:"Description"`
}

func main() {
	var data httpStatusCodes
	recvd := `[{"Code":200,"Description":"StatusOK"},
	{"Code":301,"Description":"StatusMovedPermanently"},
	{"Code":302,"Description":"StatusFound"},
	{"Code":303,"Description":"StatusSeeOther"},
	{"Code":307,"Description":"StatusTemporaryRedirect"},
	{"Code":400,"Description":"StatusBadRequest"},
	{"Code":401,"Description":"StatusUnauthorized"},
	{"Code":402,"Description":"StatusPaymentRequired"},
	{"Code":403,"Description":"StatusForbidden"},
	{"Code":404,"Description":"StatusNotFound"},
	{"Code":405,"Description":"StatusMethodNotAllowed"},
	{"Code":418,"Description":"StatusTeapot"},
	{"Code":500,"Description":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(recvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	for i, v := range data {
		fmt.Println(i, v.Code, v.Description)
	}
	fmt.Println(data)
}
