// Package andotp implements functions to encrypt/decrypt andOTP files.
package andotp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

const (
	ivLen         int = 12
	keyLen        int = 32
	iterationLen  int = 4
	saltLen       int = 12
	maxIterations int = 160000
	minIterations int = 140000
)

// Encrypt encrypts plaintext with password according to andotp encryption standard.
// It returns encrypted byte array and any error encountered.
func Encrypt(plaintext []byte, password string) ([]byte, error) {

	var finalCipher []byte
	iter := make([]byte, iterationLen)
	iv := make([]byte, ivLen)
	salt := make([]byte, saltLen)

	iterations := rand.Intn(maxIterations-minIterations) + minIterations
	binary.BigEndian.PutUint32(iter, uint32(iterations))

	_, err := rand.Read(iv)
	if err != nil {
		return nil, formatError(err.Error())
	}

	_, err = rand.Read(salt)
	if err != nil {
		return nil, formatError(err.Error())
	}

	secretKey := pbkdf2.Key([]byte(password), salt, iterations, keyLen, sha1.New)

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, formatError(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, formatError(err.Error())
	}

	cipherText := aesgcm.Seal(nil, iv, plaintext, nil)

	finalCipher = append(finalCipher, iter...)
	finalCipher = append(finalCipher, salt...)
	finalCipher = append(finalCipher, iv...)
	finalCipher = append(finalCipher, cipherText...)

	return finalCipher, nil

}

// Decrypt decrypts encryptedtext using password.
// It returns decrypted byte array and any error encountered.
func Decrypt(encryptedtext []byte, password string) ([]byte, error) {

	iterations := encryptedtext[:iterationLen]
	salt := encryptedtext[iterationLen : iterationLen+saltLen]
	iv := encryptedtext[iterationLen+saltLen : iterationLen+saltLen+ivLen]
	cipherText := encryptedtext[iterationLen+saltLen+ivLen:]
	iter := int(binary.BigEndian.Uint32(iterations))
	secretKey := pbkdf2.Key([]byte(password), salt, iter, keyLen, sha1.New)

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, formatError(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, formatError(err.Error())
	}

	plaintextbytes, err := aesgcm.Open(nil, iv, cipherText, nil)
	if err != nil {
		return nil, formatError(err.Error())
	}

	return plaintextbytes, nil
}

func formatError(e string) error {
	return fmt.Errorf("error: %s", e)
}
