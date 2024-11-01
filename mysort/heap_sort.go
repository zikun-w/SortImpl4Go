package mysort

func HeapSort(arr []int) {
	heapSortImpl(arr)
}

func heapSortImpl(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}
	for j := len(arr) - 1; j >= 0; j-- {
		swap(arr, 0, j)
		adjustHeap(arr, 0, j)
	}
}

func adjustHeap(arr []int, i int, length int) {
	curElement := arr[i]
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && arr[k] < arr[k+1] {
			k = k + 1
		}
		if arr[k] > curElement {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = curElement
}
