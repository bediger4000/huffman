package main

import (
	"fmt"
	"huffman/heap"
	"huffman/tree"
	"io"
	"log"
	"os"
)

func main() {
	h, err := constructHeap(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("/* heap has %d elements */\n", len(h))
	heap.Draw(h)
}

func constructHeap(fileName string) (heap.Heap, error) {
	fin, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer fin.Close()

	var h heap.Heap

	for lineNo := 1; true; lineNo++ {
		var ch rune
		var freq float64

		n, err := fmt.Fscanf(fin, "%c %f\n", &ch, &freq)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if n != 2 {
			fmt.Fprintf(os.Stderr, "line %d parse %d fields\n", n)
			continue
		}
		node := &tree.Leaf{
			Freq: freq,
			Char: rune(ch),
		}
		h = h.Insert(node)
	}

	return h, nil
}
