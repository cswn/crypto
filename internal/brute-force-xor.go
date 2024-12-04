package internal

import (
	"regexp"
	"strings"
)

var englishCharacterOccurances = map[string]float64{
	"A": 8.2389258,
	"B": 1.5051398,
	"C": 2.8065007,
	"D": 4.2904556,
	"E": 12.813865,
	"F": 2.2476217,
	"G": 2.0327458,
	"H": 6.1476691,
	"I": 6.1476691,
	"J": 0.1543474,
	"K": 0.7787989,
	"L": 4.0604477,
	"M": 2.4271893,
	"N": 6.8084376,
	"O": 7.5731132,
	"P": 1.9459884,
	"Q": 0.0958366,
	"R": 6.0397268,
	"S": 6.3827211,
	"T": 9.1357551,
	"U": 2.7822893,
	"V": 0.9866131,
	"W": 2.3807842,
	"X": 0.1513210,
	"Y": 1.9913847,
	"Z": 0.0746517,
}

func DecryptSingleByteXorCipher(hexString []byte) ([]byte, error) {
	bytes, err := DecodeHex(hexString)
	if err != nil {
		return nil, err
	}

	key := byte(0)
	bestScore := 0.0
	for k := range 256 {
		// get result i
		xc := XORCipher(bytes, byte(k))

		// score result i
		score := ComputeScore(xc)

		// write the results with the lowest score we have so far
		if score > bestScore {
			key = byte(k)
			bestScore = score
		}
	}

	decrypted := XORCipher(bytes, key)
	return decrypted, nil
}

func ComputeScore(buffer []byte) float64 {
	score := 0.0
	for _, b := range buffer {
		char := string(b)
		if isAlphabetic(char) {
			score += englishCharacterOccurances[strings.ToUpper(char)]
		} else {
			score -= 10.0
		}
	}

	return score
}

func isAlphabetic(str string) bool {
	var text = regexp.MustCompile("^[a-zA-Z ]$")
	return text.MatchString(str)
}
