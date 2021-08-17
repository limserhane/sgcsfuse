
package crypto

import (
	// "fmt"
	// "encoding/hex"
	// "crypto/aes"
)

func Encrypt(data *[]byte, keyString string) {

	var encrypted []byte

	////////////////////////////
	// encryption
	////////////////////////////

	encrypted = make([]byte, len(*data))

	for i := 0; i < len(*data); i++ {
		encrypted[i] = (*data)[i] + 1
	}

	encrypted = append(encrypted, keyString...)

	////////////////////////////
	
	// keyBytes, _ := bcrypt.GenerateFromPassword([]byte(keyString), bcrypt.DefaultCost)
	// fmt.Println(keyBytes)

	///////////////////////////

	*data = encrypted

	return
}



func Decrypt(data *[]byte, n *int, keyString string) { // MUST DECRYPT IN PLACE

	var decrypted [] byte

	////////////////////////////
	// decryption
	////////////////////////////

	*n = *n - len(keyString)

	decrypted = make([]byte, *n)

	for i := 0; i < *n; i++ {
		decrypted[i] = (*data)[i] - 1
	}

	////////////////////////////
	
	// *data = decrypted
	copy(*data, decrypted) // in place | *data 

	return
}