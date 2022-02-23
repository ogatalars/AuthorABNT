package CSV

import (
	"NewABNT/Error"
	"encoding/csv"
	"os"
)

func ReadCsvFile(filePath string) []string{
	var data []string

	csvFile, err0 := os.Open(filePath)
	Error.CheckError(err0)

	defer csvFile.Close()

	csvLines, err1 := csv.NewReader(csvFile).ReadAll()
	Error.CheckError(err1)

	for _, line := range csvLines {
		emp := line[0]
		data = append(data, emp)
	}

	return data
}

