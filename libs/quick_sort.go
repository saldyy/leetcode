package libs

const (
	asc  = "asc"
	desc = "desc"
)

func partition(arr []int, low int, high int, order string) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {

		isValid := false
		if isValid = arr[j] < pivot; order == desc {
			isValid = arr[j] > pivot
		}
		if isValid {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// 3 7 2 8
func QuickSort(arr []int, low int, high int, order string) {
	if order != asc && order != desc {
		panic("Invalid order")
	}
	if low < high {
		pi := partition(arr, low, high, order)
		QuickSort(arr, low, pi-1, order)
		QuickSort(arr, pi+1, high, order)
	}
}
