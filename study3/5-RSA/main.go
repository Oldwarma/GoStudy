package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// 待签名的数据
	message := []byte("hello world")

	// 计算消息摘要
	hashed := sha256.Sum256(message)

	// 对消息摘要进行数字签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}

	// 打印签名
	fmt.Printf("签名：%x\n", signature)

	// 验证签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		panic(err)
	}
	fmt.Println("签名验证通过")
}
