package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Missing the command line argument - application TTL")
	}

	// convert program ttl to int from command line arguments
	programTTL, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Error converring command line argument: %s to int", args[0])
	}

	// channel to transfer data
	data := make(chan string)
	reader := bufio.NewReader(os.Stdin)

	go func() {
		// loop through the channel and print data
		for d := range data {
			fmt.Printf("Received: %s\n", d)
		}
	}()

	// before ttl expires read user input and write to channel
	for now := time.Now(); time.Since(now) < time.Second*time.Duration(programTTL); {
		input, _ := reader.ReadString('\n')
		data <- input
	}
}
