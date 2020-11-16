package gol

// Params provides the details of how to run the Game of Life and which image to load.
type Params struct {
	Turns       int
	Threads     int
	ImageWidth  int
	ImageHeight int
}

type cell struct {
	x, y int
}

func calculateNextState(p Params, world [][]byte) [][]byte {
	newWorld := make([][]byte, len(world))
	for i := range newWorld {
		newWorld[i] = make([]byte, len(world[i]))
		copy(newWorld[i], world[i])
	}


	return world
}

func calculateAliveCells(p Params, world [][]byte) []cell {
	return []cell{}
}

// Run starts the processing of Game of Life. It should initialise channels and goroutines.
func Run(p Params, events chan<- Event, keyPresses <-chan rune) {

	ioCommand := make(chan ioCommand)
	ioIdle := make(chan bool)

	distributorChannels := distributorChannels{
		events,
		ioCommand,
		ioIdle,
	}
	go distributor(p, distributorChannels)

	ioChannels := ioChannels{
		command:  ioCommand,
		idle:     ioIdle,
		filename: nil,
		output:   nil,
		input:    nil,
	}
	go startIo(p, ioChannels)
}
