// Link:https: //leetcode.com/problems/boats-to-save-people
package main

import (
	"fmt"
	"sort"
)

func Solution881(people []int, limit int) int {
	sort.Slice(people, func(i, j int) bool {
		return people[i] > people[j]
	})
	fmt.Printf("people: %v\n", people)
	count := 0
	i := 0
	j := len(people) - 1
	for i <= j {
		if people[i]+people[j] <= limit {
			j--
		}
		i++
		count++
	}

	return count
}
