package main

import (
	"SortImpl4Go/mysort"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	N int = 1e6
)

func main() {
	arr1 := make([]int, N)
	arr2 := make([]int, N)
	arr3 := make([]int, N)
	for i := 0; i < N; i++ {
		arr1[i] = rand.Int()
		arr2[i] = arr1[i]
		arr3[i] = arr1[i]
	}
	t1 := time.Now()
	mysort.QuickSort(arr1)
	t2 := time.Now()
	fmt.Println(sort.IntsAreSorted(arr1))
	fmt.Println(t2.Sub(t1))
	t1 = time.Now()
	mysort.MergeSort(arr2)
	t2 = time.Now()
	fmt.Println(sort.IntsAreSorted(arr2))
	fmt.Println(t2.Sub(t1))
	t1 = time.Now()
	mysort.HeapSort(arr3)
	t2 = time.Now()
	fmt.Println(sort.IntsAreSorted(arr3))
	fmt.Println(t2.Sub(t1))
}
