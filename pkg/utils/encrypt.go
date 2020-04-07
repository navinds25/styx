package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

// AESEncryptCBC encrypts a message using AES CBC encryption.
func AESEncryptCBC(key, message []byte) (string, error) {
	if len(message)%aes.BlockSize != 0 {
		message = pad(message)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(message))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], message)
	encodedMsg := base64.URLEncoding.EncodeToString(ciphertext)
	return encodedMsg, err
}

// AESDecryptCBC decrypts a AES CBC encrypted message
func AESDecryptCBC(key []byte, encodedCiphertxt string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(encodedCiphertxt)
	if err != nil {
		return "", err
	}
	var block cipher.Block
	if block, err = aes.NewCipher(key); err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	plaintext := []byte{}
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(plaintext, ciphertext)
	value, err := unpad(plaintext)
	if err != nil {
		return "", err
	}
	return string(value), nil
}
