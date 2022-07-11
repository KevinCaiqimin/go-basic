package main

import (
	"fmt"
	"helloserver/timer"
	"time"

	// "strings"

	"bytes"
	"crypto/aes"

	// "io"
	// "crypto/rand"
	"crypto/cipher"
	"encoding/base64"

	"caiqimin.tech/basic/encrypt"
	"caiqimin.tech/basic/heap"
	// "caiqimin.tech/basic/mathext"
	// "caiqimin.tech/basic/xlog"
	// "github.com/huangml/log"
	// import "caiqimin.tech/basic/mathext"
	// import "strings"
)

const (
	PKCS7 = iota
	ZEROS
)

func paddingZEROS(data []byte, blockSize int) []byte {
	paddingLen := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(0)}, paddingLen)
	return append(data, padText...)
}

func unpaddingZEROS(data []byte) []byte {
	zeros := 0
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == 0 {
			zeros++
			continue
		}
		break
	}
	return data[:len(data)-zeros]
}

func paddingPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func unpaddingPKCS7(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func aes_cbc_enc(data, key, iv []byte, paddingType int) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	switch paddingType {
	case PKCS7:
		data = paddingPKCS7(data, blockSize)
		break
	case ZEROS:
		data = paddingZEROS(data, blockSize)
		break
	default:
		return nil, fmt.Errorf("unsupported padding type")
	}
	dst := make([]byte, len(data))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(dst, data)

	return dst, nil
}

func aes_cbc_dec(data, key, iv []byte, paddingType int) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	dataLen := len(data)
	if dataLen < blockSize || dataLen%blockSize != 0 {
		return nil, fmt.Errorf("invalid aes data")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	var decData []byte
	switch paddingType {
	case PKCS7:
		decData = unpaddingPKCS7(data)
		break
	case ZEROS:
		decData = unpaddingZEROS(data)
		break
	default:
		return nil, fmt.Errorf("unsupported padding type")
	}

	return decData, nil
}

type TestInt struct {
	a int
	A int
}

func (i *TestInt) Compare(o heap.HeapData) int {
	other := o.(*TestInt)
	if i.a < other.a {
		return -1
	}
	if i.a > other.a {
		return 1
	}
	return 0
}

func TestChange(t heap.HeapData) {
	t.(*TestInt).A = 999
	// t.a = 999
}

func main() {
	rawData := "{\"bundle_id\":\"com.dgkp.fg.bn\",\"client_time\":10,\"create_if_not_exists\":1,\"res_p1_ver\":0,\"session_key\":\"Yjk3MzVmMTctZGM4.TInu2MikRVIBnoFEuNQt3GZFS\\/lJ7+qqNb9N2vv1vJYLj2RLSSBrIGIMmy3qCx81qGiMWDnRS4oWB7gaOG5SW2TA6VssFF6BRmf0SW3Jpgs=.uYx19MG753e+Xhe16i\\/WR4sC3j8=\",\"uid\":6443406,\"channel_uid\":\"CUX_6443406\",\"channel_id\":\"Gami_Shadowing\",\"command\":\"GuestLogin\",\"idfv\":\"1\",\"app_version\":\"1\",\"idfa\":\"1\",\"game_index\":2701,\"platform\":\"Win\",\"api_ver\":\"1\"}"
	key := "446348d3Cz05392C"
	iv := "                "
	cbc, _ := encrypt.ToAesCbc(rawData, key, iv, ZEROS)
	b64 := base64.StdEncoding.EncodeToString(cbc)
	fmt.Println(b64)
	raw, _ := encrypt.FromAesCbc(cbc, key, iv, ZEROS)
	fmt.Println(string(raw))

	fmt.Println(encrypt.ToAesCbcBase64String(rawData, key, iv, ZEROS))

	fmt.Println(encrypt.FromAesCbcBase64String(b64, key, iv, ZEROS))

	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	ivBase := "r7BXXKkLb8qrSNn05n0qiA=="

	fmt.Println(encrypt.FromAesCbcAllBase64String(encryptedData, sessionKey, ivBase, PKCS7))

	timer.NewTimer(1000, -1, 1000, func() {
		fmt.Println("AAAAAAAAA")
	})
	timer.StartTimer()

	for {
		time.Sleep(time.Duration(1) * time.Second)
	}
}
