package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Stopwatch Application")
	fmt.Println("=====================")
	fmt.Println("Commands:")
	fmt.Println("  start - Start the stopwatch")
	fmt.Println("  stop  - Stop the stopwatch and show elapsed time")
	fmt.Println("  exit  - Exit the application")

	var startTime, endTime time.Time
	var running bool
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nEnter a command: ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "start":
			if running {
				fmt.Println("Stopwatch is already running!")
			} else {
				startTime = time.Now()
				running = true
				fmt.Println("Stopwatch started!")
			}

		case "stop":
			if !running {
				fmt.Println("Stopwatch is not running. Use 'start' to begin.")
			} else {
				endTime = time.Now()
				elapsed := endTime.Sub(startTime)
				fmt.Printf("Elapsed time: %v\n", elapsed)
				running = false
			}

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Please use 'start', 'stop', or 'exit'.")
		}
	}
}
