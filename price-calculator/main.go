package main

import (
	"fmt"
	"price-calculator/cmdmanager"
	"price-calculator/filemanager"
	pricesPackage "price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.3}

	for _, taxRate := range taxRates {
		fm := filemanager.NewFileManger("prices.txt", fmt.Sprintf("tax_included_prices_%.0f.json", taxRate*100))
		_ = cmdmanager.NewCMDManager()

		priceJob := pricesPackage.NewTaxIncludedPriceJob(taxRate, fm)
		priceJob.LoadData()
		_, err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process the job")
			fmt.Println(err)
		}
	}

}
