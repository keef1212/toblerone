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
	"path/filepath"
	"strings"
)

const keyFile = "key.txt"

var encryptionKey string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: toblerone <command> [args]")
		fmt.Println("  keygen              generate a new encryption key")
		fmt.Println("  encrypt <inputFile.txt> <outputFile.tobl> <encryptionKey>   Encrypt a file")
		fmt.Println("  decrypt <inputFile.tobl> <outputFile.txt> <senderDecryptionKey>   Decrypt a file")
		return
	}

	command := os.Args[1]
	switch command {
	case "keygen":
		err := Keygen()
		if err != nil {
			fmt.Println("Error generating key:", err)
		}
	case "encrypt":
		if len(os.Args) < 5 {
			fmt.Println("Usage: file-encrypt encrypt <inputFile> <outputFile> <encryptionKey>")
			return
		}
		inputFile := os.Args[2]
		outputFile := os.Args[3]
		encryptionKey := os.Args[4]

		err := EncryptFile(encryptionKey, inputFile, outputFile)
		if err != nil {
			fmt.Println("Error encrypting file:", err)
		}
	case "decrypt":
		if len(os.Args) < 5 {
			fmt.Println("Usage: file-encrypt decrypt <inputFile> <outputFile> <encryptionKey>")
			return
		}
		inputFile := os.Args[2]
		outputFile := os.Args[3]
		encryptionKey := os.Args[4]

		err := DecryptFile(encryptionKey, inputFile, outputFile)
		if err != nil {
			fmt.Println("Error decrypting file:", err)
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}

// Keygen generates a new encryption key and saves it to a file.
func Keygen() error {
	fmt.Print("Enter encryption key: ")
	key, err := readPassword()
	if err != nil {
		return err
	}

	return SaveKeyToFile(key)
}

// SaveKeyToFile saves the encryption key to a file after base64 encoding.
func SaveKeyToFile(key []byte) error {
	encodedKey := base64.StdEncoding.EncodeToString(key)
	err := ioutil.WriteFile(keyFile, []byte(encodedKey), 0600)
	if err != nil {
		return err
	}
	fmt.Printf("Key saved to %s\n", keyFile)
	return nil
}

// readPassword reads a password from the user without echoing it to the terminal.
func readPassword() ([]byte, error) {
	password := make([]byte, 0)

	for {
		char := make([]byte, 1)
		_, err := os.Stdin.Read(char)
		if err != nil {
			return nil, err
		}

		// Break if Enter key is pressed
		if char[0] == '\n' {
			break
		}

		password = append(password, char[0])
	}

	return password, nil
}

// EncryptFile encrypts the contents of the input file and saves it to the output file with the .tobl extension.
func EncryptFile(encryptionKey, inputFile, outputFile string) error {
	key := deriveKey(encryptionKey)

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

	// Change the output file extension to .tobl
	outputFile = changeFileExtension(outputFile, "tobl")

	return ioutil.WriteFile(outputFile, encryptedData, 0644)
}

// DecryptFile decrypts the contents of the input file and saves it to the output file.
func DecryptFile(encryptionKey, inputFile, outputFile string) error {
	key := deriveKey(encryptionKey)

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

// deriveKey derives a 32-byte key from the provided encryption key using SHA-256.
func deriveKey(keyString string) []byte {
	hash := sha256.Sum256([]byte(keyString))
	return hash[:]
}

// changeFileExtension changes the file extension of a given file path.
func changeFileExtension(filePath, newExtension string) string {
	ext := filepath.Ext(filePath)
	newPath := strings.TrimSuffix(filePath, ext) + "." + newExtension
	return newPath
}
