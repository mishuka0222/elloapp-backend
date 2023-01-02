package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAesCBCEncrypt(t *testing.T) {
	key := []byte("12345678901234561234567890123456")
	iv := []byte("1234567890123456")

	MyString := "This is my string and I want to protect it with encryption test."

	fmt.Printf("We start with a plain text: %s \n", MyString)
	myStringByte := []byte(MyString)
	fmt.Println(len(MyString))
	encryptor, _ := NewAesCBCEncrypt(key, iv)
	encryptor.CryptBlocks(myStringByte, myStringByte)

	fmt.Printf("We encrypted the string this way: %s \n", hex.EncodeToString(myStringByte))

	decryptor, _ := NewAesCBCDecrypt(key, iv)
	decryptor.CryptBlocks(myStringByte, myStringByte)
	fmt.Printf("Than we have the plain text again: %s \n", string(myStringByte))
}
