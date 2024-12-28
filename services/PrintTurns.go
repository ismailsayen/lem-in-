package services

import (
	"fmt"
	"slices"
	"strings"
)

// Function that moves ants on each turn.
func (g *GraphData) PrintTurns(ants [][]string) {
	
	turns := [][]string{}
	turn := []string{}
	rooms := []string{}

	var allPathsEmpty int
	for _, path := range ants {
		allPathsEmpty += len(path)
	}
	for allPathsEmpty > 0 {

		badTunnel := false

		for i := 0; i < len(ants); i++ {
			if len(ants[i]) == 0 {
				continue
			}
			step := fmt.Sprintf("L%v-%v", i+1, ants[i][0])
			if len(ants[i]) == 2 && ants[i][0] == g.End {
				if !badTunnel {
					ants[i] = ants[i][2:]
					turn = append(turn, step)
					badTunnel = true
					allPathsEmpty -= 2
				}
			} else if !slices.Contains(rooms, ants[i][0]) {
				if ants[i][0] != g.End {
					rooms = append(rooms, ants[i][0])
				}
				ants[i] = ants[i][1:]
				turn = append(turn, step)
				allPathsEmpty--
			}

		}

		turns = append(turns, turn)
		rooms = []string{}
		turn = []string{}

	}
	// we print here...
	for _, rooms := range turns {
		fmt.Println(strings.Join(rooms, " "))
	}
}
