package main

func GetItinerary(transfers [][]string) [][]string {
	graph := make(map[string][]string)
	count := make(map[string]int)

	for _, t := range transfers {
		graph[t[0]] = append(graph[t[0]], t[1])
		count[t[0]]++
		count[t[1]]--
	}

	start := transfers[0][0]
	for vertex, c := range count {
		if c > 0 {
			start = vertex
			break
		}
	}

	stack, path := []string{start}, make([]string, 0, len(graph))

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if len(graph[cur]) > 0 {
			last := len(graph[cur]) - 1
			stack = append(stack, graph[cur][last])
			graph[cur] = graph[cur][:last]
		} else {
			last := len(stack) - 1
			path = append(path, stack[last])
			stack = stack[:last]
		}
	}

	itinerary := make([][]string, 0, len(graph))
	for i := len(path) - 1; i > 0; i-- {
		itinerary = append(itinerary, []string{path[i], path[i-1]})
	}

	return itinerary
}

func GetTransfers(s []string) [][]string {
	var t [][]string
	i, j := 0, 0

	for i < len(s) {
		j = i + 1
		t = append(t, []string{s[i], s[j]})
		i = j + 1
	}
	return t
}

func GetCount(s [][]string) int {
	return len(s)
}

func GetDestination(transfers [][]string) []string {
	m := map[string]bool{}
	d := make([]string, 0, 2)

	for _, tr := range transfers {
		for _, t := range tr {
			_, ok := m[t]
			if ok {
				delete(m, t)
			} else {
				m[t] = true
			}
		}
	}

	for k := range m {
		d = append(d, k)
	}

	return d
}
