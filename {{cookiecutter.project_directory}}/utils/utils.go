package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

// PrintPretty is a function to print pretty
func PrintPretty(i interface{}) {
	empJSON, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(empJSON))
}
