package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Draft - algorithm and implementation steps:
// create type Cell for abstracting the coordinators for each element.
// need a way to handle the directions for checking 8 directions for the neighbours, can create directs slice
// 1. read from input file to a set to record all the live cells
// 2. build a candidate set 9 times than current live cells.
// 3. count all candidate cells alive neighbors
// 4. create next state cells' set, decide which cell should be live and put in the set
//    a. for all current alive cells, if alive neighbors >= 2 and <=3, still live
//    b. for all dead cells, if alive neighbors = 3, become live
// 5. run 10 iteration then write the output to the life 1.06 format file.

type Cell struct {
	X int64
	Y int64
}

func nextGeneration(currents map[Cell]bool) map[Cell]bool {
	// candidates represents all the possible cells for current state
	candidates := make(map[Cell]bool)
	// directions are 8 directions one cell can move
	directions := []Cell{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}

	// create all the possible cells for the next state
	for c := range currents {
		for _, d := range directions {
			// TODO: can check exists or not to avoid duplicate assignment
			candidates[Cell{c.X + d.X, c.Y + d.Y}] = true
		}
		// add all current alive cells
		candidates[c] = true
	}

	next := make(map[Cell]bool)

	for c := range candidates {
		// count alive neighbors
		// TODO: can remove to another countLiveNeighbors function
		count := 0
		for _, d := range directions {
			neighbor := Cell{c.X + d.X, c.Y + d.Y}
			if _, exists := currents[neighbor]; exists {
				count++
			}
		}

		if _, exists := currents[c]; exists {
			// for current alive cells, if alive neighbors >= 2 and <=3, still live
			if count >= 2 && count <= 3 {
				next[c] = true
			}
		} else {
			// for all dead cells, if alive neighbors = 3, become live
			if count == 3 {
				next[c] = true
			}
		}
	}
	return next
}

func readFile(r io.Reader) map[Cell]bool {
	live := make(map[Cell]bool)
	scanner := bufio.NewScanner(r)
	// skip the header line
	// TODO: check some edge cases for invalid inputs
	if scanner.Scan() {
	}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		x, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			continue
		}
		y, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			continue
		}
		cell := Cell{
			X: x,
			Y: y,
		}
		live[cell] = true
	}
	// TODO: handle scanner errors
	return live
}

func printOutput(cells map[Cell]bool) {
	fmt.Println("#Life 1.06")
	for c := range cells {
		fmt.Printf("%v %v\n", c.X, c.Y)
	}
}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not open file %s: %v\n", fileName, err)
		os.Exit(1)
	}
	defer file.Close()
	live := readFile(file)
	// TODO: customize the iteration rounds
	for i := 0; i < 10; i++ {
		live = nextGeneration(live)
	}
	printOutput(live)
	// TODO: game visualization
	// TODO: clean up the comments to make them more professional
	// TODO: unit tests
	// TODO: log errors
}
