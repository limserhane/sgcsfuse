
package crypto

import (
	"io"
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"

	// "github.com/limserhane/sgcsfuse/internal/logger"
)

func GetKey(password string) (key []byte) {
	key32 := sha256.Sum256([]byte(password))
	key = key32[:]

	return
}

func Encrypt(data *[]byte, key []byte) {

	var encrypted []byte

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	encrypted = aesgcm.Seal(nil, nonce, *data, nil)

	*data = append(nonce, encrypted...)

	return
}



func Decrypt(data *[]byte, n *int, key []byte) { // MUST DECRYPT IN PLACE

	var decrypted [] byte
	
	nonce := (*data)[:12]
	encrypted := (*data)[12:*n]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}


	decrypted, err = aesgcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		panic(err.Error())
	}

	// *data = decrypted
	*n = copy(*data, decrypted) // in place | *data

	return
}