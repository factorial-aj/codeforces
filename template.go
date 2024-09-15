package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var x int
	Fscan(in, &x)
	Fprint(out, "yo", " ", x)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve(in, out)
	}
	Fprintln(out)
}