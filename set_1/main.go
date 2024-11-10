package main

import fixedxor "github.com/sboy99/cryptopals/set_1/fixed_xor"

func main() {
	str1 := "1c0111001f010100061a024b53535009181c"
	str2 := "686974207468652062756c6c277320657965"
	xorResult := fixedxor.FixedXOR(str1, str2)
	println(xorResult) // Output: 746865206b696420646f6e277420706c6179
}
