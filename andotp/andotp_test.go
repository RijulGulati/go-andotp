package andotp_test

import (
	"testing"

	"github.com/grijul/go-andotp/andotp"
)

const (
	jsonstr  string = "[{\"secret\":\"SOMEAWESOMESECRET\",\"issuer\":\"SOMEAWESOMEISSUER\",\"label\":\"TESTLABEL\",\"digits\":6,\"type\":\"TOTP\",\"algorithm\":\"SHA1\",\"thumbnail\":\"Default\",\"last_used\":1000000000000,\"used_frequency\":0,\"period\":30,\"tags\":[]}]"
	password string = "testpass"
)

func TestEncryptDecrypt(t *testing.T) {

	encryptedtext, err := andotp.Encrypt([]byte(jsonstr), password)
	if err != nil {
		t.Error(err)
	}

	decryptedtext, err := andotp.Decrypt(encryptedtext, password)
	if err != nil {
		t.Error(err)
	}

	if string(decryptedtext) != jsonstr {
		t.Error("Encryption/Decryption failed. Text mismatch")
	}

	// With wrong password

	encryptedtext = nil
	decryptedtext = nil
	err = nil

	encryptedtext, err = andotp.Encrypt([]byte(jsonstr), password)
	if err != nil {
		t.Error(err)
	}

	_, err = andotp.Decrypt(encryptedtext, "someotherpass")
	if err != nil {
		if err.Error() != "error: cipher: message authentication failed" {
			t.Error(err)
		}
	}
}
