package Functions

import (
	"NewABNT/Structs"
	"strings"
)

func ABNT(authors []string) []Structs.DataABNT {
	var authorsABNT []Structs.DataABNT

	for i := 0; i < len(authors); i++ {
		textL, textD, textS := abntName(authors[i])

		data := Structs.DataABNT{
			AuthorName:    authors[i],
			TextABNTLong:  textL,
			TextABNTnoDot: textD,
			TextABNTSmall: textS,
		}

		authorsABNT = append(authorsABNT, data)
	}

	return authorsABNT

}

func abntName(name string) (string, string, string) {
	var abnt string
	var textABNTSmall string
	var textABNTLong string

	var finalName string
	var sonName string
	var initialNames []string
	var initialLetters []string

	splitNames, qtdNames := splitName(name)

	lastName := splitNames[qtdNames-1]

	if juniorName(lastName) == true {
		finalName = splitNames[qtdNames-2]
		sonName = splitNames[qtdNames-1]
	} else {
		finalName = lastName
		sonName = ""
	}

	for i := 0; i < len(splitNames); i++ {
		if preposicion(splitNames[i]) == false && splitNames[i] != finalName && splitNames[i] != sonName {
			initialLetters = append(initialLetters, splitNames[i][0:1]+". ")
			initialNames = append(initialNames, splitNames[i]+" ")
		} else if splitNames[i] != finalName && splitNames[i] != sonName {
			initialNames = append(initialNames, splitNames[i]+" ")
		}
	}

	if sonName != "" {
		abnt = strings.ToUpper(finalName) + " " + strings.ToUpper(sonName) + ", "
	} else {
		abnt = strings.ToUpper(finalName) + ", "
	}

	textABNTSmall = abnt
	textABNTLong = abnt

	for i := 0; i < len(initialLetters); i++ {
		textABNTSmall += initialLetters[i]
	}

	for i := 0; i < len(initialNames); i++ {
		textABNTLong += initialNames[i]
	}

	textABNTnoDot := strings.Replace(textABNTSmall, ".", "", -1)

	return textABNTLong, textABNTnoDot, textABNTSmall
}

func splitName(name string) ([]string, int) {
	text := strings.Split(name, " ")
	namesQtd := len(text)

	return text, namesQtd
}

func juniorName(name string) bool {
	returnBool := false
	var juniorNameArray []string

	juniorNameArray = append(juniorNameArray, "filho", "filha", "neto", "neta", "junior", "júnior", "segundo", "segunda", "terceiro", "terceira")

	for i := 0; i < len(juniorNameArray); i++ {
		if strings.ToLower(name) == juniorNameArray[i] {
			returnBool = true
		}
	}

	return returnBool
}

func preposicion(name string) bool {
	returnBool := false
	var preposicionArray []string

	preposicionArray = append(preposicionArray, "do", "da", "de", "dos", "das", "e")

	for i := 0; i < len(preposicionArray); i++ {
		if strings.ToLower(name) == preposicionArray[i] {
			returnBool = true
		}
	}

	return returnBool
}
