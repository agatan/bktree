package bktree

type BKTree struct {
	root *node
}

type node struct {
	entry    Entry
	children map[int]*node
}

type Entry interface {
	Distance(Entry) int
}

func (bk *BKTree) Add(entry Entry) {
	if bk.root == nil {
		bk.root = &node{
			entry:    entry,
			children: make(map[int]*node),
		}
		return
	}
	nd := bk.root
	for {
		d := nd.entry.Distance(entry)
		if next, ok := nd.children[d]; ok {
			nd = next
			return
		}
		nd.children[d] = &node{
			entry:    entry,
			children: make(map[int]*node),
		}
	}
}

type Result struct {
	Distance int
	Entry    Entry
}

func (bk *BKTree) Search(needle Entry, tolerance int) []*Result {
	results := make([]*Result, 0)
	if bk.root == nil {
		return results
	}
	candidates := []*node{bk.root}
	for len(candidates) != 0 {
		c := candidates[len(candidates)-1]
		candidates = candidates[:len(candidates)-1]
		d := c.entry.Distance(needle)
		if d <= tolerance {
			results = append(results, &Result{
				Distance: d,
				Entry:    c.entry,
			})
		}

		low, high := d-tolerance, d+tolerance
		for k, child := range c.children {
			if low <= k && k <= high {
				candidates = append(candidates, child)
			}
		}
	}
	return results
}
