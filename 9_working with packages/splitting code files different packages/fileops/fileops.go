package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFloat(fileName string) (floatValue float64) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("No data file found.")
		return
	} else {
		floatText := string(fileData)
		floatValue, err = strconv.ParseFloat(floatText, 64)
		if err != nil {
			fmt.Println("Data corruption detected.")
			return
		} else {
			return
		}
	}
}

func WriteFloat(fileName string, floatValue float64) {
	floatText := fmt.Sprint(floatValue)
	os.WriteFile(fileName, []byte(floatText), 0644)
}
