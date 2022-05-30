package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"go-boiler-clean/internal/config"
	"io"

	"golang.org/x/crypto/bcrypt"
)

type (
	Hasher interface {
		Encrypt(content string) (string, error)
		Decrypt(content string) (string, error)

		HashPassword(password string) (string, error)
		CompareHashAndPassword(hash, password string) error
	}
	hasher struct {
		key []byte
	}
)

func NewHasher() *hasher {
	return &hasher{
		key: []byte(config.Config.Server.AppKey),
	}
}

func (h *hasher) Encrypt(content string) (string, error) {
	bytesContent := []byte(content)
	block, err := aes.NewCipher(h.key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := aesGCM.Seal(nonce, nonce, bytesContent, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func (h *hasher) Decrypt(content string) (string, error) {
	enc, _ := hex.DecodeString(content)

	block, err := aes.NewCipher(h.key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func (h *hasher) HashPassword(password string) (string, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), nil
}

func (h *hasher) CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
