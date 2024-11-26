package problem2924

func findChampion(n int, edges [][]int) int {
	m := map[int][]int{}

	if n == 1 {
		return 0
	}

	for _, edge := range edges {
		parent := edge[0]
		child := edge[1]
		if _, ok := m[parent]; !ok {
			m[parent] = []int{}
		}
		if marked, ok := m[child]; ok {
			m[child] = append(marked, parent)
		} else {
			m[child] = []int{parent}
		}
	}

	if len(m) != n {
		return -1
	}

	r := []int{}
	for team, weakerTeams := range m {
		if len(weakerTeams) != 0 {
			continue
		}
		r = append(r, team)
	}

	if len(r) == 1 {
		return r[0]
	}

	return -1
}

func Execute(n int, edges [][]int) int {
	return findChampion(n, edges)
}
