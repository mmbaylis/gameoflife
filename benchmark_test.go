package main

import (
	"fmt"
	"testing"
	"uk.ac.bris.cs/gameoflife/gol"
)

func BenchmarkGol(b *testing.B) {
	// run a gol program, 512x512, with 1 thread, for 25 turns
	testParams := gol.Params{25, 1, 512,512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for event := range events {
				switch event.(type) {
				case gol.FinalTurnComplete:
					//correct?
				}
			}

		}
	})

	// run a gol program, 512x512, with 1 thread, for 100 turns
	testParams = gol.Params{100, 1, 512,512}
	events = make(chan gol.Event)
	benchName = fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for event := range events {
				switch event.(type) {
				case gol.FinalTurnComplete:
					//correct?
				}
			}

		}
	})

	// run a gol program, 512x512, with 2 threads, for 25 turns
	testParams = gol.Params{25, 2, 512,512}
	events = make(chan gol.Event)
	benchName = fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for event := range events {
				switch event.(type) {
				case gol.FinalTurnComplete:
					//correct?
				}
			}

		}
	})

	// run a gol program, 512x512, with 2 threads, for 100 turns
	testParams = gol.Params{100, 2, 512,512}
	events = make(chan gol.Event)
	benchName = fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for event := range events {
				switch event.(type) {
				case gol.FinalTurnComplete:
					//correct?
				}
			}

		}
	})

	// run a gol program, 512x512, with 10 threads, for 25 turns
	testParams = gol.Params{25, 10, 512,512}
	events = make(chan gol.Event)
	benchName = fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for event := range events {
				switch event.(type) {
				case gol.FinalTurnComplete:
					//correct?
				}
			}

		}
	})

	// run a gol program, 512x512, with 10 threads, for 100 turns
	testParams = gol.Params{100, 10, 512,512}
	events = make(chan gol.Event)
	benchName = fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for event := range events {
				switch event.(type) {
				case gol.FinalTurnComplete:
					//correct?
				}
			}

		}
	})
}
