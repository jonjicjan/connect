package main

import (
	"fmt"

	"github.com/jonjicjan/cryptit/decrypt"
	"github.com/jonjicjan/cryptit/encrypt"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)
	encryptedMessage := encrypt.Nimbus("hello")
	fmt.Println(encryptedMessage)
	fmt.Println(decrypt.Nimbus(encryptedMessage))
}
