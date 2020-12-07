package gol

import (
	"fmt"
	"strconv"
	"uk.ac.bris.cs/gameoflife/util"
	"sync"
)

type distributorChannels struct {
	events     chan<- Event
	ioCommand  chan<- ioCommand
	ioIdle     <-chan bool
	ioFilename chan<- string
	ioInput    chan uint8
}


func findAliveNeighbours(world [][]byte, col int, row int) int {
	aliveNeighbours := 0
	for _, i := range []int{-1,0,1} {
		for _, j := range []int{-1,0,1} {
			if i == 0 && j == 0 {
				continue
			}

			living := world[(col+i+len(world))%len(world)][(row+j+len(world[0]))%len(world[0])] !=0
			if living {
				aliveNeighbours++
			}
		}
	}
	return aliveNeighbours
}

func calculateNextState(p Params, world [][]byte) [][]byte {

	newWorld := make([][]byte, len(world))

	for i := range newWorld {
		newWorld[i] = make([]byte, len(world[i]))
		copy(newWorld[i], world[i])
	}

	for col := 0; col < len(world); col++ {
		for row := 0; row < len (world[0]); row++ {
			/*find number of alive neighbours */
			aliveNeighbours := findAliveNeighbours(world, col, row)

			if world[col][row] !=0 {
				/* if current cell is not dead */

				if aliveNeighbours < 2 {
					newWorld[col][row] = 0
				}

				if aliveNeighbours > 3 {
					newWorld[col][row] = 0
				}
			}
			if world[col][row] == 0 {
				/* if current cell is dead */

				if aliveNeighbours == 3 {
					newWorld[col][row] = 0xFF
				}
			}
		}
	}

	return newWorld
}

func calculateAliveCells(p Params, world [][]byte) []util.Cell {

	var aliveCells []util.Cell

	for x, col := range world {
		for y, v := range col {
			if v != 0 {
				aliveCells = append(aliveCells, util.Cell{y, x})
			}
		}
	}

	return aliveCells
}

// don't necessarily need to update the same file you read from??
func executeATurn(threadedWorld [][]byte, p Params, results chan<- [][]byte, waiter *sync.WaitGroup){
	defer waiter.Done()

	threadedWorld = calculateNextState(p, threadedWorld)

	results <- threadedWorld
}

// distributor divides the work between workers and interacts with other goroutines.
func distributor(p Params, c distributorChannels) {

	//create WaitGroup
	var waiter sync.WaitGroup

	// TODO: Create a 2D slice to store the world.
	world := make([][]byte, p.ImageHeight)
	for i := range world {
		world[i] = make([]byte, p.ImageWidth)
	}

	height := strconv.Itoa(p.ImageHeight)
	width := strconv.Itoa(p.ImageWidth)

	c.ioCommand <- ioInput
	c.ioFilename <- height + "x" + width

	// add input to world, send cellflipped event for each initially alive cell

	for i := 0; i < p.ImageHeight; i++ {
		for j := 0; j < p.ImageWidth; j++ {
			world[i][j] = <-c.ioInput
			if world[i][j] == 255 {
				c.events <- CellFlipped{0, util.Cell{i, j}}
				//send a cell flipped event
			}
		}
	}

	// make single channel to collect results

	results := make(chan [][]byte, p.Threads)

	// make slices within slice to create split up worlds

	newWorlds := make([][][]byte, p.Threads)

	//each slice should have "width" of p.ImageWidth/p.Threads

	threadWidth := p.ImageWidth/p.Threads

	fmt.Println(threadWidth)

	for i := 0; i < p.Threads; i++ {
		newWorlds[i] = world
		fmt.Printf("original world %v %v \n", i, len(newWorlds[i]))

		startX := threadWidth*i
		if startX > 0 {
			startX--
		}
		endX := threadWidth*(i+1)

		fmt.Printf("start %v, end %v \n", startX, endX)

		//make sure there is an additional margin for correct processing
		columns := newWorlds[i][startX:endX]
		fmt.Printf("sliced world %v %v \n", i, len(columns))
	}

	fmt.Println("World sliced")

	// go through each turn, go through each thread, execute turn

	combinedWorld := make([][]byte, p.ImageHeight)

	for i := 1; i <= p.Turns; i++ {
		for j := 0; j < p.Threads; j++ {
			waiter.Add(1)
			go executeATurn(newWorlds[j], p, results, &waiter)
			combinedWorld = append(combinedWorld, <-results...)
		}
		fmt.Printf("combined world turn %v: %v \n", i, len(combinedWorld))
		waiter.Wait()

		aliveCells := calculateAliveCells(p, combinedWorld)


		for _, cell := range aliveCells{
			fmt.Printf("cell flipping: %v \n", cell)
			c.events <- CellFlipped{i, cell}
		}


		fmt.Printf("number of alive cells: %v \n", len(aliveCells))

		c.events <- TurnComplete{i}
	}

	c.events <- FinalTurnComplete{p.Turns, calculateAliveCells(p, world)}

	// TODO: Send correct Events when required, e.g. CellFlipped, TurnComplete and FinalTurnComplete.
	// 	See event.go for a list of all events.

	// Make sure that the Io has finished any output before exiting.
	c.ioCommand <- ioCheckIdle
	<-c.ioIdle

	c.events <- StateChange{p.Turns, Quitting}
	// Close the channel to stop the SDL goroutine gracefully. Removing may cause deadlock.
	close(c.events)
}
