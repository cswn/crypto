package internal

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func HexToBase64(input []byte) []byte {
	dh, err := DecodeHex(input)
	if err != nil {
		fmt.Printf("error while decoding hex: %s", err)
		return nil
	}

	return EncodeBase64(dh)
}

func DecodeHex(src []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

func EncodeHex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	return dst
}

func EncodeBase64(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

func XORCombinationTwoEqualLengthBuffers(b1 []byte, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, errors.New("the inputs have mismatched lengths")
	}

	db1, err := DecodeHex(b1)
	if err != nil {
		fmt.Printf("error while decoding hex: %s", err)
		return nil, err
	}
	b1 = db1

	db2, err := DecodeHex(b2)
	if err != nil {
		fmt.Printf("error while decoding hex: %s", err)
		return nil, err
	}
	b2 = db2

	xc := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		xc[i] = b1[i] ^ b2[i]
	}
	return EncodeHex(xc), nil
}
