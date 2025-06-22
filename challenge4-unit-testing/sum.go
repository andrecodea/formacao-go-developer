package calculator

import "fmt"

func main() {
	x := sum(1, 2, 3)
	y := multiply(10, 20)
	z := divide(200, 10)
	w := subtract(19, 9)
	fmt.Println(x, w, y, z)
}

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total
}

func subtract(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}

func multiply(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}

	return total
}

func divide(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		total = total / v
	}
	return total
}
