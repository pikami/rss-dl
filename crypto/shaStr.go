package crypto

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func ShaStr(input string) string {
	h := sha1.New()
	io.WriteString(h, input)
	return hex.EncodeToString(h.Sum(nil))
}
