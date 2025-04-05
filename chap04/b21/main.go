package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func main() {
	defer wr.Flush()
	buf := make([]byte, bufio.MaxScanTokenSize)
	sc.Buffer(buf, 1024*1024)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	s := []rune(sc.Text())

	fmt.Fprintln(wr, maxLspLen(s, n))
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxLspLen(s []rune, n int) int {
	var dp [1009][1009]int

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for len := 2; len <= n; len++ {
		for i := 0; i <= n-len; i++ {
			j := i + len - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}
