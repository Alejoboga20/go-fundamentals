package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string][]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	var lines []string
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("An error occurred while opening the file")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("An error occurred while reading the file")
		fmt.Println(err)
		return
	}

	fmt.Println(lines)
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{100, 200, 300},
	}
}
