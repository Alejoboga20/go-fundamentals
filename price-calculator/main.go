package main

import (
	"price-calculator/cmdmanager"
	pricesPackage "price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.3}

	for _, taxRate := range taxRates {
		// _ := filemanager.NewFileManger("prices.txt", fmt.Sprintf("tax_included_prices_%.0f.json", taxRate*100))
		cmdm := cmdmanager.NewCMDManager()

		priceJob := pricesPackage.NewTaxIncludedPriceJob(taxRate, cmdm)
		priceJob.LoadData()
		priceJob.Process()
	}

}
