package main

import (
	"bufio"
	"fmt"
	"math"
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
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	a := make([][]int, m+1)
	for i := range a {
		a[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sc.Scan()
			a[i][j], _ = strconv.Atoi(sc.Text())
		}
	}

	// dp[i][S]: クーポン券1, 2, ..., iの中から最小何枚選べば品物の集合Sが無料になるか
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, 1<<n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt16
		}
	}

	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		for j := 0; j < (1 << n); j++ {
			coupon := 0
			for k := n; k >= 1; k-- {
				coupon <<= 1
				coupon |= a[i][k]
			}

			dp[i][j] = min(dp[i][j], dp[i-1][j])
			if (j | coupon) < (1 << n) {
				dp[i][j|coupon] = min(dp[i][j|coupon], dp[i-1][j]+1)
			}
		}
	}

	if dp[m][(1<<n)-1] == math.MaxInt16 {
		fmt.Fprintln(wr, -1)
	} else {
		fmt.Fprintln(wr, dp[m][(1<<n)-1])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
