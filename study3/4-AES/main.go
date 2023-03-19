package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	key := []byte("this is a 16 byte key")
	iv := []byte("this is a 16 byte iv")

	plaintext := []byte("hello world")

	// 加密
	ciphertext, err := encrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密结果：%s\n", base64.StdEncoding.EncodeToString(ciphertext))

	// 解密
	decrypted, err := decrypt(ciphertext, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密结果：%s\n", decrypted)
}

func encrypt(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

func decrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	return plaintext, nil
}
