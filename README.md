## bktree

This is an implementation of [BK-tree](https://en.wikipedia.org/wiki/BK-tree) for golang.
BK-tree is a tree data structure for similarity search in a metric space.
Using BK-tree, you can search neighbors of a data from the metric space efficiently.

### Performance

Search similar values from 1,000,000 data.
Data is 64 bits integer, and distance function is hamming distance.
(see `bktree_test.go` for detail)

```
BenchmarkSearch_ExactForLargeTree-4              1000000              1108 ns/op
BenchmarkSearch_1FuzzyForLargeTree-4               50000             29468 ns/op
BenchmarkSearch_2FuzzyForLargeTree-4                5000            328753 ns/op
BenchmarkSearch_8FuzzyForLargeTree-4                  20          68182122 ns/op
BenchmarkSearch_32FuzzyForLargeTree-4                  3         353715305 ns/op

BenchmarkLinearSearchForLargeSet-4                   300           4132926 ns/op
```

If the tolerance is small enough, BK-tree is much faster than naive linear search.


### Example

see `_example/` direcotry.

### Install

```
$ go get github.com/agatan/bktree
```

### License

MIT

