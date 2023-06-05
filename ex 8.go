package main

import "fmt"

func triangleArea(base float64, height float64) float64 {
	area := 0.5 * base * height
	return area
}

func main() {
	var base, height float64
	fmt.Print("Digite a base do triângulo: ")
	fmt.Scan(&base)
	fmt.Print("Digite a altura do triângulo: ")
	fmt.Scan(&height)
	area := triangleArea(base, height)
	fmt.Printf("A área do triângulo é %.2f\n", area)
}
