package problem2070

import (
	"reflect"
	"testing"
)

type testCase struct {
	items    [][]int
	queries  []int
	expected []int
}

func TestMaximumBeauty(t *testing.T) {
	cases := []testCase{
		// {
		// 	items:    [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}},
		// 	queries:  []int{1, 2, 3, 4, 5, 6},
		// 	expected: []int{2, 4, 5, 5, 6, 6},
		// },
		// {
		// 	items:    [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 2}},
		// 	queries:  []int{1},
		// 	expected: []int{4},
		// },
		// {
		// 	items:    [][]int{{10, 10000}},
		// 	queries:  []int{5},
		// 	expected: []int{0},
		// },
		// {
		// 	items:    [][]int{{3, 3}, {3, 4}, {3, 5}, {1, 3}},
		// 	queries:  []int{1, 2},
		// 	expected: []int{3, 3},
		// },
		{
			items:    [][]int{{193, 732}, {781, 962}, {864, 954}, {749, 627}, {136, 746}, {478, 548}, {640, 908}, {210, 799}, {567, 715}, {914, 388}, {487, 853}, {533, 554}, {247, 919}, {958, 150}, {193, 523}, {176, 656}, {395, 469}, {763, 821}, {542, 946}, {701, 676}},
			queries:  []int{885, 1445, 1580, 1309, 205, 1788, 1214, 1404, 572, 1170, 989, 265, 153, 151, 1479, 1180, 875, 276, 1584},
			expected: []int{962, 962, 962, 962, 746, 962, 962, 962, 946, 962, 962, 919, 746, 746, 962, 962, 962, 919, 962},
		},
	}

	for _, value := range cases {
		result := MaximumBeauty(value.items, value.queries)

		if !reflect.DeepEqual(result, value.expected) {
			t.Log("Running test case")
			t.Logf("Items: %v\n", value.items)
			t.Logf("Queries: %v\n", value.queries)
			t.Errorf("Incorrect, expecting: %v, received: %v", value.expected, result)
		}
	}

}
