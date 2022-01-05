package main

import (
	"flag"
	"fmt"
	"kadai1/image_converter"
	"log"
	"os"
)

var srcExt = flag.String("src", "", "src extension")
var destExt = flag.String("dest", "", "dest extension")

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) != 1 || *srcExt == "" || *destExt == "" {
		log.Fatal("Usage: go run main.go -srcExt <src extension> -destExt <dest extension> <root path>")
	}

	fmt.Println(args)

	if _, err := os.Stat(args[0]); err != nil {
		log.Fatal("Error: path does not exist")
	}

	ic, err := image_converter.NewImageConverter(*srcExt, *destExt)
	if err != nil {
		log.Fatal(err)
	}
	err = ic.ConvertImageExt(args[0])
	if err != nil {
		log.Fatal(err)
	}
}
