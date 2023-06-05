package main

import (
	"fmt"
	"math"
)

func compoundInterest(principal float64, rate float64, time float64) float64 {
	amount := principal * math.Pow(1+rate, time)
	return amount
}

func main() {
	var principal, rate, time float64
	fmt.Print("Digite o valor principal: ")
	fmt.Scan(&principal)
	fmt.Print("Digite a taxa de juros anual em decimal (por exemplo, 0,05 para 5%): ")
	fmt.Scan(&rate)
	fmt.Print("Digite o período de tempo em anos: ")
	fmt.Scan(&time)
	amount := compoundInterest(principal, rate, time)
	fmt.Printf("O montante após %g anos é de R$ %.2f\n", time, amount)
}
