package quicksort

func Sort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, low, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		quicksort(arr, low, partitionIndex-1)
		quicksort(arr, partitionIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pIndex := low
	pivot := arr[high]

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			arr[pIndex], arr[i] = arr[i], arr[pIndex]
			pIndex++
		}
	}

	arr[pIndex], arr[high] = arr[high], arr[pIndex]
	return pIndex
}
