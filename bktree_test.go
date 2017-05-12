package bktree

import (
	"fmt"
	"math/rand"
	"testing"

	popcount "github.com/hideo55/go-popcount"
)

type testEntry uint64

func (e testEntry) Distance(x Entry) int {
	a := uint64(e)
	b := uint64(x.(testEntry))

	return int(popcount.Count(a ^ b))
}

func TestEmptySearch(t *testing.T) {
	var tree BKTree
	results := tree.Search(testEntry(0), 0)
	if len(results) != 0 {
		t.Fatalf("empty tree should return empty results, bot got %d results", len(results))
	}
}

func TestExactMatch(t *testing.T) {
	var tree BKTree
	for i := 0; i < 100; i++ {
		tree.Add(testEntry(i))
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("searching %d", i), func(st *testing.T) {
			results := tree.Search(testEntry(i), 0)
			if len(results) != 1 {
				st.Fatalf("exact match should return only one result, but got %d results (%#v)", len(results), results)
			}
			if results[0].Distance != 0 {
				st.Fatalf("exact match result should have 0 as Distance field, but got %d", results[0].Distance)
			}
			if int(results[0].Entry.(testEntry)) != i {
				st.Fatalf("expected result entry value is %d, but got %d", i, int(results[0].Entry.(testEntry)))
			}
		})
	}
}

func TestFuzzyMatch(t *testing.T) {
	var tree BKTree
	for i := 0; i < 100; i++ {
		tree.Add(testEntry(i))
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("searching %d", i), func(st *testing.T) {
			results := tree.Search(testEntry(i), 2)
			for _, result := range results {
				if result.Distance > 2 {
					st.Fatalf("Distance fields of results should be less than or equal to 2, but got %d", result.Distance)
				}
				if result.Entry.Distance(testEntry(i)) > 2 {
					st.Fatalf("distances of result entries should be less than or equal to 2, but got %d", result.Distance)
				}
			}
		})
	}
}

func BenchmarkConstruct(b *testing.B) {
	randoms := make([]uint64, 10000)
	for i := range randoms {
		randoms[i] = rand.Uint64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tree BKTree
		for _, r := range randoms {
			tree.Add(testEntry(r))
		}
	}
}

const largeSize int = 1000000
const smallSize int = 1000

func makeRandomTree(size int) *BKTree {
	randoms := make([]int, size)
	for i := range randoms {
		randoms[i] = rand.Int()
	}
	var tree BKTree
	for _, r := range randoms {
		tree.Add(testEntry(r))
	}
	return &tree
}

func BenchmarkSearch_ExactForLargeTree(b *testing.B) {
	tree := makeRandomTree(largeSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 0)
	}
}

func BenchmarkSearch_1FuzzyForLargeTree(b *testing.B) {
	tree := makeRandomTree(largeSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 1)
	}
}

func BenchmarkSearch_2FuzzyForLargeTree(b *testing.B) {
	tree := makeRandomTree(largeSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 2)
	}
}

func BenchmarkSearch_8FuzzyForLargeTree(b *testing.B) {
	tree := makeRandomTree(largeSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 8)
	}
}

func BenchmarkSearch_32FuzzyForLargeTree(b *testing.B) {
	tree := makeRandomTree(largeSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 32)
	}
}

func BenchmarkLinearSearchForLargeSet(b *testing.B) {
	randoms := make([]uint64, largeSize)
	for i := range randoms {
		randoms[i] = rand.Uint64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		cnt := 0
		for _, c := range randoms {
			if int(popcount.Count(c^needle)) <= 1 {
				cnt++
			}
		}
	}
}

func BenchmarkSearch_ExactForSmallTree(b *testing.B) {
	tree := makeRandomTree(smallSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 0)
	}
}

func BenchmarkSearch_1FuzzyForSmallTree(b *testing.B) {
	tree := makeRandomTree(smallSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 1)
	}
}

func BenchmarkSearch_2FuzzyForSmallTree(b *testing.B) {
	tree := makeRandomTree(smallSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 2)
	}
}

func BenchmarkSearch_8FuzzyForSmallTree(b *testing.B) {
	tree := makeRandomTree(smallSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 8)
	}
}

func BenchmarkSearch_32FuzzyForSmallTree(b *testing.B) {
	tree := makeRandomTree(smallSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		tree.Search(testEntry(needle), 32)
	}
}

func BenchmarkLinearSearchForSmallSet(b *testing.B) {
	randoms := make([]uint64, smallSize)
	for i := range randoms {
		randoms[i] = rand.Uint64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := rand.Uint64()
		cnt := 0
		for _, c := range randoms {
			if int(popcount.Count(c^needle)) <= 1 {
				cnt++
			}
		}
	}
}
