package main

import (
	"flag"
	"fmt"
	"huffman/h"
	"huffman/tree"
	"log"
	"os"
)

func main() {
	fileName := flag.String("i", "", "file name of symbols and frequencies")
	outFileName := flag.String("o", "", "output file name")
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

	n := flag.NArg()
	if n < 1 {
		log.Fatal("need filename of huffman-encoded bits to decode\n")
		return
	}

	bitsFile := flag.Arg(0)
	fmt.Fprintf(os.Stderr, "decoding contents of %q\n", bitsFile)

	bitsBuffer, err := os.ReadFile(bitsFile)
	if err != nil {
		log.Fatal(err)
	}

	bits := NewBits(bitsBuffer, bitsFile)

	node := encodingTree

	symbolOutputCount := 0
	bitOutputCount := 0
	for !bits.Empty() {
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
	fmt.Fprintf(os.Stderr, "Decoded %d symbols in %d bits\n", symbolOutputCount, bitOutputCount)
}

type Bits struct {
	filename     string
	buffer       []byte
	bufferLength int
	bitInByte    int
	byteInBuffer int
}

func NewBits(buffer []byte, filename string) *Bits {
	return &Bits{
		filename:     filename,
		buffer:       buffer,
		bufferLength: len(buffer),
	}
}

func (b *Bits) NextBit() byte {
	bit := (b.buffer[b.byteInBuffer] >> (7 - b.bitInByte)) & 0x01
	b.bitInByte++
	if b.bitInByte == 8 {
		b.bitInByte = 0
		b.byteInBuffer++
	}
	return bit
}

func (b *Bits) Empty() bool {
	if b.byteInBuffer == b.bufferLength {
		return true
	}
	return false
}
