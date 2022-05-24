package bytesutil

import "encoding/base64"

func Base64To(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Base64From(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(b)
	}
	return b
}
