package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd *CMDManager) ReadLines() ([]string, error) {
	var prices []string
	fmt.Println("Please enter the prices separated by a new line")

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd *CMDManager) WriteResult(data any) (any, error) {
	fmt.Println(data)
	return nil, nil
}

func NewCMDManager() *CMDManager {
	return &CMDManager{}
}
