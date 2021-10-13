package main

import (
	"flag"
	"fmt"
	"log"
	"unicode"

	"huffman/h"
	"huffman/heap"
	"huffman/tree"
)

func main() {

	drawHeap := flag.Bool("h", false, "dot-format heap representation on stdout")
	drawTree := flag.Bool("t", false, "dot-format encoding tree representation on stdout")
	fileName := flag.String("i", "", "file name of symbols and frequencies")
	flag.Parse()

	dict, err := h.ConstructDictFromFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	hp := h.ConstructHeapFromDict(dict)

	if *drawHeap {
		fmt.Printf("/* heap has %d elements */\n", len(hp))
		heap.Draw(hp)
		return
	}

	root := h.ConstructTree(hp)

	if *drawTree {
		tree.Draw(root)
		return
	}

	var total float64
	freqs := make(map[rune]float64)
	for _, node := range dict {
		freqs[node.Char] = node.Freq
		total += node.Freq
	}
	for r, f := range freqs {
		freqs[r] = f / total
	}

	outputEncoding(root, freqs)
	fmt.Printf("%d symbols in %d total bits, ave %.03f\n", totalSymbols, totalBits, aveBitsPerSymbol)
}

func outputEncoding(root tree.Node, freqs map[rune]float64) {
	var path []rune
	proportions = freqs
	followPath(root, path)
}

var totalSymbols int
var totalBits int
var aveBitsPerSymbol float64
var proportions map[rune]float64

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
		pro := proportions[node.(*tree.Leaf).Char]
		aveBitsPerSymbol += float64(len(path)) * pro
	}
}
