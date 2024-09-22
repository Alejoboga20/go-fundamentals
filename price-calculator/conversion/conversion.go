package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for stringIndex, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, errors.New("an error occurred while converting the strings to floats")
		}

		floats[stringIndex] = floatValue
	}

	return floats, nil
}
