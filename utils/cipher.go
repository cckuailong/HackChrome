package utils

import (
	"crypto/aes"
	"crypto/cipher"
)

func AesGCMDecrypt(crypted, key, nounce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode, _ := cipher.NewGCM(block)
	origData, err := blockMode.Open(nil, nounce, crypted, nil)
	if err != nil{
		return nil, err
	}
	return origData, nil
}
