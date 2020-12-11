package gol

import (
	"strconv"
	"sync"
	"uk.ac.bris.cs/gameoflife/util"
)

type distributorChannels struct {
	events     chan<- Event
	ioCommand  chan<- ioCommand
	ioIdle     <-chan bool
	ioFilename chan<- string
	ioInput    chan uint8
	//ioOutput	chan uint8
	//tickTurn	chan int
	//tickFinish chan bool
}


func findAliveNeighbours(world [][]byte, col int, row int) int {
	aliveNeighbours := 0
	for _, i := range []int{-1,0,1} {
		for _, j := range []int{-1,0,1} {
			if i == 0 && j == 0 {
				continue
			}

			x := (col+i+len(world))%len(world)
			y := (row+j+len(world[0]))%len(world[0])

			living := world[x][y] !=0
			if living {
				aliveNeighbours++
			}
		}
	}
	return aliveNeighbours
}

func calculateNextState(world [][]byte) [][]byte {

	newWorld := make([][]byte, len(world))

	for i := range newWorld {
		newWorld[i] = make([]byte, len(world[i]))
		copy(newWorld[i], world[i])
	}

	for col := 0; col < len(world); col++ {
		for row := 0; row < len (world[0]); row++ {
			aliveNeighbours := findAliveNeighbours(world, col, row)

			if world[col][row] !=0 {

				if aliveNeighbours < 2 {
					newWorld[col][row] = 0
				}

				if aliveNeighbours > 3 {
					newWorld[col][row] = 0
				}
			}
			if world[col][row] == 0 {

				if aliveNeighbours == 3 {
					newWorld[col][row] = 255
				}
			}
		}
	}

	return newWorld
}

func calculateAliveCells(world [][]byte) []util.Cell {

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


func calculateDeadCells(world [][]byte) []util.Cell {

	var deadCells []util.Cell

	for x, col := range world {
		for y, v := range col {
			if v == 0 {
				deadCells = append(deadCells, util.Cell{y, x})
			}
		}
	}

	return deadCells
}


func isAlive(givenCell util.Cell, world [][]byte) bool {
	x := givenCell.X
	y := givenCell.Y

	if world[x][y] == 255 {
		return true
	} else {
		return false
	}
}

func executeATurn(threadedWorld [][]byte, results chan<- [][]byte, waiter *sync.WaitGroup){
	defer waiter.Done()

	threadedWorld = calculateNextState(threadedWorld)

	results <- threadedWorld
}

func splitWorld(p Params, world[][]byte, threadWidth float64) [][][]byte{
	// identify if threadWidth is decimal, in order to compensate
	isDecimal := threadWidth != float64(int(threadWidth))

	// make 3D slice to store the 2D slices
	newWorlds := make([][][]byte, p.Threads)

	for i := 0; i < p.Threads; i++ {
		startX := int(threadWidth)*i
		endX := int(threadWidth)*(i+1)

		if (startX > 0) {
			startX--
		}
		if (i == (p.Threads-1)) && (isDecimal) {
			endX = endX + (p.ImageWidth%p.Threads)
		}

		// slice world according to calculated indices
		newWorlds[i] = world[startX:endX]
	}

	// return 3D slice of 2D slices
	return newWorlds
}

// removed code for ticker task
/*
func backgroundTicker(c distributorChannels, world [][]byte) {
	ticker := time.NewTicker(2 * time.Second)

	for _ = range ticker.C {
		select {
		case <- c.tickFinish:
			return
		default:
			turn := <- c.tickTurn
			c.events <- AliveCellsCount{turn, len(calculateAliveCells(world))}
		}
	}
}

 */

// distributor divides the work between workers and interacts with other goroutines.
func distributor(p Params, c distributorChannels) {

	//create WaitGroup
	var waiter sync.WaitGroup

	// Create a 2D slice to store the world.
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
			if world[i][j] != 0 {
				c.events <- CellFlipped{0, util.Cell{i, j}}
				//send a cell flipped event
			}
		}
	}

	//go backgroundTicker(c, world)

	// make slice of channels to collect results
	results := make([]chan [][]byte, p.Threads)

	for i := 0; i < p.Threads; i++ {
		results[i] = make(chan [][]byte, 1)
	}

	// calculate width of each section of world depending on number of threads
	// this will give threadWidth as decimal
	threadWidth := float64(p.ImageWidth)/float64(p.Threads)

	// go through each turn, go through each thread, execute turn
	for i := 1; i <= p.Turns; i++ {
		//c.tickTurn <- i

		// split world according to given threadWidth
		newWorlds := splitWorld(p, world, threadWidth)

		// create a 2D slice to store the processed world in
		var combinedWorld [][]byte

		// add to waitgroup, execute turn with a world from the split worlds and a results channel
		for j := 0; j < p.Threads; j++ {
			waiter.Add(1)
			go executeATurn(newWorlds[j], results[j], &waiter)
		}

		//wait for processing to finish before continuing
		waiter.Wait()

		// collect results from each results channel
		// if the resulting 2D slice is longer than expected (including overlap), remove the leading slice
		// append corrected slice to combinedworld
		for j := 0; j < p.Threads; j++ {
			output := <- results[j]
			if len(output) != int(threadWidth) {
				output = output[1:]
			}
			combinedWorld = append(combinedWorld, output...)
		}

		// calculate slices of all the alive and dead cells in the new world
		aliveCells := calculateAliveCells(combinedWorld)
		deadCells := calculateDeadCells(combinedWorld)

		// if a cell is alive, but was dead in the previous state, send a CellFlipped event
		for _, cell := range aliveCells{
			if !isAlive(cell, world) {
				c.events <- CellFlipped{i, cell}
			}
		}

		// if a cell is dead, but was alive in the previous state, send a CellFlipped event
		for _, cell := range deadCells{
			if isAlive(cell, world) {
				c.events <- CellFlipped{i, cell}
			}
		}

		// replace the original world with the newly processed world
		world = combinedWorld

		// mark turn as complete
		c.events <- TurnComplete{i}
	}

	// removed code for ticker task
	/*
	c.tickFinish <- true
	close(c.tickFinish)
	close(c.tickTurn)
	 */

	// mark that final turn has been completed
	c.events <- FinalTurnComplete{p.Turns, calculateAliveCells(world)}

	// removed code to output final state of board as pgm
	/*a
	for i := 0; i < p.ImageHeight; i++ {
		for j := 0; j < p.ImageWidth; j++ {
			c.ioOutput <- world[i][j]
		}
	}
	 */

	// Make sure that the Io has finished any output before exiting.
	c.ioCommand <- ioCheckIdle
	<-c.ioIdle

	c.events <- StateChange{p.Turns, Quitting}
	// Close the channel to stop the SDL goroutine gracefully. Removing may cause deadlock.
	close(c.events)
}
