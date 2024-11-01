package mysort

func MergeSort(arr []int) {
	mergeSortImpl(arr, 0, len(arr)-1)
}

func mergeSortImpl(arr []int, l, r int) {
	if l < r {
		mid := (l + r) / 2
		mergeSortImpl(arr, l, mid)
		mergeSortImpl(arr, mid+1, r)
		merge(arr, l, mid, r)
	}
}

func merge(arr []int, l, mid, r int) {
	p, q, tot := l, mid+1, 0
	toolArr := make([]int, r-l+1)
	for p <= mid && q <= r {
		if arr[p] < arr[q] {
			toolArr[tot] = arr[p]
			tot = tot + 1
			p = p + 1
		} else {
			toolArr[tot] = arr[q]
			tot = tot + 1
			q = q + 1
		}
	}
	for p <= mid {
		toolArr[tot] = arr[p]
		tot = tot + 1
		p = p + 1
	}
	for q <= r {
		toolArr[tot] = arr[q]
		tot = tot + 1
		q = q + 1
	}
	copy(arr[l:r+1], toolArr)
}
