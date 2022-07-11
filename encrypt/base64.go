package encrypt

import(
	"encoding/base64"
)

func ToBase64(s string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(s))
	
	return sEnc
}

func FromBase64(b64 string) (string, error) {
	bs, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}
	return string(bs), err
}