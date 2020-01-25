package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	n := 3
	if arg != "" {
		nArg, err := strconv.Atoi(arg)

		if err != nil {
			fmt.Println("Wrong argument, using n = 3")
			n = 3
		} else {
			n = nArg
		}
	}

	var towers [3]*Tower

	for i := range towers {
		towers[i] = NewTower(i)
	}

	for i := n; i > 0; i-- {
		towers[0].Add(i)
	}

	fromTower := towers[0]
	buffTower := towers[1]
	destTower := towers[2]

	callbackFunc := displayTowersCallback(fromTower, buffTower, destTower)

	fromTower.SetCallback(callbackFunc)
	buffTower.SetCallback(callbackFunc)
	destTower.SetCallback(callbackFunc)

	// Display Initial Configuration
	// fromTower.ShowTower()
	// buffTower.ShowTower()
	// destTower.ShowTower()
	// callbackFunc()

	towers[0].MoveDisks(n, towers[2], towers[1])

	// Display Resulting Configuration
	// fromTower.ShowTower()
	// buffTower.ShowTower()
	// destTower.ShowTower()
	// callbackFunc()
}

func displayTowersCallback(fromTower *Tower, buffTower *Tower, destTower *Tower) func() {
	return func() {
		fromTower.ShowTower()
		buffTower.ShowTower()
		destTower.ShowTower()
	}
}
