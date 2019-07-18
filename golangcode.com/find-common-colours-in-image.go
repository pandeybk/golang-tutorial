package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/EdlinOrg/prominentcolor"
)

func main() {
	fmt.Println("Processing image ...")
	img, err := loadImage("/Users/bkpandey/Desktop/example.jpg")

	if err != nil {
		log.Fatal("Failed to load image", err)
	}

	colors, err := prominentcolor.Kmeans(img)

	if err != nil {
		log.Fatal("Failed to process image", err)
	}

	fmt.Println("Dominant Colors:")

	for _, color := range colors {
		fmt.Println("#", color.AsString())
	}

}

func loadImage(fileInput string) (image.Image, error) {
	f, err := os.Open(fileInput)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}
