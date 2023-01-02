package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type aesCbcCryptor struct {
	aesKey []byte
	aesIV  []byte
}

func NewAesCBCEncrypt(aesKey, aesIV []byte) (cipher.BlockMode, error) {
	if len(aesIV) != aes.BlockSize {
		return nil, fmt.Errorf("invalid iv")
	}

	c, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	return cipher.NewCBCEncrypter(c, aesIV), nil
}

func NewAesCBCDecrypt(aesKey, aesIV []byte) (cipher.BlockMode, error) {
	if len(aesIV) != aes.BlockSize {
		return nil, fmt.Errorf("invalid iv")
	}

	c, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	return cipher.NewCBCDecrypter(c, aesIV), nil
}
