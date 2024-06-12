package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"math"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numTests int
	fmt.Fscan(in, &numTests)

	for i := 0; i < numTests; i++ {
		var numCars, carCapacity int
		fmt.Fscan(in, &numCars, &carCapacity)

		var numBoxes int
		fmt.Fscan(in, &numBoxes)
		boxesWeights := make([]int, numBoxes)

		for i := 0; i < numBoxes; i++ {
			var weight int
			fmt.Fscan(in, &weight)
			boxesWeights[i] = int(math.Pow(2, float64(weight)))
		}

		sort.Ints(boxesWeights)

		numTransportations, numLoadedCars, curWeight := 0, 0, 0
		for len(boxesWeights) > 0 {
			if curWeight + boxesWeights[len(boxesWeights)-1] <= carCapacity {
				curWeight += boxesWeights[len(boxesWeights)-1]
				boxesWeights = boxesWeights[:len(boxesWeights)-1]
			} else {
				for j := len(boxesWeights) - 1; j >= 0; j-- {
					if curWeight + boxesWeights[j] <= carCapacity {
						curWeight += boxesWeights[j]
						boxesWeights = append(boxesWeights[:j], boxesWeights[j+1:]...)
					}
				}
				numLoadedCars++
				if numLoadedCars >= numCars {
					numTransportations++
					numLoadedCars = 0
				}
				curWeight = 0
			}
		}

		if curWeight > 0 || numLoadedCars > 0 {
			numTransportations++
		}

		fmt.Fprintln(out, numTransportations)
	}
}
