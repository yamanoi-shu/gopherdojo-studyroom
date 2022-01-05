package main

import (
	"flag"
	"fmt"
	"kadai1/image_converter"
	"log"
)

var srcExt = flag.String("src", "", "src extension")
var destExt = flag.String("dest", "", "dest extension")

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) != 1 || *srcExt == "" || *destExt == "" {
		log.Fatal("Usage: go run main.go -src <src extension> -dest <dest extension> <root path>")
	}

	fmt.Println(args)

	ic := image_converter.NewImageConverter()
	err := ic.ConvertImageExt(args[0], *srcExt, *destExt)
	fmt.Println(err)
}
