package main

import (
	"fmt"
	"log"
	"os"

	"oss.ac/hidapp/pkg/hidapp"
)

func main() {
	configFile := "./regex.json"
	args := os.Args[1:]
	var str string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-c":
			configFile = args[i+1]
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

	fd, err := os.Open(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	pp, err := hidapp.NewProcessorFrom(fd)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(pp.Process(str))
}
