package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var p, a [2009]int
	for i := 1; i <= n; i++ {
		sc.Scan()
		tokens := strings.Fields(sc.Text())
		p[i], _ = strconv.Atoi(tokens[0])
		a[i], _ = strconv.Atoi(tokens[1])
	}

	fmt.Fprintln(wr, getMaxScore(n, p, a))
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func getMaxScore(n int, p, a [2009]int) int {
	var dp [2009][2009]int

	// dp[l][r] := max(dp[l-1][r] + score1, dp[l][r+1] + score2)
	dp[1][n] = 0
	for l := 1; l <= n; l++ {
		for r := n; r >= 1; r-- {
			if (l == 1 && r == n) || l > r {
				continue
			}

			score1 := 0
			if l <= p[l-1] && p[l-1] <= r {
				score1 = a[l-1]
			}

			score2 := 0
			if l <= p[r+1] && p[r+1] <= r {
				score2 = a[r+1]
			}

			dp[l][r] = max(dp[l-1][r]+score1, dp[l][r+1]+score2)
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		if dp[i][i] >= res {
			res = dp[i][i]
		}
	}

	return res
}
