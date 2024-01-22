package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://codeforces.com/contest/1894/problem/A
func CF1894A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		var res string
		Fscan(in, &res)
		Fprint(out, string(res[n-1]), "\n")
	}
	Fprintln(out)
}

func main() { CF1894A(os.Stdin, os.Stdout) }