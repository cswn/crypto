package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func DecryptSingleByteXorCipher(msg []byte) ([]byte, error) {
	bytes := HexToBase64(msg)

	// brute force
	var buffer [256]byte
	for i := 0; i < len(buffer); i++ {
		BruteForceSingleByteXorCipher(bytes, buffer[i])
	}

	// score results
	mostValid, err := ScoreBruteForceResults()
	if err != nil {
		fmt.Printf("error while scoring brute force results: %s", err)
		return nil, err
	}

	return mostValid, nil
}

func BruteForceSingleByteXorCipher(bytes []byte, keyByte byte) {
	xc := make([]byte, len(bytes))
	for i := 0; i < len(bytes); i++ {
		xc[i] = XORCipher(bytes[i], keyByte)
	}

	// write all results to a txt file
	f, err := os.Create("results.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(string(xc))
	if err2 != nil {
		log.Fatal(err2)
	}
}

func ScoreBruteForceResults() ([]byte, error) {
	// go through text file
	lines, err := readLinesOfResult("result.txt")
	if err != nil {
		return nil, err
	}
	lines = []byte(lines)

	// score each result and return highest scored one
}

func readLinesOfResult(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = scanner.Bytes()
	}
	return lines, scanner.Err()
}
