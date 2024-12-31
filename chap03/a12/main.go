package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func getInt() int {
	sc.Scan()
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func getInts(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = getInt()
	}
	return res
}

func input(n, k *int, arr *[]int) {
	*n = getInt()
	*k = getInt()
	*arr = getInts(*n)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)

	var n, k int
	var arr []int
	input(&n, &k, &arr)

	solve(n, k, arr)
}

func solve(n, k int, arr []int) {
	sort.Ints(arr)

	left := 1
	right := 1_000_000_000
	for left < right {
		mid := (left + right) / 2
		if check(arr, mid, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Fprintln(wr, left)
}

func check(arr []int, x, k int) bool {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += x / arr[i]
		if sum >= k {
			return true
		}
	}
	return false
}
