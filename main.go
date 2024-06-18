package main

import (
    "bufio"
    "fmt"
    "os"

    "github.com/skip2/go-qrcode"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter website URL: ")
    url, _ := reader.ReadString('\n')
    url = url[:len(url)-1] // Remove the newline character

    err := qrcode.WriteFile(url, qrcode.Medium, 256, "qrcode.png")
    if err != nil {
        fmt.Println("Error generating QR code:", err)
    } else {
        fmt.Println("QR code generated and saved as qrcode.png")
    }
}
