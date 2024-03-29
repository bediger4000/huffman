package main

/*
 * Read a file (named on command line), output a table
 * of up to 256 hex values and their associated frequencies
 * on stdout. Table suitable for cmd/encode.go -t input.
 */

import (
	"fmt"
	"log"
	"os"
)

func main() {
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var byteCount [256]int
	totalByteCount := 0.0

	for i := range buf {
		byteCount[buf[i]]++
		totalByteCount += 1.0
	}

	for i := range byteCount {
		if byteCount[i] == 0 {
			continue
		}
		fmt.Printf("%02x %.08f\n", i, float64(byteCount[i])/totalByteCount)
	}
}
