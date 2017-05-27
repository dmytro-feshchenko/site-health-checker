package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/scrypt"
)

//EncryptPassword - generate unique hash from user password
//using AES-256 encryption
func EncryptPassword(key string, password string) (encrypted []byte, err error) {
	dk, err := scrypt.Key([]byte(password), []byte(key), 16384, 8, 1, 32)
	if err != nil {
		return nil, err
	}
	return dk, nil
}

// ComparePasswords compares hash and string password and returns the result
// of compating
func ComparePasswords(key string, hash []byte, password string) (result bool, err error) {
	dk, err := EncryptPassword(key, password)
	if err != nil {
		return false, err
	}
	return dk == hash
}

// Encrypt returns the text encrypted with key
// using AES encryption
func Encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

// Decrypt return the source text with key using AES cipher
func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
