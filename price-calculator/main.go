package main

import (
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0.1, 0.2, 0.15, 0.07}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
