package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func piEstimation(points int) float64 {
	circlePointsInside := 0
	for i := 0; i < points; i++ {
		x := rand.Float64()
		y := rand.Float64()
		distance := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
		if distance <= 1 {
			circlePointsInside++
		}
	}
	return 4 * (float64(circlePointsInside) / float64(points))
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please introduce some point value in the cmd")
		os.Exit(1)
	}
	points, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("You need to use a valid value")
		os.Exit(1)
	}

	fmt.Println("PI Estimation: ", piEstimation(points))
}
