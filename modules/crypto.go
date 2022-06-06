package modules

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/pbkdf2"
	"strings"

)

func deriveKey(passphrase string, salt []byte) []byte {

  return pbkdf2.Key([]byte(passphrase), salt, 10000, 32, sha256.New)

}

func EncryptString(passphrase, plaintext string) string {
	salt := make([]byte, 8)
	rand.Read(salt)
	key := deriveKey(passphrase, salt)
	iv := make([]byte, 12)
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte("STRING"+plaintext), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
}

// this is for Meta
func decrypt(data []byte, key []byte, iv []byte) (string, error) {
	b, err1 := aes.NewCipher(key)
	if err1 != nil {
		return "", err1
	}
	aesgcm, err2 := cipher.NewGCMWithNonceSize(b, 16)
	//NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)
	if err2 != nil {
		return "", err2
	}

	dec_dat, err3 := aesgcm.Open(nil, iv, data, nil)
	if err3 != nil {
		return "", err3
	}
	return string(dec_dat), nil
}


func DecryptString(passphrase, ciphertext string) (string, error) {
	arr := strings.Split(ciphertext, "-")
	if len(arr) < 3 {
		return "", errors.New("Invalid ciphertext")
	}
	salt, e1 := hex.DecodeString(arr[0])
	iv, e2 := hex.DecodeString(arr[1])
	data, e3 := hex.DecodeString(arr[2])
	if e1 != nil || e2 != nil || e3 != nil {
		return "", errors.New("Invalid ciphertext(2)")
	}
	key := deriveKey(passphrase, salt)
	b, e4 := aes.NewCipher(key)
	if e4 != nil {
		return "", errors.New("Invalid ciphertext(3)")
	}
	aesgcm, e5 := cipher.NewGCM(b)
	if e5 != nil {
		return "", errors.New("Invalid ciphertext(4)")
	}
	data_dec, e6 := aesgcm.Open(nil, iv, data, nil)
	if e6 != nil {
		return "", e6
	}
	result := ""
	if strings.HasPrefix(string(data_dec), "STRING") {
		result = string(data_dec)[6:]
	} else {
		return "", errors.New("Invalid ciphertext(6)")
	}
	return result, nil
}
