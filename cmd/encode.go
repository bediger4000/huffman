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

	encoding, err := constructEncoding(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	bwo, err := NewBitwiseOut(*outFileName, encoding)
	if err != nil {
		log.Fatal(err)
	}

	n := flag.NArg()
	if n < 1 {
		log.Fatal("need filename of symbols to encode\n")
		return
	}

	symbolFile := flag.Arg(0)
	fmt.Fprintf(os.Stderr, "encoding contents of %q\n", symbolFile)

	plainBuffer, err := os.ReadFile(symbolFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range []rune(string(plainBuffer)) {
		bwo.Write(r)
	}

	bwo.Flush()
}

type EncodedSymbol struct {
	Symbol   rune
	Bits     []byte
	BitCount int // len(Bits) but it's handy
}

func constructEncoding(freqFileName string) (map[rune]*EncodedSymbol, error) {
	hp, err := h.ConstructHeapFromFile(freqFileName)
	if err != nil {
		return nil, err
	}

	encodingTree := h.ConstructTree(hp)

	return findEncoding(encodingTree), nil
}

func findEncoding(root tree.Node) map[rune]*EncodedSymbol {
	var path []byte
	followPath(root, path)
	fmt.Fprintf(os.Stderr, "Found %d leaf nodes\n", len(leaves))
	encoding := make(map[rune]*EncodedSymbol)
	for _, sym := range leaves {
		encoding[sym.Symbol] = sym
	}
	return encoding
}

var leaves []*EncodedSymbol

func followPath(node tree.Node, path []byte) {
	switch node.(type) {
	case *tree.Interior:
		newpath := append(path, 0)
		followPath(node.LeftChild(), newpath)
		newpath = append(path, 1)
		followPath(node.RightChild(), newpath)
	case *tree.Leaf:
		npath := make([]byte, len(path))
		copy(npath, path)
		sym := &EncodedSymbol{
			Symbol:   node.(*tree.Leaf).Char,
			Bits:     npath,
			BitCount: len(npath),
		}
		leaves = append(leaves, sym)
	}
}

type BitwiseOutput struct {
	FileName  string
	Encoding  map[rune]*EncodedSymbol
	Fout      *os.File
	ByteCount int
	BitCount  int
	RuneCount int

	currentByte byte
	bitsInByte  int
}

func NewBitwiseOut(outFileName string, encoding map[rune]*EncodedSymbol) (*BitwiseOutput, error) {
	fout, err := os.Create(outFileName)
	if err != nil {
		return nil, err
	}
	return &BitwiseOutput{
		FileName: outFileName,
		Encoding: encoding,
		Fout:     fout,
	}, nil
}

func (bwo *BitwiseOutput) Write(r rune) {
	bwo.RuneCount++
	encoded, ok := bwo.Encoding[r]
	if !ok {
		fmt.Fprintf(os.Stderr, "No encoding for symbol #%d, %c (%0x)\n", bwo.RuneCount, r, r)
		return
	}
	for i := 0; i < encoded.BitCount; i++ {
		if bwo.bitsInByte == 8 {
			bwo.ByteCount++
			if _, err := bwo.Fout.Write([]byte{bwo.currentByte}); err != nil {
				fmt.Fprintf(os.Stderr,
					"Output byte %d, rune %d (%c): %v\n",
					bwo.ByteCount, bwo.RuneCount, r, err,
				)
				return
			}
			bwo.currentByte = 0
			bwo.bitsInByte = 0
		}
		bwo.currentByte <<= 1
		bwo.currentByte |= (0x01 & encoded.Bits[i])
		bwo.bitsInByte++
		bwo.BitCount++
	}
}

func (bwo *BitwiseOutput) Flush() {
	bitsPerRune := float64(bwo.BitCount) / float64(bwo.RuneCount)
	if bwo.bitsInByte == 0 {
		fmt.Fprintf(os.Stderr, "Output %d bits in %d bytes, representing %d runes, %.02f bits/rune\n", bwo.BitCount, bwo.ByteCount, bwo.RuneCount, bitsPerRune)
		return
	}
	fmt.Fprintf(os.Stderr, "Flushing final %d bits\n", bwo.bitsInByte)
	bwo.currentByte <<= (8 - bwo.bitsInByte)
	bwo.ByteCount++
	if _, err := bwo.Fout.Write([]byte{bwo.currentByte}); err != nil {
		fmt.Fprintf(os.Stderr,
			"Output byte %d, rune %d: %v\n",
			bwo.ByteCount, bwo.RuneCount, err,
		)
	}
	fmt.Fprintf(os.Stderr, "Output %d bits in %d bytes, representing %d runes, %02f bits/rune\n", bwo.BitCount, bwo.ByteCount, bwo.RuneCount, bitsPerRune)
	bwo.Fout.Close()
}
