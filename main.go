package main

import (
	"fmt"

	"github.com/cloudblockandbeyond/shift_cipher/decrypt"
	"github.com/cloudblockandbeyond/shift_cipher/encrypt"
)

func main() {
	encryptedString := encrypt.GetDigest("HelloWorld")
	fmt.Println(encryptedString)

	decryptedString := decrypt.GetMessage(encryptedString)
	fmt.Println(decryptedString)
}
