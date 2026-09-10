package main

import (
	"crypto/sha256"
	"fmt"
	"encoding/hex"
)

func hash(input string) string {
	h := sha256.New()
  	h.Write([]byte(input))
 	out := hex.EncodeToString(h.Sum(nil))
  	return out
}

func main() {
	fmt.Println("Test 1: Hash a string and see if it matches the stored hash")
	coolString := "guest"
	storedStringHash := "84983c60f7daadc1cb8698621f802c0d9f9a3c3c295c810748fb048115c186ec"
	stringHash := hash(coolString)
	var hashMatch string
	if stringHash == storedStringHash {
		hashMatch = " yes"
	} else {
		hashMatch = " no"
	}
	fmt.Println("Hash equivalent to stored hash?" + hashMatch)
}