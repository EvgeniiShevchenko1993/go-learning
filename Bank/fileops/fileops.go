package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fielName string) {
	valueText := fmt.Sprint(fielName)
	os.WriteFile(fielName, []byte(valueText), 0644)
}
