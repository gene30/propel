package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Propel Launched")
	debug("Propel Launched.")

}

func debug(message string) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "Version 1.0b", log.LstdFlags)
	logger.Println(message)
}
