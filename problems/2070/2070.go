package problem2070

import (
	"sort"
)

type item struct {
	Price         int
	HighestBeauty int
}

func MaximumBeauty(items [][]int, queries []int) []int {
	result := make([]int, len(queries))
	beautyMap := map[int]int{}
	lowestPrice := int(1e9)
	highestPrice := 0

	for _, i := range items {
		price := i[0]
		beauty := i[1]

		lowestPrice = min(lowestPrice, price)
		highestPrice = max(highestPrice, price)
		if curPrice, ok := beautyMap[price]; ok {
			beautyMap[price] = max(curPrice, beauty)
			continue
		}

		beautyMap[price] = beauty
	}

	priceWithHighestBeauty := []item{}
	for key, value := range beautyMap {
		priceWithHighestBeauty = append(priceWithHighestBeauty, item{Price: key, HighestBeauty: value})
	}

	sort.Slice(priceWithHighestBeauty, func(i, j int) bool {
		return priceWithHighestBeauty[i].Price < priceWithHighestBeauty[j].Price
	})

	l := len(priceWithHighestBeauty)
	for i := 1; i < l; i++ {
		item := &priceWithHighestBeauty[i]
		prevItem := &priceWithHighestBeauty[i-1]
		item.HighestBeauty = max(item.HighestBeauty, prevItem.HighestBeauty)
	}


	for index, q := range queries {
		bestBeauty := 0
		left, right := 0, l-1

		if q < lowestPrice {
			result[index] = 0
			continue
		}
		if q > highestPrice {
			result[index] = priceWithHighestBeauty[l-1].HighestBeauty
			continue
		}
		for left <= right {
			mid := (left + right) / 2
			if q == priceWithHighestBeauty[mid].Price {
				bestBeauty = priceWithHighestBeauty[mid].HighestBeauty
				break
			}
			if q < priceWithHighestBeauty[mid].Price {
				right = mid - 1
			} else {
				bestBeauty = max(priceWithHighestBeauty[mid].HighestBeauty, bestBeauty)
				left = mid + 1
			}
		}

		result[index] = bestBeauty
	}

	return result
}
