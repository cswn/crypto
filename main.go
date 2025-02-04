package main

import (
	"cswn/crypto/internal"
	"fmt"
	"reflect"
)

func main() {
	// challenge 1
	// hexString := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	// base64String := internal.HexToBase64(hexString)
	//fmt.Printf("%s", base64String)

	// challenge 2
	// bufferString1 := []byte("1c0111001f010100061a024b53535009181c")
	// bufferString2 := []byte("686974207468652062756c6c277320657965")
	// xc, _ := internal.XORCombinationTwoEqualLengthBuffers(bufferString1, bufferString2)
	// fmt.Printf("%s", xc)

	// challenge 3
	// hexString := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	// dc, _ := internal.DecryptSingleByteXorCipher(hexString)
	// fmt.Printf("%s", dc)

	// challenge 4
	// result := internal.DetectSingleByteXorFromMultipleStrings("data/4.txt")
	// fmt.Printf("%s", result)

	// challenge 5
	key := []byte("ICE")
	plainText := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	solution := []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	result := internal.RepeatingKeyXor(key, plainText)
	fmt.Printf("%s", result)
	if reflect.DeepEqual(result, solution) {
		fmt.Println("success")
	}
}
