package main

import (
	"fmt"

	"github.com/jonjicjan/cryptit.git/decrypt"
	"github.com/jonjicjan/cryptit.git/encrypt"
)

func main() {
	encryptpackage := encrypt.Nimbus("hello")
	fmt.Println(encryptpackage)
	fmt.Println(decrypt.Nimbus(encryptpackage))
}
