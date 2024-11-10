package main

import converthextobase64 "github.com/sboy99/cryptopals/set_1/convert_hex_to_base64"

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64String := converthextobase64.ConvertHexToBase64(hexString)
	println(base64String) // Output: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
}
