package main

import (
	"cswn/crypto/internal"
	"fmt"
)

func main() {
	// challenge 1
	// hexString := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	// base64String := internal.HexToBase64(hexString)
	//fmt.Printf("%s", base64String)

	// challenge 2
	bufferString1 := []byte("1c0111001f010100061a024b53535009181c")
	bufferString2 := []byte("686974207468652062756c6c277320657965")
	xc, _ := internal.XORCombinationTwoEqualLengthBuffers(bufferString1, bufferString2)
	fmt.Printf("%s", xc)
}
