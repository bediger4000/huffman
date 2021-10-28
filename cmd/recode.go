package main

import (
	"flag"
	"fmt"
	"huffman/h"
	"huffman/tree"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	fileName := flag.String("i", "", "file name of symbols and frequencies")
	outFileName := flag.String("o", "", "output file name")
	max := flag.Int("N", 1024*1024, "number of input bits")
	flag.Parse()

	if *fileName == "" {
		log.Fatal("need symbols and frequencies file (-i filename)\n")
	}

	if *outFileName == "" {
		log.Fatal("need name of output file (-o filename)\n")
	}

	fout, err := os.Create(*outFileName)
	if err != nil {
		log.Fatal(err)
	}

	hp, err := h.ConstructHeapFromFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	encodingTree := h.ConstructTree(hp)

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	node := encodingTree
	bits := NewBitSource()

	symbolOutputCount := 0
	bitOutputCount := 0

	for i := 0; i < *max; i++ {
		if bit := bits.NextBit(); bit == 0 {
			node = node.LeftChild()
		} else {
			node = node.RightChild()
		}
		bitOutputCount++
		if l, ok := node.(*tree.Leaf); ok {
			fmt.Fprintf(fout, "%c", l.Char)
			symbolOutputCount++
			node = encodingTree
		}
	}
	fmt.Fprintf(os.Stderr, "Output %d symbols in %d bits\n", symbolOutputCount, bitOutputCount)
}

type BitSource struct {
	current uint64
	used    int
}

func NewBitSource() *BitSource {
	return &BitSource{
		current: rand.Uint64(),
	}
}

func (b *BitSource) NextBit() byte {
	if b.used == 64 {
		b.used = 0
		b.current = rand.Uint64()
	}
	bit := byte(b.current & 0x01)
	b.current >>= 1
	b.used++
	return bit
}
