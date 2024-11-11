package singlebytexorcipher

import (
	"encoding/hex"
	"strings"
)

// Function to score text based on English letter frequency
func scoreText(text string) int {
	score := 0
	frequency := "etaoin shrdlu"
	for _, char := range text {
		if strings.ContainsRune(frequency, char) {
			score++
		}
	}
	return score
}

// Function to decrypt a hex string XOR'd with a single character
func SingleByteXORCipher(hexString string) (byte, string) {
	decodedBytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	bestScore := 0
	var bestKey byte
	var bestMessage string

	for key := 0; key < 256; key++ {
		decryptedBytes := make([]byte, len(decodedBytes))
		for i, b := range decodedBytes {
			decryptedBytes[i] = b ^ byte(key)
		}

		message := string(decryptedBytes)
		score := scoreText(message)
		if score > bestScore {
			bestScore = score
			bestKey = byte(key)
			bestMessage = message
		}

	}

	return bestKey, bestMessage
}
