package internal

func RepeatingKeyXor(key []byte, plaintext []byte) []byte {
	var result []byte
	for i, val := range plaintext {
		xordByte := XORCipherByByte(val, key[i%len(key)])
		result = append(result, xordByte)
	}
	return EncodeHex(result)
}
