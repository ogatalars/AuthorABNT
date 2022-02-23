package main

import (
	"NewABNT/CSV"
	"NewABNT/Functions"
)

func main() {

	authors := CSV.ReadCsvFile("NameCSV/Authors.csv")

	abnt := Functions.ABNT(authors)

	CSV.WriteCSV("Autores_ABNT", abnt)

}
