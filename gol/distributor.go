package gol

import (
	//"fmt"
	"strconv"
	"time"
	"uk.ac.bris.cs/gameoflife/util"
	"sync"
)

type distributorChannels struct {
	events     chan<- Event
	ioCommand  chan<- ioCommand
	ioIdle     <-chan bool
	ioFilename chan<- string
	ioInput    chan uint8
	tickTurn	chan int
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

func executeATurn(threadedWorld [][]byte, p Params, results chan<- [][]byte, waiter *sync.WaitGroup){
	defer waiter.Done()

	threadedWorld = calculateNextState(p, threadedWorld)

	results <- threadedWorld
}

func splitWorld(p Params, world[][]byte, threadWidth float64) [][][]byte{
	isDecimal := threadWidth != float64(int(threadWidth))
	//fmt.Println(isDecimal)
	//fmt.Println(threadWidth)

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
		//fmt.Printf("start: %v, end: %v \n", startX, endX)
		newWorlds[i] = world[startX:endX]
	}
	return newWorlds
}

func backgroundTicker(c distributorChannels, world [][]byte) {
	ticker := time.NewTicker(2 * time.Second)
	for _ = range ticker.C {
		turn := <- c.tickTurn
		c.events <- AliveCellsCount{turn, len(calculateAliveCells(world))}
	}
}

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

	go backgroundTicker(c, world)

	// make slice of channels to collect results
	results := make([]chan [][]byte, p.Threads)

	for i := 0; i < p.Threads; i++ {
		results[i] = make(chan [][]byte, 1)
	}

	// calculate width of each section of world depending on number of threads
	//this will give threadWidth as decimal
	threadWidth := float64(p.ImageWidth)/float64(p.Threads)

	// go through each turn, go through each thread, execute turn
	for i := 1; i <= p.Turns; i++ {
		c.tickTurn <- i

		//fmt.Printf("Turn: %v \n", i)
		// split world according to given threadWidth
		newWorlds := splitWorld(p, world, threadWidth)

		var combinedWorld [][]byte

		for j := 0; j < p.Threads; j++ {
			waiter.Add(1)
			go executeATurn(newWorlds[j], p, results[j], &waiter)
		}
		waiter.Wait()
		for j := 0; j < p.Threads; j++ {
			output := <- results[j]
			if len(output) != int(threadWidth) {
				output = output[1:]
			}
			combinedWorld = append(combinedWorld, output...)
		}


		aliveCells := calculateAliveCells(combinedWorld)
		deadCells := calculateDeadCells(combinedWorld)

		for _, cell := range aliveCells{
			if !isAlive(cell, world) {
				c.events <- CellFlipped{i-1, cell}
			}
		}

		for _, cell := range deadCells{
			if isAlive(cell, world) {
				c.events <- CellFlipped{i-1, cell}
			}
		}

		world = combinedWorld

		c.events <- TurnComplete{i}
	}

	c.events <- FinalTurnComplete{p.Turns, calculateAliveCells(world)}

	// TODO: Send correct Events when required, e.g. CellFlipped, TurnComplete and FinalTurnComplete.
	// 	See event.go for a list of all events.

	// Make sure that the Io has finished any output before exiting.
	c.ioCommand <- ioCheckIdle
	<-c.ioIdle

	c.events <- StateChange{p.Turns, Quitting}
	// Close the channel to stop the SDL goroutine gracefully. Removing may cause deadlock.
	close(c.events)
}
