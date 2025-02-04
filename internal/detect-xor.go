package internal

import (
	"bufio"
	"log"
	"os"
)

func DetectSingleByteXorFromMultipleStrings(path string) []byte {
	var result []byte
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Couldn't open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bestScore := 0.0
	for scanner.Scan() {
		line := scanner.Text()
		hexString := []byte(line)
		dc, _ := DecryptSingleByteXorCipher(hexString)

		// score result i
		score := ComputeScore(dc)

		// write the results with the lowest score we have so far
		if score > bestScore {
			bestScore = score
			result = dc
		}
	}

	return result
}
