package main

import "fmt"

func weightedAverage(numbers []float64, weights []float64) float64 {
	sum := 0.0
	weightSum := 0.0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i] * weights[i]
		weightSum += weights[i]
	}
	average := sum / weightSum
	return average
}

func main() {
	numbers := []float64{5.0, 7.0, 8.5, 6.0, 9.0} // lista de números
	weights := []float64{2.0, 3.0, 4.0, 1.0, 5.0} // lista de pesos correspondentes
	average := weightedAverage(numbers, weights)
	fmt.Printf("A média ponderada é %.2f\n", average)
}
