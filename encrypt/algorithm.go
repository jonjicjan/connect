package encrypt

func Nimbus(str string) string {
	encryptedstr := " "
	for _, c := range str {
		asciiCode := int(c)
		char := string(asciiCode + 3)
		encryptedstr += char
	}
	return encryptedstr
}
