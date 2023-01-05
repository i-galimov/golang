package main

import (
	"fmt"
	"log"

	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	password := hexutil.Encode(privateKeyBytes)[2:]
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	hashSum := GenHash(password)
	encHashSum := Encrypt(hashSum)
	fmt.Println(encHashSum)
}

const keyAES = "keyAES!"

func GenHash(password string) string {
	hash := sha256.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Encrypt(hash string) string {
	key := []byte(keyAES)
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	out := make([]byte, len(hash))

	block.Encrypt(out, []byte(hash))

	return hex.EncodeToString(out)
}

func Decrypt(cipherText string) (string, error) {
	text, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	key := []byte(keyAES)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	pt := make([]byte, len(text))

	block.Decrypt(pt, text)

	s := string(pt[:])

	return s, nil
}
