package encrypt

import(
	"fmt"
	"crypto/md5"
)

func ToMD5(s string) string {
	md := md5.Sum([]byte(s))

	return fmt.Sprintf("%x", md)
}