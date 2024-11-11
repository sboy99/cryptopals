package main

import (
	singlebytexorcipher "github.com/sboy99/cryptopals/set_1/single_byte_xor_cipher"
)

func main() {
	encString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	key, message := singlebytexorcipher.SingleByteXORCipher(encString)
	println("Key:", key)         // Key: 88
	println("Message:", message) // Message: Cooking MC's like a pound of bacon
}
