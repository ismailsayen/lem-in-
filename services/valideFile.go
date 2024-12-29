package services

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"lemin/utils"
)

type GraphData struct {
	NbOfAnts int
	Start    string
	End      string
	Rooms    map[string][]string
	Tunnels  map[string][]string
	Paths    []*PathInfos
	Groups   []*Groups
}

type Groups struct {
	key      *PathInfos
	Comb     []*PathInfos
	lenPaths int
}

type PathInfos struct {
	len  int
	Path []string
}

// Create an instance from the struct GraphData
func NewGraphData() *GraphData {
	return &GraphData{
		Rooms:   make(map[string][]string),
		Tunnels: make(map[string][]string),
	}
}

// ValidateFileContent is a method that validates the content inside the file
func (g *GraphData) ValidateFileContent(file *os.File) string {
	var err error

	scanner := bufio.NewScanner(file)
	myfile := []string{}
	var count int
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		myfile = append(myfile, line)

		// For virefy number of ants
		if count == 0 {
			g.NbOfAnts, err = strconv.Atoi(line)
			if err != nil || g.NbOfAnts <= 0 || g.NbOfAnts > 10000 {
				return "ERROR: invalid data format, invalid number of Ants"
			}

			count = 1
			continue
		}

		if line == "" || (line[0] == '#' && line != "##start" && line != "##end") {
			continue
		}

		room := strings.Fields(line)
		if count == 2 || count == 3 {
			if !utils.IsValidRoom(line) {
				if count == 2 {
					return "ERROR: invalid data format, invalid start room"
				} else {
					return "ERROR: invalid data format, invalid end room"
				}
			}
			if count == 2 {
				g.Start = room[0]
			} else {
				g.End = room[0]
			}
			count = 1
		}

		if line == "##start" {
			if g.Start != "" {
				return "ERROR: invalid data format, Duplicated Start Room"
			}
			count = 2
			continue
		}

		if line == "##end" {
			if g.End != "" {
				return "ERROR: invalid data format, duplicated end room"
			}
			count = 3
			continue
		}
		if !utils.IsValidRoom(line) && !utils.IsValidTunnel(line) {
			return "ERROR: invalid data format, invalid room or tunnel"
		}
		if utils.IsValidRoom(line) {
			err := g.AddRoom(line)
			if err != nil {
				return "ERROR: invalid data format, failed to add Room"
			}
			continue
		}
		if utils.IsValidTunnel(line) {
			err := g.AddNeighbor(line)
			if err != nil {
				return "ERROR: invalid data format, failed to add tunnel"
			}
		}

	}

	if g.Start == "" {
		return "ERROR: invalid data format, missing start room"
	}
	if g.End == "" {
		return "ERROR: invalid data format, missing end room"
	}
	if g.Start == g.End {
		return "ERROR: invalid data format, start and end rooms are identical"
	}
	g.Rooms = map[string][]string{}
	// Range over the slice of neighbors to find the shortest path for each neighbor using BFS (Breadth-First Search).
	for i := 0; i < len(g.Tunnels[g.Start]); i++ {
		g.BFS(g.Tunnels[g.Start][i])
	}

	if len(g.Paths) < 1 {
		return "ERROR: invalid data format, no valid paths found"
	}

	g.GroupMaker()

	g.Tunnels = map[string][]string{}
	// After creating the groups, we need to choose one to pass the ants, and we select the group with the most paths.
	g.FilterPaths()

	for i, ele := range myfile {
		if len(ele) == 0 {
			continue
		}
		if i != 0 && i < len(myfile)-1 {
			fmt.Println(ele, " ")
			continue
		}
		fmt.Println(ele)
	}
	fmt.Println("")
	fmt.Println("")

	g.assignAntPaths()

	return ""
}

// AddRoom adds a new room to the collection.
func (g *GraphData) AddRoom(line string) error {
	room := strings.Fields(line)
	if _, exist := g.Rooms[room[0]]; exist {
		return errors.New("ERROR: invalid data format, room already exists")
	}

	g.Rooms[room[0]] = append(g.Rooms[room[0]], room[1], room[2])
	return nil
}

// AddNeighbor adds a new tunnel to the collection.
func (g *GraphData) AddNeighbor(line string) error {
	if !utils.ContainsRoom(line, g.Rooms) {
		return errors.New("ERROR: invalid data format")
	}

	tunnel := strings.Split(line, "-")

	if slices.Contains(g.Tunnels[tunnel[0]], tunnel[1]) {
		return errors.New("ERROR: invalid data format,tunnel already exists")
	}
	if slices.Contains(g.Tunnels[tunnel[1]], tunnel[0]) {
		return errors.New("ERROR: invalid data format,tunnel already exists")
	}
	g.Tunnels[tunnel[0]] = append(g.Tunnels[tunnel[0]], tunnel[1])
	g.Tunnels[tunnel[1]] = append(g.Tunnels[tunnel[1]], tunnel[0])
	return nil
}
