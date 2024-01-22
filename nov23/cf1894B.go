package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://codeforces.com/contest/1894/problem/B
func CF1894B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		res := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
			res[i] = 1
		}
		num2Idx := make(map[int][]int)
		for i, num := range a {
			num2Idx[num] = append(num2Idx[num], i)
		}
		k := 2
		for _, idxs := range num2Idx {
			if len(idxs) >= 2 {
				res[idxs[0]] = k
				k++
				if k > 3 {
					break
				}
			}
		}
		if k <= 3 {
			Fprint(out, -1, "\n")
			continue
		}
		for i := 0; i < n; i++ {
			Fprint(out, res[i], " ")
		}
		Fprint(out, "\n")
	}
	Fprintln(out)
}

func main() { CF1894B(os.Stdin, os.Stdout) }
