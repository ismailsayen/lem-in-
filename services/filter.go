package services

import (
	"log"
	"slices"
)

// GroupMaker creates groups for each path by initializing a group for every path in g.Paths
// and then combining them using the CombBfs function.
func (g *GraphData) GroupMaker() {
	for _, path := range g.Paths {
		group := &Groups{key: path, lenPaths: 1}
		g.Groups = append(g.Groups, group)
	}
	// make the g.Paths empty
	g.Paths = []*PathInfos{}
	for _, grp := range g.Groups {
		g.CombBFS(grp)
	}
}

// Unique checks if the given path (currentPath) shares any common room with the group's key path
func Unique(p *Groups, currentPath []string) bool {
	if len(currentPath) == 1 {
		return false
	}
	for i := 0; i < len(currentPath)-1; i++ {
		if slices.Contains(p.key.Path, currentPath[i]) {
			return false
		}
		for j := 0; j < len(p.Comb); j++ {
			if slices.Contains(p.Comb[j].Path, currentPath[i]) {
				return false
			}
		}
	}
	return true
}

// FilterPaths give the max len group
func (g *GraphData) FilterPaths() {
	var matrix [][]string

	help := false

	for _, grp := range g.Groups {
		if grp.lenPaths == 1 {
			continue
		}
		g.Paths = nil
		g.Paths = append(g.Paths, grp.key)
		g.Paths = append(g.Paths, grp.Comb...)

		for _, e := range g.Paths {
			matrix = append(matrix, e.Path)
		}

		status := g.IsGoodGroup(matrix)

		if status {
			help = true
			g.Paths = append(g.Paths, grp.key)
			g.Paths = append(g.Paths, grp.Comb...)
			break
		}

		matrix = nil
	}

	if !help {

		if len(g.Groups) < 1 {
			log.Fatal("errr")
		}

		max := g.Groups[0]
		for _, grp := range g.Groups {
			if grp.lenPaths > max.lenPaths {
				max = grp
			}
		}
		g.Paths = append(g.Paths, max.key)
		g.Paths = append(g.Paths, max.Comb...)
		g.Groups = []*Groups{}
	}
}

func (y *GraphData) IsGoodGroup(matrix [][]string) bool {
	return y.NbOfAnts <= CoUnt(matrix)
}

func CoUnt(g [][]string) int {
	c := 0
	for _, v := range g {
		for range v {
			c++
		}
	}
	return c
}
