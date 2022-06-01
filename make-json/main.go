package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	var inputName, inputAddress string

	inputName = readString("Enter a name")
	inputAddress = readString("Enter an address")

	// Add the name and address to a map
	person := make(map[string]string)
	person["name"] = inputName
	person["address"] = inputAddress

	// Create a JSON object from the map
	myJson, _ := json.Marshal(person)
	fmt.Println(string(myJson))
}

func readString(s string) string {
	for {
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput != "" {
			return userInput
		}
	}
}
