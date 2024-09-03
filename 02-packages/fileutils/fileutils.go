package fileutils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, floatValue float64) {
	dataText := fmt.Sprint(floatValue)
	os.WriteFile(fileName, []byte(dataText), 0644)
}

func GetFloatFromFile(fileName string, defaultValue ...float64) (float64, error) {
	// ...float64 is a variadic parameter: it allows the function to accept 0 or more float64 values
	var defVal float64
	data, err := os.ReadFile(fileName)

	if len(defaultValue) > 0 {
		defVal = defaultValue[0]
	} else {
		defVal = 0
	}

	if err != nil {
		return defVal, errors.New("error reading from file")
	}

	dataString := string(data)
	dataFloat, err := strconv.ParseFloat(dataString, 64)

	if err != nil {
		return defVal, errors.New("error parsing from file")
	}

	return dataFloat, nil
}
