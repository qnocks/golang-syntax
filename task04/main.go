package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
)

func processWorker(workerID int, data chan string) {
	// if channel is close workers will be stopped
	for num := range data {
		fmt.Printf("#%d got - %s\n", workerID, num)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Missing the command line argument - number of workers")
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Error converring command line argument: %s to int\n", args[0])
	}

	reader := bufio.NewReader(os.Stdin)

	// channel to transfer data
	data := make(chan string)

	// channel to handle Ctrl+C signal
	interrupt := make(chan os.Signal, 1)

	// assign interrupt signals to the interrupt channel
	signal.Notify(interrupt, os.Interrupt)

	for i := 0; i < n; i++ {
		// start workers
		go processWorker(i, data)
	}

	for {
		// handle Ctrl+C signal
		select {
		case <-interrupt:
			// close data channel to stop workers
			close(data)
			return
		default:
		}

		// read user input
		input, _ := reader.ReadString('\n')

		// write data to the channel
		data <- input
	}
}
