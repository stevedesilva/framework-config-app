package main

import (
	"fmt"

	"github.com/labstack/gommon/log"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	conf, err := yaml.ReadFile("cmd/yaml/conf.yaml")
	if err != nil {
		log.Fatalf("Failed to read file, error msg: %v", err)
	}

	fmt.Println(conf.Get("path"))
	fmt.Println(conf.GetBool("enabled"))

}
