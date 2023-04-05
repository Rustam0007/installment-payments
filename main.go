package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	req1 := Request{
		ProductName: "phone",
		Amount:      1000,
		Phone:       "+992919010101",
		Period:      18,
	}

	req2 := Request{
		ProductName: "computer",
		Amount:      1000,
		Phone:       "+992919010101",
		Period:      24,
	}

	req3 := Request{
		ProductName: "tv",
		Amount:      1000,
		Phone:       "+992919010101",
		Period:      24,
	}

	response1, err := InstallmentPayments(req1)
	if err != nil {
		log.Println(err)
	}
	response2, err := InstallmentPayments(req2)
	response3, err := InstallmentPayments(req3)

	if err != nil {
		log.Println("Error - ", err)
	}
	out, _ := json.MarshalIndent(response1, "", "  ")
	out2, _ := json.MarshalIndent(response2, "", "  ")
	out3, _ := json.MarshalIndent(response3, "", "  ")

	fmt.Println(string(out), string(out2), string(out3))
}
