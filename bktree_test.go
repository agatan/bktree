package bktree

import (
	"fmt"
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
