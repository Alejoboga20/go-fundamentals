package main

import (
	pricesPackage "price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.3}

	for _, taxRate := range taxRates {
		priceJob := pricesPackage.NewTaxIncludedPriceJob((taxRate))
		priceJob.Process()
	}

}
