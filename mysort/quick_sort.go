package mysort

func QuickSort(arr []int) {
	quickSortImpl(arr, 0, len(arr)-1)
}

func quickSortImpl(arr []int, l, r int) {
	if l < r {
		lLimit, rLimit := partition(arr, l, r)
		quickSortImpl(arr, l, lLimit-1)
		quickSortImpl(arr, rLimit+1, r)
	}
}

func partition(arr []int, l, r int) (lLimit, rLimit int) {
	pivot := arr[l]
	curPos := l + 1
	lLimit = l
	rLimit = r
	for curPos <= rLimit {
		if arr[curPos] < pivot {
			swap(arr, curPos, lLimit)
			lLimit = lLimit + 1
			curPos = curPos + 1
		} else if arr[curPos] > pivot {
			swap(arr, curPos, rLimit)
			rLimit = rLimit - 1
		} else {
			curPos = curPos + 1
		}
	}
	return lLimit, rLimit
}
