package problem1352

type ProductOfNumbers struct {
	m []int
}

func Constructor() ProductOfNumbers {

	return ProductOfNumbers{
		m: []int{},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.m = []int{}
		return
	}

	l := len(this.m)

	if l == 0 {
		this.m = []int{num}
		return
	}

	this.m = append(this.m, this.m[len(this.m)-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	l := len(this.m)

	if k > l {
		return 0
	}

	if k == l {
		return this.m[l-1]
	}

	return this.m[l-1] / this.m[l-k-1]
}
