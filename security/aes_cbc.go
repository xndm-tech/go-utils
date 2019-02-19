package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"strings"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize //填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesCbcEncryptHex(encodeStr, key, iv string) string { //编码为大写的hex字符串
	encodeBytes := []byte(encodeStr) //根据key 生成密文
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)
	return strings.ToUpper(hex.EncodeToString(crypted))
}

func AesCbcDecryptHex(decodeStr, key, iv string) string { //先解码hex字符串
	decodeBytes, err := hex.DecodeString(decodeStr)
	if err != nil {
		return ""
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))
	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return string(origData)
}
