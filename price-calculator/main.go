package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0.1, 0.2, 0.15, 0.07}

	for _, taxRate := range taxRates {
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
