package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	Fscan(in, &n)
	var s string
	Fscan(in, &s)
	adjList := make([][]int, n)
	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 2)
		Fscan(in, &adjList[i][0], &adjList[i][1])
	}
	q := [][]int{{0, 0}}
	res := n
	for len(q) > 0 {
		node, cnt := q[0][0], q[0][1]
		q = q[1:]
		l, r := adjList[node][0], adjList[node][1]
		if l+r == 0 {
			res = min(res, cnt)
			continue
		}
		if l != 0 {
			newCnt := cnt
			if s[node] != 'L' {
				newCnt++
			}
			q = append(q, []int{l - 1, newCnt})
		}
		if r != 0 {
			newCnt := cnt
			if s[node] != 'R' {
				newCnt++
			}
			q = append(q, []int{r - 1, newCnt})
		}
	}
	Fprint(out, res)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve(in, out)
		Fprint(out, "\n")
	}
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
