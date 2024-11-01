package main

import (
	"SortImpl4Go/mysort"
	"fmt"
	"math/rand"
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

	fmt.Printf("Sorting %v Elements Using Following Methods:\n", N)

	t1 := time.Now()
	mysort.QuickSort(arr1)
	t2 := time.Now()
	fmt.Printf("quick sort: %v\n", t2.Sub(t1))

	t1 = time.Now()
	mysort.MergeSort(arr2)
	t2 = time.Now()
	fmt.Printf("merge sort: %v\n", t2.Sub(t1))

	t1 = time.Now()
	mysort.HeapSort(arr3)
	t2 = time.Now()
	fmt.Printf("heap sort: %v\n", t2.Sub(t1))

	input := "Press Any Key To Continue..."
	fmt.Println(input)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
}
