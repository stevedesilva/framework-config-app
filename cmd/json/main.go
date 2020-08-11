package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	fmt.Println("json")
	file, err := os.Open("cmd/json/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed to read file, error msg: %v", err)
	}

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatalf("Failed to decode file, error msg: %v", err)
	}
	fmt.Printf("Path %v, Enabled %t\n", conf.Path, conf.Enabled)

}
