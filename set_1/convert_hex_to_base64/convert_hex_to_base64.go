package converthextobase64

import (
	"encoding/base64"
	"encoding/hex"
)

func ConvertHexToBase64(hexString string) string {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	base64String := base64.StdEncoding.EncodeToString(bytes)
	return base64String
}
