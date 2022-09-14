package main

import (
	"github.com/skip2/go-qrcode"
	"io"
	"log"
	"os"
)

func main() {
	a, err := os.Open("text-qr.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	b, err := io.ReadAll(a)
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	newQR := "new-qr.png"
	err = qrcode.WriteFile(str, qrcode.Medium, 256, newQR)

	if err != nil {
		log.Fatal(err)
	}
}
