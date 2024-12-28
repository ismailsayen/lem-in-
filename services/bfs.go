package services

import (
	"slices"
)

// bfs function implements the Breadth-First Search algorithm to find paths.
// This BFS does not consider squares of a room that are duplicated on other paths. It only ensures that the rooms are unique within each path.
func (g *GraphData) BFS(neighborStart string) {
	var queue [][]string
	var currentPath []string
	visited := make(map[string]bool)
	queue = append(queue, []string{g.Start, neighborStart})
	visited[g.Start] = true
	visited[neighborStart] = true
	for len(queue) > 0 {
		currentPath = queue[0]
		queue = queue[1:]
		lastRoom := currentPath[len(currentPath)-1]
		if lastRoom == g.End {
			path := &PathInfos{len: len(currentPath), Path: currentPath[1:]}
			g.Paths = append(g.Paths, path)
			break
		}

		for _, neighbor := range g.Tunnels[lastRoom] {
			if !slices.Contains(currentPath, neighbor) && !visited[neighbor] {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
				if neighbor != g.End {
					visited[neighbor] = true
				}
			}
		}
	}
}

// CombBfs finds additional paths and adds them to the group if they match its criteria.
// This second BFS function, which takes a group with one path, is designed to find all other paths that can include this path, ensuring no duplicated rooms across all paths.
func (g *GraphData) CombBFS(grp *Groups) {
	var queue [][]string
	var currentPath []string
	queue = append(queue, []string{g.Start})
	for len(queue) > 0 {
		currentPath = queue[0]
		queue = queue[1:]
		lastRoom := currentPath[len(currentPath)-1]
		if len(g.Tunnels[lastRoom]) > 4 {
			continue
		}
		if g.End == lastRoom {
			if Unique(grp, currentPath[1:]) {
				path := &PathInfos{len: len(currentPath), Path: currentPath[1:]}
				grp.Comb = append(grp.Comb, path)
				grp.lenPaths++
				continue
			}
		}

		for _, neighbor := range g.Tunnels[lastRoom] {
			if !slices.Contains(currentPath, neighbor) {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
}
