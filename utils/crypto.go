package utils

import (
	"blog/config"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

const (
	hashSize  = 32
	iteration = 600000
)

func MakePassword(password string) string {
	salt := config.GetSettings().SecretKey

	return hex.EncodeToString(pbkdf2.Key(
		[]byte(password),
		[]byte(salt),
		iteration,
		hashSize,
		sha256.New,
	))
}

func CheckPassword(password string, hash string) bool {
	newHash := MakePassword(password)
	return hash == newHash
}
