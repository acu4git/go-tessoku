package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func input(n *int64, a *[]int64, q *int64, x *[]int64) error {
	var err error
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)

	sc.Scan()
	if *n, err = strconv.ParseInt(sc.Text(), 10, 0); err != nil {
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

	sc.Scan()
	if *q, err = strconv.ParseInt(sc.Text(), 10, 0); err != nil {
		return err
	}

	*x = make([]int64, 0, *q)
	for i := int64(0); i < *q; i++ {
		sc.Scan()
		v, err := strconv.ParseInt(sc.Text(), 10, 0)
		if err != nil {
			return err
		}
		*x = append(*x, v)
	}

	return nil
}

func main() {
	var arrSize int64
	var arr []int64
	var numQuery int64
	var queries []int64
	if err := input(&arrSize, &arr, &numQuery, &queries); err != nil {
		log.Fatal(err)
	}

	solve(arr, queries)
}

func solve(arr, queries []int64) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for _, target := range queries {
		fmt.Println(lowerBound(arr, target))
	}
}

func lowerBound(arr []int64, target int64) int {
	idx := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
	return idx
}
