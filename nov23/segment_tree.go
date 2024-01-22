package main

import (
	"bufio"
	. "fmt"
	"os"
)

type SegmentTree struct {
	tree 	[]int
	arr		[]int
	size	int
}

func NewSegmentTree(arr []int) *SegmentTree {
	st := &SegmentTree{arr: arr, size: len(arr)}
	st.tree = make([]int, 4*st.size)
	st.build(0, 0, st.size-1)
	return st
}

func (st *SegmentTree) build(i, s, e int) {
	if s == e {
		st.tree[i] = st.arr[s]
		return
	}
	m := s + (e-s)/2
	st.build(2*i+1, s, m)
	st.build(2*i+2, m+1, e)
	st.tree[i] = min(st.tree[2*i+1], st.tree[2*i+2])
}

func (st *SegmentTree) query(i, s, e, l, r int) int {
	if r < s || e < l {
		return 1<<31 - 1
	}
	if s >= l && r >= e {
		return st.tree[i]
	}
	m := s + (e-s)/2
	left := st.query(2*i+1, s, m, l, r)
	right := st.query(2*i+2, m+1, e, l, r)
	return min(left, right)
}

func (st *SegmentTree) update(i, s, e, k, v int) {
	if s == e && s == k {
		st.tree[i] = v
		return
	}
	m := s + (e-s)/2
	if k <= m {
		st.update(2*i+1, s, m, k, v)
	} else {
		st.update(2*i+2, m+1, e, k, v)
	}
	st.tree[i] = min(st.tree[2*i+1], st.tree[2*i+2])
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, q int
	Fscan(in, &n, &q)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &arr[i])
	}
	st := NewSegmentTree(arr)
	for ; q > 0; q-- {
		var op string
		var x, y int
		Fscan(in, &op, &x, &y)
		if op == "q" {
			Fprint(out, st.query(0, 0, n-1, x-1, y-1), "\n")
		} else {
			st.update(0, 0, n-1, x-1, y)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	solve(in, out)
	Fprintln(out)
}

func min(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v < res {
			res = v
		}
	}
	return res
}

func max(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
