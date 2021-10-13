# Daily Coding Problem: Problem #1016 [Easy]

This problem was asked by Amazon.

Huffman coding is a method of encoding characters based on their frequency.
Each letter is assigned a variable-length binary string,
such as 0101 or 111110,
where shorter lengths correspond to more common letters.
To accomplish this,
a binary tree is built such that the path from the root to any leaf
uniquely maps to a character.
When traversing the path,
descending to a left child corresponds to a 0 in the prefix,
while descending right corresponds to 1.

Here is an example tree (note that only the leaf nodes have letters):

```
        *
      /   \
    *       *
   / \     / \
  *   a   t   *
 /             \
c               s
```

With this encoding, cats would be represented as 0000110111.

Given a dictionary of character frequencies,
build a Huffman tree,
and use it to determine a mapping between characters
and their encoded binary strings.

## Build and run

```sh
$ go build .
$ ./huffman -t -i better.tbl > better.dot
$ dot -Tpng -o better.png better.dot
$ ./huffman -i better.tbl > better.encodings
```

The `huffman` program can output [GraphViz](https://graphviz.org/) `dot` format representations
of both the min-priority tree it creates before constructing the
encoding tree, and the encoding tree.

Without a `-t` (dot-format encoding tree) or `-h` (dot-format heap) flag,
the `huffman` program prints out human readable table of symboly numeric values
and corresponding bits for each value.
I believe this satisfies the problem statement.

* `english.tbl` is a very simple proportion of letters in English text
table. Doesn't include anything other than upper-case ASCII values,
but that makes the encoding tree simpler to navigate.
* `better.tbl` is a byte frequency table based on a bunch of english text
I found on my laptop a while back. I also use this table of byte frequency
in [single-byte Xor decoding](https://github.com/bediger4000/singlexor),
to determine "closeness" of possibly decoded text to english text.

## Analysis

"[Easy]".
I had no idea how to do this,
even after pondering for a while.
The [Huffman coding](https://en.wikipedia.org/wiki/Huffman_coding) wikipedia page
gives a "simple algorithm" that involves building a tree from a minimum priority queue,
where the priority of a character is that character's frequency.
Once you have the tree,
you can determine the bit patterns for each character by traversing the tree.
A left child gets a 0-bit, a right child gets a 1-bit at each level of recursion.
When the traverse reaches a leaf node,
it has all the bits for the symbol at that leaf.

## Interview Analysis

This really does not strike me as "[Easy]".
It involves 2 different data structures, a heap (used as a min-priority queue)
and a binary tree.
There's algorithms for adding and deleting elements to the min-priority queue,
and there's a traverse of the tree built from the queue,
which involves keeping track of left or right child recursion.

The O(n) method of constructing the encoding tree is completely inobvious.

Unless the candidate has recently gone through either of these algorithms,
they're not going to do well at this problem at all.
