package journeytomoon

func journeyToMoon(n int32, astronaut [][]int32) int32 {
	arr := make([][]int, len(astronaut))
	for i := range astronaut {
		arr[i] = make([]int, len(astronaut[i]))
		for j := range astronaut[i] {
			arr[i][j] = int(astronaut[i][j])
		}
	}
	return int32(JourneyToMoon(int(n), arr))
}

func JourneyToMoon(n int, astronaut [][]int) int {
	components := splitIntoConnected(n, astronaut)
	var pairs int
	for _, component := range components {
		pairs += len(component) * (n - len(component))
	}
	return pairs / 2
}

func splitIntoConnected(n int, edges [][]int) [][]int {
	visited := make(map[int]bool)
	edgesMap := make([][]int, n)
	for _, edge := range edges {
		edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
		edgesMap[edge[1]] = append(edgesMap[edge[1]], edge[0])
	}
	var components [][]int
	var component []int
	var f func(v int)
	f = func(v int) {
		visited[v] = true
		component = append(component, v)
		for _, node := range edgesMap[v] {
			if !visited[node] {
				f(node)
			}
		}
	}
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		component = nil
		f(i)
		components = append(components, component)
	}
	return components
}
