package main

import (
	"crypto/des"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	key := []byte{61, 22, 11, 57, 110, 89, 236, 255, 0, 99, 111, 136, 55, 4, 247, 10, 11, 45, 71, 167, 21, 157, 54, 51}

	ct := os.Args[1]
	ct = strings.TrimPrefix(ct, "zxx")

	ctb, err := hex.DecodeString(ct)
	if err != nil {
		log.Fatal(err)
	}

	cipher, err := des.NewTripleDESCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(ctb); i += cipher.BlockSize() {
		cipher.Decrypt(ctb[i:], ctb[i:])
	}

	fmt.Println(string(ctb))
}
