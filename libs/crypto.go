package libs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

//DecryptWechatAppletUser 解密
func DecryptWechatAppletUser(encryptedData string, session_key string, iv string) ([]byte, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(encryptedData)
	key, _ := base64.StdEncoding.DecodeString(session_key)
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
		return nil, err
	}
	iv_b, _ := base64.StdEncoding.DecodeString(iv)
	blockModel := cipher.NewCBCDecrypter(block, iv_b)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

//PKCS7UnPadding PKSC#7解包
func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

//Sha1Signature 使用 sha1( rawData + sessionkey ) 得到字符串，用于校验用户信息
func Sha1Signature(rawData, sessionkey string) string {
	s := sha1.New()
	io.WriteString(s, rawData+sessionkey)
	return fmt.Sprintf("%x", s.Sum(nil))
}
