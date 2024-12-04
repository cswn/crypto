package internal

import (
	"fmt"
)

func DecryptSingleByteXorCipher(msg []byte) ([]byte, error) {
	bytes := HexToBase64(msg)

	var buffer [256]byte
	minFq := 0.0
	var key byte
	var result []byte
	for i := 0; i < len(buffer); i++ {
		// get result i
		xc := XORCipher(bytes, buffer[i])

		// score result i
		fq := ComputeFittingQuotient(xc)

		if (minFq == 0) || (fq < minFq) {
			key = buffer[i]
			result = xc
		}
	}

	fmt.Printf("key was: %s", key)

	return result, nil
}

func ComputeFittingQuotient(input []byte) float64 {

}
