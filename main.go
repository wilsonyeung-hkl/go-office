package main

import (
	"fmt"

	docx "clyeung.hkl/go-office/pkg/docx"
)

func main() {
	imagePath := "data/example.png"
	outputPath := "output/example.docx"

	fmt.Println("start")
	docx.ImageToDocx(imagePath, outputPath)

	fmt.Println("end")
}
