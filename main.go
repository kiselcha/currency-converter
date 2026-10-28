package main

import "fmt"

func main() {
	var rub float64

	fmt.Print("Введите рубли: ")
	fmt.Scan(&rub)
	var usdRate float64 = 92.0 
	var eurRate float64 = 100.0 
	dollars := convertRubToTarget(rub, usdRate)
	fmt.Printf("%.2f RUB = %.2f USD\n", rub, dollars)
	euros := convertRubToTarget(rub, eurRate)
	fmt.Printf("%.2f RUB = %.2f EUR\n", rub, euros)
	dollarsToEuros := convertUsdToEur(dollars, usdRate, eurRate)
	fmt.Printf("%.2f USD = %.2f EUR\n", dollars, dollarsToEuros)
}
func convertRubToTarget(rubSum float64, targetRate float64) float64 {
	return rubSum / targetRate
}
func convertUsdToEur(usdSum float64, usdRate float64, eurRate float64) float64 {
	return (usdSum * usdRate) / eurRate
}