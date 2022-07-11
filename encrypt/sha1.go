package encrypt

import(
	"fmt"
	"crypto/sha1"
)

func ToSHA1(s string) string {
	h1 := sha1.New()
	h1.Write([]byte(s))
	bs := h1.Sum(nil)

	return fmt.Sprintf("%x", bs)
}