package problem1865

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	m1    map[int]int
	m2    map[int]int
}

func toMapOccurence(nums []int) map[int]int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	return m
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	r := FindSumPairs{}
	r.nums1 = nums1
	r.nums2 = nums2
	r.m1 = toMapOccurence(nums1)
	r.m2 = toMapOccurence(nums2)

	return r
}

func (this *FindSumPairs) Add(index int, val int) {
	num := this.nums2[index]
	this.m2[num]--
	this.m2[num+val]++
	this.nums2[index] += val
}

func (this *FindSumPairs) Count(tot int) int {
	ans := 0

	for num, occurence := range this.m1 {
		ans += occurence * this.m2[(tot-num)]
	}

	return ans
}
