package main

import (
	"go-os/aws"
	"log"
	"os/exec"
)

func main() {
	log.Println("Testing the logging: in main")
	bytes, err := aws.Version(exec.Command)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(bytes))
}
