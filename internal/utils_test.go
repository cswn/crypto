package internal_test

import (
	"cswn/crypto/internal"
	"testing"
)

type testDataHexToB64 struct {
	input  []byte
	output []byte
}

type testDataXOR struct {
	b1     []byte
	b2     []byte
	output []byte
}

var testDataMapHexToB64 = []testDataHexToB64{
	{
		input:  []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"),
		output: []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"),
	},
}

var testDataMapXOR = []testDataXOR{
	{
		b1:     []byte("1c0111001f010100061a024b53535009181c"),
		b2:     []byte("686974207468652062756c6c277320657965"),
		output: []byte("746865206b696420646f6e277420706c6179"),
	},
}

func TestHexToBase64(t *testing.T) {
	for _, data := range testDataMapHexToB64 {
		result := internal.HexToBase64(data.input)
		if string(result) != string(data.output) {
			t.Fatalf(`HexToBase64 returned %q, should have been "%v"`, result, data.output)
		}
	}
}

func TestXORCombinationTwoEqualLengthBuffers(t *testing.T) {
	for _, data := range testDataMapXOR {
		result, _ := internal.XORCombinationTwoEqualLengthBuffers(data.b1, data.b2)
		if string(result) != string(data.output) {
			t.Fatalf(`XORCombinationTwoEqualLengthBuffers returned %q, should have been "%v"`, result, data.output)
		}
	}
}
