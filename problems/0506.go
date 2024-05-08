package main

import (
	"sort"
	"strconv"
)

func Solution506(score []int) []string {
	c := make([]int, len(score))
	result := make([]string, len(score))
	copy(c[:], score[:])

	sort.Ints(c)

	for index, val := range score {
		i := sort.SearchInts(c, val)

		if i == len(score)-1 {
			result[index] = "Gold Medal"
		} else if i == len(score)-2 {
			result[index] = "Silver Medal"
		} else if i == len(score)-3 {
			result[index] = "Bronze Medal"
		} else {
			result[index] = strconv.Itoa(len(score) - i)
		}

	}
	return result
}
