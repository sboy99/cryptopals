package singlebytexorcipher

import (
	"encoding/hex"
	"fmt"
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
	// Decode the hex string to bytes
	bytesData, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return 0, ""
	}

	bestScore := 0
	var bestKey byte
	var bestMessage string

	// Try each possible key (0-255)
	for key := 0; key < 256; key++ {
		decrypted := make([]byte, len(bytesData))
		for i, b := range bytesData {
			decrypted[i] = b ^ byte(key)
		}
		decryptedText := string(decrypted)
		score := scoreText(decryptedText)

		if score > bestScore {
			bestScore = score
			bestKey = byte(key)
			bestMessage = decryptedText
		}
	}

	return bestKey, bestMessage
}
