package main

import (
	"fmt"
	"price-calculator/cmdmanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0.1, 0.2, 0.15, 0.07}

	for _, taxRate := range taxRates {
		// fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%0f.json", taxRate*100)) // swapabble struct
		cmdm := cmdmanager.NewCmdManager()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Error processing job")
			fmt.Println(err)
		}
	}
}
