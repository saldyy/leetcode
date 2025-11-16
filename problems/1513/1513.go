package problem1513

const mod = 1e9 + 7

func numSub(s string) int {
	ans := 0
	c := 0
	for _, val := range s {
		if val == '0' {
			ans += ((c + 1) * c / 2) % mod
			c = 0
			continue
		}
		c++
	}

	ans += ((c + 1) * c / 2) % mod

	return ans
}

func Execute(a string) int {

	return numSub(a)
}
