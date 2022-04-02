package dynamicconnectivity

import (
	"fmt"
)

type QuickFind struct {
	indexes  []int
	capacity int
	sz       []int
}

// NewQuickFind index is data set
func NewQuickFind(n int) *QuickFind {
	ids := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
		sz[i] = 1
	}
	return &QuickFind{capacity: n, indexes: ids, sz: sz}
}

// Connected p and q are connected if they have the same id
func (qf *QuickFind) Connected(p, q int) bool {
	return qf.indexes[p] == qf.indexes[q]
}

// Union assign id of q to id of p. O(N)= M * N
func (qf *QuickFind) Union(p, q int) {
	ids := qf.indexes
	pid, qid := ids[p], ids[q]

	for i := range ids {
		if ids[i] == pid {
			ids[i] = qid
		}
	}
}

// root get root of passed index
func (qf *QuickFind) root(index int) int {
	for index != qf.indexes[index] {
		index = qf.indexes[index]
	}
	return index
}

// QuickConnected p and q are connected if they have the same root
func (qf *QuickFind) QuickConnected(p, q int) bool {
	return qf.root(p) == qf.root(q)
}

// QuickUnion assign q's root to p's root. O(N) = N + M*logN
func (qf *QuickFind) QuickUnion(p, q int) {
	i := qf.root(p)
	j := qf.root(q)
	qf.indexes[i] = j
}

func (qf *QuickFind) PrintData() {
	fmt.Println(qf.indexes)
}

// WeightedQuickUnion assign q's root to p's root with weighted. O(N) = N + M*logN
func (qf *QuickFind) WeightedQuickUnion(p, q int) {
	i := qf.root(p)
	j := qf.root(q)
	if i == j {
		return
	}
	if qf.sz[i] < qf.sz[j] {
		qf.indexes[i] = j
		qf.sz[j] += qf.sz[i]
	} else {
		qf.indexes[j] = i
		qf.sz[i] += qf.sz[j]
	}
}

func (qf *QuickFind) rootWithPathCompression(index int) int {
	for index != qf.indexes[index] {
		qf.indexes[index] = qf.indexes[qf.indexes[index]]
		index = qf.indexes[index]
	}
	return index
}

// WeightedQuickUnionWithPathCompression assign q's root to p's root with weighted and path Compression. O(N) = N + M*logN
func (qf *QuickFind) WeightedQuickUnionWithPathCompression(p, q int) {
	i := qf.rootWithPathCompression(p)
	j := qf.rootWithPathCompression(q)
	if i == j {
		return
	}
	if qf.sz[i] < qf.sz[j] {
		qf.indexes[i] = j
		qf.sz[j] += qf.sz[i]
	} else {
		qf.indexes[j] = i
		qf.sz[i] += qf.sz[j]
	}
}
