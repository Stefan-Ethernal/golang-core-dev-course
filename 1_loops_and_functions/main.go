package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(0.5)
	for i := 0; i < 10; i++ {
		previousZ := z
		z -= (z*z - x) / (2 * z)
		if previousZ == z {
			break
		}
		fmt.Printf("Iteration #%d, z=%f\n", i+1, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
}
