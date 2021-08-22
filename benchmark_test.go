package main

import (
	"fmt"
	"testing"
	"uk.ac.bris.cs/gameoflife/gol"
)

func BenchmarkGol1(b *testing.B) {
	// run a gol program, 512x512, with 1 thread, for 100 turns
	testParams := gol.Params{100, 1, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol2(b *testing.B) {
	// run a gol program, 512x512, with 2 threads, for 100 turns
	testParams := gol.Params{100, 2, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol3(b *testing.B) {
	// run a gol program, 512x512, with 3 threads, for 100 turns
	testParams := gol.Params{100, 3, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol4(b *testing.B) {
	// run a gol program, 512x512, with 4 threads, for 100 turns
	testParams := gol.Params{100, 4, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol5(b *testing.B) {
	// run a gol program, 512x512, with 5 threads, for 100 turns
	testParams := gol.Params{100, 5, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol6(b *testing.B) {
	// run a gol program, 512x512, with 6 threads, for 100 turns
	testParams := gol.Params{100, 6, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol7(b *testing.B) {
	// run a gol program, 512x512, with 7 threads, for 100 turns
	testParams := gol.Params{100, 7, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol8(b *testing.B) {
	// run a gol program, 512x512, with 8 threads, for 100 turns
	testParams := gol.Params{100, 8, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol9(b *testing.B) {
	// run a gol program, 512x512, with 9 threads, for 100 turns
	testParams := gol.Params{100, 9, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol10(b *testing.B) {
	// run a gol program, 512x512, with 10 threads, for 100 turns
	testParams := gol.Params{100, 10, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol11(b *testing.B) {
	// run a gol program, 512x512, with 11 threads, for 100 turns
	testParams := gol.Params{100, 11, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol12(b *testing.B) {
	// run a gol program, 512x512, with 12 threads, for 100 turns
	testParams := gol.Params{100, 12, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol13(b *testing.B) {
	// run a gol program, 512x512, with 13 threads, for 100 turns
	testParams := gol.Params{100, 13, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol14(b *testing.B) {
	// run a gol program, 512x512, with 14 threads, for 100 turns
	testParams := gol.Params{100, 14, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol15(b *testing.B) {
	// run a gol program, 512x512, with 15 threads, for 100 turns
	testParams := gol.Params{100, 15, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}

func BenchmarkGol16(b *testing.B) {
	// run a gol program, 512x512, with 16 threads, for 100 turns
	testParams := gol.Params{100, 16, 512, 512}
	events := make(chan gol.Event)
	benchName := fmt.Sprintf("%dx%dx%d-%d", testParams.ImageWidth, testParams.ImageHeight, testParams.Turns, testParams.Threads)
	b.Run(benchName, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gol.Run(testParams, events, nil)
			for range events {
				// correct?
			}

		}
	})
}