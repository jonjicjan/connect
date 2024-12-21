package decrypt

func Nimbus(str string) string { //function name should start caps
	decryptedstr := " "
	for _, c := range str {
		asciiCode := int(c)
		char := string(asciiCode - 3)
		decryptedstr += char
	}
	return decryptedstr
}
