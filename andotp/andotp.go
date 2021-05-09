package andotp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

const (
	IV_LEN         int = 12
	KEY_LEN        int = 32
	ITERATION_LEN  int = 4
	SALT_LEN       int = 12
	MAX_ITERATIONS int = 160000
	MIN_ITERATIONS int = 140000
)

func Encrypt(plaintext []byte, password string) ([]byte, error) {

	var finalCipher []byte
	iter := make([]byte, ITERATION_LEN)
	iv := make([]byte, IV_LEN)
	salt := make([]byte, SALT_LEN)

	iterations := rand.Intn(MAX_ITERATIONS-MIN_ITERATIONS) + MIN_ITERATIONS
	binary.BigEndian.PutUint32(iter, uint32(iterations))

	_, err := rand.Read(iv)
	if err != nil {
		return nil, FormatError(err.Error())
	}

	_, err = rand.Read(salt)
	if err != nil {
		return nil, FormatError(err.Error())
	}

	secretKey := pbkdf2.Key([]byte(password), salt, iterations, KEY_LEN, sha1.New)

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, FormatError(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, FormatError(err.Error())
	}

	cipherText := aesgcm.Seal(nil, iv, plaintext, nil)

	finalCipher = append(finalCipher, iter...)
	finalCipher = append(finalCipher, salt...)
	finalCipher = append(finalCipher, iv...)
	finalCipher = append(finalCipher, cipherText...)

	return finalCipher, nil

}

func Decrypt(encryptedtext []byte, password string) ([]byte, error) {

	iterations := encryptedtext[:ITERATION_LEN]
	salt := encryptedtext[ITERATION_LEN : ITERATION_LEN+SALT_LEN]
	iv := encryptedtext[ITERATION_LEN+SALT_LEN : ITERATION_LEN+SALT_LEN+IV_LEN]
	cipherText := encryptedtext[ITERATION_LEN+SALT_LEN+IV_LEN:]
	iter := int(binary.BigEndian.Uint32(iterations))
	secretKey := pbkdf2.Key([]byte(password), salt, iter, KEY_LEN, sha1.New)

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, FormatError(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, FormatError(err.Error())
	}

	plaintextbytes, err := aesgcm.Open(nil, iv, cipherText, nil)
	if err != nil {
		return nil, FormatError(err.Error())
	}

	return plaintextbytes, nil
}

func FormatError(e string) error {
	return fmt.Errorf("error: %s", e)
}

func ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}
