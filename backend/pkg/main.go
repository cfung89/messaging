package main

import (
	"fmt"
	"os"

	"github.com/cfung89/messaging/backend/pkg/assert"
)

func main() {
	iss, err := os.ReadFile("./iss.env")
	if err != nil {
		fmt.Printf("Unable to read issuer from file: %s\n", err)
	}
	a := string(iss)
	fmt.Println(a)
	b := assert.EqualIn{}
}
