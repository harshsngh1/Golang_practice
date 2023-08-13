package main

import (
	"fmt"

	"github.com/harshsngh1/cryptit/decrypt"
	"github.com/harshsngh1/cryptit/encrypt"
)

func main() {
	str := "KodeKloud"
	encryptstr := encrypt.Nimbus(str)
	fmt.Println("encrypted string", encryptstr)
	fmt.Println("decrypted string", decrypt.Nimbus(encryptstr))
}
