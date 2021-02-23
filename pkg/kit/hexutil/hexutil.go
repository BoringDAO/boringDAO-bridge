package hexutil

import "encoding/hex"

func Encode(b []byte) string {
	hx := hex.EncodeToString(b)
	// Prefer output of "0x0" instead of "0x"
	if len(hx) == 0 {
		hx = "0"
	}

	return "0x" + hx
}

func Decode(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" {
			s = s[2:]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
		if len(s) >= 2 && s[0:2] == "0x" {
			s = s[2:]
		}
		h, _ := hex.DecodeString(s)

		return h
	}
	return nil
}
