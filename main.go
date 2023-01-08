package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bekcpear/hidepass/pkg/hidepass"
)

func main() {
	config := "./regex.json"
	args := os.Args[1:]
	var str string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-c":
			config = args[i+1]
			args[i+1] = ""
		case "":
			continue
		default:
			if str == "" {
				str = args[i]
			} else {
				log.Fatalln("only one string acceptable")
			}
		}
	}

	if str == "" {
		log.Fatalln("no input string")
	}

	err := hidepass.ReadConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(hidepass.Hide(str))
}
