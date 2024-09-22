package main

import (
	"fmt"
	"price-calculator/filemanager"
	pricesPackage "price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.3}

	for _, taxRate := range taxRates {
		fm := filemanager.NewFileManger("prices.txt", fmt.Sprintf("tax_included_prices_%.0f.json", taxRate*100))
		priceJob := pricesPackage.NewTaxIncludedPriceJob(taxRate, *fm)
		priceJob.LoadData()
		priceJob.Process()
	}

}
