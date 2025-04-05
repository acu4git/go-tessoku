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

	a := make([]int, n+1)
	for i := 2; i <= n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}

	b := make([]int, n+1)
	for i := 3; i <= n; i++ {
		sc.Scan()
		b[i], _ = strconv.Atoi(sc.Text())
	}

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[1] = 0

	for i := 1; i <= n; i++ {
		if i <= n-1 {
			dp[i+1] = min(dp[i]+a[i+1], dp[i+1])
		}
		if i <= n-2 {
			dp[i+2] = min(dp[i]+b[i+2], dp[i+2])
		}
	}

	fmt.Fprintln(wr, dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
