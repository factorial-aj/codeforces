package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://codeforces.com/contest/1878/problem/A
func CF1878A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &k)
		a := make([]int, n)
		res := "no"
		for i := range a {
			Fscan(in, &a[i])
			if a[i] == k {
				res = "yes"
			}
		}
		Fprint(out, res, "\n")
	}
	Fprintln(out)
}

func main() { CF1878A(os.Stdin, os.Stdout) }
