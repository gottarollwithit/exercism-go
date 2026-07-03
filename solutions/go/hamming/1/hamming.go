package hamming

import "errors"

func Distance(a, b string) (int, error) {
	runaA, runeB := []rune(a), []rune(b)
	if len(runaA) != len(runeB) {
		return 0, errors.New("lengths must be equal")
	}
	distance := 0

	for index, v := range a {
		if v != runeB[index] {
			distance++
		}
	}
	return distance, nil
}
