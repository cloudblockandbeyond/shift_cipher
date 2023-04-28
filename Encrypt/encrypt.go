package encrypt

import "fmt"

func GetDigest(str string) string {
	encryptedString := ""

	for _, char := range str {
		asciiCode := int(char)
		character := fmt.Sprintf("%c", asciiCode+3)
		encryptedString += character
	}

	return encryptedString
}
