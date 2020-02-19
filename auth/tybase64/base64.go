package tybase64

import "encoding/base64"

func B64Encode(src []byte) string {
	if len(src) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString(src)
}

func B64Decode(s string) []byte {
	if s == "" {
		return nil
	}

	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil
	}
	return data
}
