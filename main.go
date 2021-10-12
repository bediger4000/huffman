package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"huffman/heap"
	"huffman/tree"
)

func main() {
	h, err := constructHeap(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("/* heap has %d elements */\n", len(h))

	for len(h) > 1 {
		var hn1, hn2 heap.Node

		h, hn1 = h.Delete()
		h, hn2 = h.Delete()

		in1 := &tree.Interior{
			Left:  hn1.(tree.Node),
			Right: hn2.(tree.Node),
			Freq:  hn1.Value() + hn2.Value(),
		}

		h = h.Insert(in1)
	}

	h, root := h.Delete()

	tree.Draw(root.(tree.Node))
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
			fmt.Fprintf(os.Stderr, "line %d parse %d fields\n", lineNo, n)
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
