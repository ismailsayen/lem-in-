package services

// We assign the shortest path to each ant.
func (g *GraphData) assignAntPaths() {
	// this is for bad link with the start and the end room.
	if g.BadLink && g.NbOfAnts == 1 {
		var p []string
		p = append(p, g.End)
		g.OneAnt = append(g.OneAnt, p)
		g.PrintTurns(g.OneAnt)
		return
	}

	// 
	ants := make([][]string, g.NbOfAnts)

	for i := 1; i <= g.NbOfAnts; i++ {

		shortPath := g.Paths[0]
		for j := 0; j < len(g.Paths); j++ {
			if g.Paths[j].len < shortPath.len {
				shortPath = g.Paths[j]
			}
		}

		shortPath.len++

		ants[i-1] = append(ants[i-1], shortPath.Path...)
	}
	// Here we mark the paths directly linked directly to the start room.
	for i := 0; i < len(ants); i++ {
		if len(ants[i]) == 1 {
			ants[i] = append(ants[i], g.End)
		}
	}

	g.PrintTurns(ants)
}
