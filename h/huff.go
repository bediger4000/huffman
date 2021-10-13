package h

import (
	"fmt"
	"huffman/heap"
	"huffman/tree"
	"io"
	"os"
)

func ConstructDictFromFile(fileName string) ([]*tree.Leaf, error) {
	fin, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer fin.Close()

	var dict []*tree.Leaf

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
		dict = append(dict, node)
	}

	return dict, nil
}

func ConstructHeapFromFile(fileName string) (heap.Heap, error) {
	dict, err := ConstructDictFromFile(fileName)
	if err != nil {
		return nil, err
	}

	return ConstructHeapFromDict(dict), nil
}

func ConstructHeapFromDict(dict []*tree.Leaf) heap.Heap {
	var hp heap.Heap
	for _, node := range dict {
		hp = hp.Insert(node)
	}
	return hp
}

func ConstructTree(hp heap.Heap) tree.Node {
	for len(hp) > 1 {
		var hn1, hn2 heap.Node

		hp, hn1 = hp.Delete()
		hp, hn2 = hp.Delete()

		in1 := &tree.Interior{
			Left:  hn1.(tree.Node),
			Right: hn2.(tree.Node),
			Freq:  hn1.Value() + hn2.Value(),
		}

		hp = hp.Insert(in1)
	}

	hp, root := hp.Delete()

	return root.(tree.Node)
}
