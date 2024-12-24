package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func input(n, x *int64, a *[]int64) error {
	var err error
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)

	sc.Scan()
	token := strings.Fields(sc.Text())
	*n, err = strconv.ParseInt(token[0], 10, 0)
	if err != nil {
		return err
	}
	*x, err = strconv.ParseInt(token[1], 10, 0)
	if err != nil {
		return err
	}

	sc.Scan()
	*a = make([]int64, 0, *n)
	for _, token := range strings.Fields(sc.Text()) {
		v, err := strconv.ParseInt(token, 10, 0)
		if err != nil {
			return err
		}
		*a = append(*a, v)
	}

	return nil
}

func main() {
	var n, x int64
	var a []int64
	if err := input(&n, &x, &a); err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(a, x))
}

func solve(arr []int64, target int64) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			return mid + 1
		}
	}

	return -1
}
