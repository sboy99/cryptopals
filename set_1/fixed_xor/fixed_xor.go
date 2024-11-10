package fixedxor

import "encoding/hex"

func FixedXOR(hexString1, hexString2 string) string {
	bytes1, err := hex.DecodeString(hexString1)
	if err != nil {
		panic(err)
	}
	bytes2, err := hex.DecodeString(hexString2)
	if err != nil {
		panic(err)
	}
	if len(bytes1) != len(bytes2) {
		panic("The length of the two byte arrays must be equal")
	}
	bytes := make([]byte, len(bytes1))
	for i := range bytes1 {
		bytes[i] = bytes1[i] ^ bytes2[i]
	}
	return hex.EncodeToString(bytes)
}
