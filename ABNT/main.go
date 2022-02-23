package main

import (
	"NewABNT/CSV"
	"NewABNT/Functions"
	"fmt"
)

func main()  {

	authors := CSV.ReadCsvFile("NameCSV/Authors.csv")

	abnt := Functions.ABNT(authors)

	CSV.WriteCSV("Autores_ABNT",abnt)

	fmt.Println(abnt[123].TextABNTLong)

}


