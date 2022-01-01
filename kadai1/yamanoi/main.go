package main

import (
	"fmt"
	"kadai1/image_converter"
)

func main() {
	ic := image_converter.NewImageConverter()
	err := ic.ConvertImageExt("/Users/yamanoishu/ghq/github.com/yamanoi-shu/image_test", ".jpg", ".png")
	fmt.Println(err)
}
