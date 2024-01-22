package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"bytes"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	Fscan(in, &n)
	var s string
	Fscan(in, &s)
	var resultBuffer bytes.Buffer

	for i := n - 1; i >= 0; i-- {
		if s[i] == 'a' || s[i] == 'e' {
			resultBuffer.WriteByte(s[i])
			resultBuffer.WriteByte(s[i-1])
			resultBuffer.WriteString(".")
			i -= 1
		} else {
			resultBuffer.WriteByte(s[i])
			resultBuffer.WriteByte(s[i-1])
			resultBuffer.WriteByte(s[i-2])
			resultBuffer.WriteString(".")
			i -= 2
		}
	}

	resultString := resultBuffer.String()
	res := reverseString(resultString[:len(resultString)-1])

	Fprint(out, res)
}

func reverseString(s string) string {
	runeSlice := []rune(s)
	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}
	return string(runeSlice)
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

func lowerBound(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}