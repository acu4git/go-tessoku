package main

import (
	"bufio"
	"fmt"
	"math"
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

	a := make([]int, n)
	b := make([]int, n)
	sc.Scan()
	for i, v := range strings.Fields(sc.Text()) {
		a[i+1], _ = strconv.Atoi(v)
	}
	sc.Scan()
	for i, v := range strings.Fields(sc.Text()) {
		b[i+1], _ = strconv.Atoi(v)
	}

	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = math.MinInt // required
	}
	for i := 1; i < n; i++ {
		dp[a[i]] = max(dp[i]+100, dp[a[i]])
		dp[b[i]] = max(dp[i]+150, dp[b[i]])
	}

	fmt.Fprintln(wr, dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
