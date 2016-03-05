package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	defaultDuration = 1 * time.Second
)

func main() {
	rate := flag.Duration("rate", defaultDuration, "The maximum output rate.")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	for {
		fmt.Print(text)
		next := time.Now().Add(*rate)
		for next.After(time.Now()) {
			text, _ = reader.ReadString('\n')
		}
	}
}
