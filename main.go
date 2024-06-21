package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	var urlInput string
	flag.StringVar(&urlInput, "url", "", "URL to encode as QR code")
	flag.Parse()

	if urlInput == "" {
		fmt.Println("Error: Please provide a URL to encode as QR code")
		flag.PrintDefaults()
		os.Exit(1)
	}


	_, err := url.ParseRequestURI(urlInput)
	if err != nil {
		fmt.Println("Error: Invalid URL format provided")
		os.Exit(1)
	}

	qrCode, err := qrcode.Encode(urlInput, qrcode.Medium, 256)
	if err != nil {
		log.Fatal("Error generating QR code:", err)
	}


	fileName := "qrcode.png"


	currentDir, err := getCurrentDirectory()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}
	filePath := filepath.Join(currentDir, fileName)
	err = saveQRCode(filePath, qrCode)
	if err != nil {
		log.Fatal("Error saving QR code:", err)
	}

	fmt.Printf("QR code saved as %s\n", filePath)
}
func saveQRCode(filePath string, qrCode []byte) error {
	// Create a new file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(qrCode)
	if err != nil {
		return err
	}

	return nil
}
func getCurrentDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}
