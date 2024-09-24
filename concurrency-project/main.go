package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChannels[index] = make(chan bool)
		errorChannels[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChannels[index], errorChannels[index])
	}

	for index, _ := range taxRates {
		select {
		case err := <-errorChannels[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChannels[index]:
			fmt.Println("Done")
		}
	}
}
