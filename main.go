package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"huffman/heap"
	"huffman/tree"
)

func main() {

	drawHeap := flag.Bool("h", false, "dot-format heap representation on stdout")
	drawTree := flag.Bool("t", false, "dot-format encoding tree representation on stdout")
	fileName := flag.String("i", "", "file name of symbols and frequencies")
	flag.Parse()

	h, err := constructHeap(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	if *drawHeap {
		fmt.Printf("/* heap has %d elements */\n", len(h))
		heap.Draw(h)
		return
	}

	root := constructTree(h)

	if *drawTree {
		tree.Draw(root)
		return
	}

	outputEncoding(root)
	fmt.Printf("%d symbols in %d total bits\n", totalSymbols, totalBits)
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

		n, err := fmt.Fscanf(fin, "%02x %f\n", &ch, &freq)
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

func constructTree(h heap.Heap) tree.Node {
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

	return root.(tree.Node)
}

func outputEncoding(root tree.Node) {
	var path []rune
	followPath(root, path)
}

var totalSymbols int
var totalBits int

func followPath(node tree.Node, path []rune) {
	switch node.(type) {
	case *tree.Interior:
		newpath := append(path, '0')
		followPath(node.LeftChild(), newpath)
		newpath = append(path, '1')
		followPath(node.RightChild(), newpath)
	case *tree.Leaf:
		r := node.(*tree.Leaf).Char
		if unicode.IsPrint(r) && (unicode.IsLetter(r) || unicode.IsNumber(r)) {
			fmt.Printf("%c: %v\n", r, string(path))
		} else {
			fmt.Printf("%02x: %v\n", r, string(path))
		}
		totalSymbols++
		totalBits += len(path)
	}
}
