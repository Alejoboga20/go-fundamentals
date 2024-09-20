package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		prices[lineIndex], err = strconv.ParseFloat(line, 64)

		if err != nil {
			file.Close()
			fmt.Println("An error occurred while parsing the file")
			fmt.Println(err)
			return
		}
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println("taxRate: ", job.TaxRate)
	fmt.Println("prices: ", result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}
