package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numTests int
	fmt.Fscan(in, &numTests)

	for i := 0; i < numTests; i++ {
		var numPlayers int
		fmt.Fscan(in, &numPlayers)

		type player struct {
			time int
			idx  int
		}

		times := make([]player, numPlayers)
		for j := 0; j < numPlayers; j++ {
			fmt.Fscan(in, &times[j].time)
			times[j].idx = j
		}

		sort.Slice(times, func(i, j int) bool {
			return times[i].time < times[j].time
		})

		places := make([]int, numPlayers)
		place, prevTime, numSamePlaces := 0, -1, 1
		for _, p := range times {
			if p.time-prevTime > 1 {
				place += numSamePlaces
				numSamePlaces = 1
			} else {
				numSamePlaces++
			}
			places[p.idx] = place
			prevTime = p.time
		}

		for _, place := range places {
			fmt.Fprint(out, place, " ")
		}
		fmt.Fprintln(out)
	}
}
