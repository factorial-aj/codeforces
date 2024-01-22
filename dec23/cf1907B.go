package main

import (
	"bufio"
	. "fmt"
	"os"
	"bytes"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s string
	Fscan(in, &s)
	b, B := 0, 0
	var buffer bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'b' {
			b++
		} else if s[i] == 'B' {
			B++
		} else if 'a' <= s[i] && s[i] <= 'z' {
			if b > 0 {
				b--
				continue
			}
			buffer.WriteByte(s[i])
		} else {
			if B > 0 {
				B--
				continue
			}
			buffer.WriteByte(s[i])
		}
	}
	ret := buffer.Bytes()
	var res bytes.Buffer
	for i := len(ret) - 1; i >= 0; i-- {
		res.WriteByte(ret[i])
	}
	Fprint(out, res.String())
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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
