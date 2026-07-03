package bottlesong

import (
	"fmt"
)

var uppercaseNumberMap = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}
var lowercaseNumberMap = map[int]string{
	0: "no",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

const (
	verse1 = "%s green bottle%s hanging on the wall,"
	verse2 = "And if one green bottle should accidentally fall,"
	verse3 = "There'll be %s green bottle%s hanging on the wall."
)

func Recite(startBottles, takeDown int) []string {
	var result []string

	for i := startBottles; i > startBottles-takeDown; i-- {
		verse1Final := fmt.Sprintf(verse1, uppercaseNumberMap[i], getPluralSuffix(i))
		verse3Final := fmt.Sprintf(verse3, lowercaseNumberMap[i-1], getPluralSuffix(i-1))
		result = append(result, verse1Final, verse1Final, verse2, verse3Final)
		if i-1 > startBottles-takeDown {
			result = append(result, "")
		}
	}
	return result
}

func getPluralSuffix(number int) string {
	if number != 1 {
		return "s"
	}
	return ""
}
