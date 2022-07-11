package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
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

func aesCbcEnc(data, key, iv []byte, paddingType int) ([]byte, error) {
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

func aesCbcDec(data, key, iv []byte, paddingType int) ([]byte, error) {
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

// convert a string with specified key, iv to a bytes array encrypted by aes cbc,
// mode can be ZEROS or PKCS7
func ToAesCbc(rawData, key string, iv string, mode int) ([]byte, error) {
	return aesCbcEnc([]byte(rawData), []byte(key), []byte(iv), mode)
}

// ToAesCbcBytes convert a string with specified key, iv to a bytes array encrypted by aes cbc,
// mode can be ZEROS or PKCS7
func ToAesCbcBytes(rawData []byte, key string, iv string, mode int) ([]byte, error) {
	return aesCbcEnc(rawData, []byte(key), []byte(iv), mode)
}

func FromAesCbc(cbcBytes []byte, key string, iv string, mode int) ([]byte, error) {
	return aesCbcDec(cbcBytes, []byte(key), []byte(iv), mode)
}

func ToAesCbcBase64String(rawData, key string, iv string, mode int) (string, error) {
	cbc, err := aesCbcEnc([]byte(rawData), []byte(key), []byte(iv), mode)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cbc), nil
}

func FromAesCbcBase64String(encryptedData string, key string, iv string, mode int) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	rawBytes, err := aesCbcDec([]byte(data), []byte(key), []byte(iv), mode)
	if err != nil {
		return "", err
	}
	return string(rawBytes), nil
}

func FromAesCbcAllBase64String(encryptedData string, key string, iv string, mode int) (string, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	rawBytes, err := aesCbcDec(dataBytes, keyBytes, ivBytes, mode)
	if err != nil {
		return "", err
	}
	return string(rawBytes), nil
}
