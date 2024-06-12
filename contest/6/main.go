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
		var numOrders int
		fmt.Fscan(in, &numOrders)

		arrival := make([][2]int, numOrders)
		for j := 0; j < numOrders; j++ {
			fmt.Fscan(in, &arrival[j][0])
			arrival[j][1] = j
		}

		sort.Slice(arrival, func(i, j int) bool {
			return arrival[i][0] < arrival[j][0]
		})

		var numCars int
		fmt.Fscan(in, &numCars)

		carCharacteristics := make([][4]int, numCars)
		for j := 0; j < numCars; j++ {
			fmt.Fscan(in, &carCharacteristics[j][0], &carCharacteristics[j][1], &carCharacteristics[j][2])
			carCharacteristics[j][3] = j
		}

		sort.Slice(carCharacteristics, func(i, j int) bool {
			if carCharacteristics[i][0] == carCharacteristics[j][0] {
				return carCharacteristics[i][3] < carCharacteristics[j][3]
			}
			return carCharacteristics[i][0] < carCharacteristics[j][0]
		})

		carSchedule := make([]int, numOrders)
		for _, arr := range arrival {
			for j := 0; j < len(carCharacteristics); j++ {
				if carCharacteristics[j][0] <= arr[0] && arr[0] <= carCharacteristics[j][1] && carCharacteristics[j][2] > 0 {
					carSchedule[arr[1]] = carCharacteristics[j][3] + 1
					carCharacteristics[j][2]--
					break
				}
			}
		}

		for _, car := range carSchedule {
			if car == 0 {
				fmt.Fprint(out, -1, " ")
			} else {
				fmt.Fprint(out, car, " ")
			}
		}
		fmt.Fprintln(out)
	}
}
