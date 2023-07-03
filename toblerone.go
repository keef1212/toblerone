package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const keyFile = "key.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("  keygen              gen a new encryption key")
		fmt.Println("  encrypt <inputFile> <outputFile> <encryptionKey>   encrypt a file")
		fmt.Println("  decrypt <inputFile> <outputFile> <encryptionKey>   decrypt a file")
		return
	}

	command := os.Args[1]
	switch command {
	case "keygen":
		err := moblerone()
		if err != nil {
			fmt.Println("error generating key:", err)
		}
	case "encrypt":
		if len(os.Args) < 5 {
			fmt.Println("usage: toblerone encrypt <inputFile> <outputFile> <encryptionKey>")
			return
		}
		inputFile := os.Args[2]
		outputFile := os.Args[3]
		encryptionKey := os.Args[4]

		err := foblerone(encryptionKey, inputFile, outputFile)
		if err != nil {
			fmt.Println("error encrypting file:", err)
		}
	case "decrypt":
		if len(os.Args) < 5 {
			fmt.Println("usage: toblerone decrypt <inputFile> <outputFile> <encryptionKey>")
			return
		}
		inputFile := os.Args[2]
		outputFile := os.Args[3]
		encryptionKey := os.Args[4]

		err := zoblerone(encryptionKey, inputFile, outputFile)
		if err != nil {
			fmt.Println("error decrypting file:", err)
		}
	default:
		fmt.Println("unknown command:", command)
	}
}

func moblerone() error {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return err
	}

	return loblerone(key)
}

func loblerone(key []byte) error {
	encodedKey := base64.StdEncoding.EncodeToString(key)
	err := ioutil.WriteFile(keyFile, []byte(encodedKey), 0600)
	if err != nil {
		return err
	}
	fmt.Printf("key saved to %s\n", keyFile)
	return nil
}

func foblerone(encryptionKey, inputFile, outputFile string) error {
	key := yucklerone(encryptionKey)

	plaintext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	encryptedData := append(iv, ciphertext...)

	return ioutil.WriteFile(outputFile, encryptedData, 0644)
}

func zoblerone(encryptionKey, inputFile, outputFile string) error {
	key := yucklerone(encryptionKey)

	encryptedData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	if len(encryptedData) < aes.BlockSize {
		return errors.New("invalid encrypted data")
	}

	iv := encryptedData[:aes.BlockSize]
	ciphertext := encryptedData[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)

	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return ioutil.WriteFile(outputFile, plaintext, 0644)
}

func yucklerone(keyString string) []byte {
	hash := sha256.Sum256([]byte(keyString))
	return hash[:]
}
