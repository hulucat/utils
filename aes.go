package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
)

func AesDecrypt(key, iv []byte, text string) ([]byte, string, error) {
	var msgDecrypt []byte
	var id string
	deciphered, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil, "", err
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, "", err
	}
	cbc := cipher.NewCBCDecrypter(c, iv)
	cbc.CryptBlocks(deciphered, deciphered)
	decoded := PKCS7Decode(deciphered)
	buf := bytes.NewBuffer(decoded[16:20])
	var msgLen int32
	binary.Read(buf, binary.BigEndian, &msgLen)
	msgDecrypt = decoded[20 : 20+msgLen]
	id = string(decoded[20+msgLen:])
	return msgDecrypt, id, nil
}

func PKCS7Decode(text []byte) []byte {
	pad := int(text[len(text)-1])
	if pad < 1 || pad > 32 {
		pad = 0
	}
	return text[:len(text)-pad]
}

// PKCS7Encode 方法用于对需要加密的明文进行填充补位
func PKCS7Encode(text []byte) []byte {
	const BlockSize = 32
	amountToPad := BlockSize - len(text)%BlockSize
	for i := 0; i < amountToPad; i++ {
		text = append(text, byte(amountToPad))
	}
	return text
}
