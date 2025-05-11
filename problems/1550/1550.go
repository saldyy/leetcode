package problem1550

func threeConsecutiveOdds(arr []int) bool {
    for i := range(len(arr) - 2) {
			if arr[i] % 2 == 1 && arr[i + 1] % 2 == 1 && arr[i + 2] % 2 == 1 {
				return true
			}
		}
    return false
}

func Execute(a []int) bool {

  return threeConsecutiveOdds(a)
}
