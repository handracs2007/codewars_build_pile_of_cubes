package main

import (
	"fmt"
	"math"
)

func solvePositive(a, b, c int) float64 {
	var x1, x2 float64

	x1 = (float64(-b) + math.Sqrt(float64(b*b-(4*a*c)))) / float64(2*a)
	x2 = (float64(-b) - math.Sqrt(float64(b*b-(4*a*c)))) / float64(2*a)

	if x1 > 0 {
		return x1
	} else {
		return x2
	}
}

func FindNb(m int) int {
	var a = 1
	var b = 1

	var res = m * 4
	var cFloat = math.Sqrt(float64(res))
	var cInt = int(cFloat)

	if cInt*cInt != res {
		return -1
	}

	var c = -cInt

	var sol = solvePositive(a, b, c)

	if math.Ceil(sol) == math.Floor(sol) {
		return int(sol)
	}

	return -1
}

func main() {
	fmt.Println(FindNb(4))
	fmt.Println(FindNb(1071225))
	fmt.Println(FindNb(91716553919377))
	fmt.Println(FindNb(2304422822859502501))
}
