package counting

import "fmt"

func counting(a []int, x, low, high int) int {
	leftEdge := countingLeftEdge(a, x, low, high)
	rightEdge := countingRightEdge(a, x, low, high)
	fmt.Println(leftEdge, rightEdge)
	return rightEdge - leftEdge
}

func countingLeftEdge(a []int, x, low, high int) int {
	fmt.Printf("Left: %v %v", low, high)
	if low > high {
		fmt.Println()
		return low
	}
	m := (low + high) / 2
	fmt.Printf(" %v\n", m)
	if a[m] >= x {
		return countingLeftEdge(a, x, low, m-1)
	}
	return countingLeftEdge(a, x, m+1, high)
}
func countingRightEdge(a []int, x, low, high int) int {
	fmt.Printf("Right: %v %v", low, high)
	if low > high {
		fmt.Println()
		return low
	}
	m := (low + high) / 2
	fmt.Printf(" %v\n", m)
	if a[m] > x {
		return countingRightEdge(a, x, low, m-1)
	}
	return countingRightEdge(a, x, m+1, high)
}
