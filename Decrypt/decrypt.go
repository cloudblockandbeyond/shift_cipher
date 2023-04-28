package decrypt

import "fmt"

func GetMessage(str string) string {
	decryptedString := ""

	for _, char := range str {
		asciiCode := int(char)
		character := fmt.Sprintf("%c", asciiCode-3)
		decryptedString += character
	}

	return decryptedString
}
