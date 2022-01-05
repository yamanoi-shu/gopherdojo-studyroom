package image_converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type ImageConverter struct {
	srcExt  string
	destExt string
}

var extMap = map[string]string{
	"jpg":  "jpg",
	"jpeg": "jpg",
	"png":  "png",
}

func NewImageConverter(srcExt string, destExt string) (ic *ImageConverter, err error) {
	if _, ok := extMap[srcExt]; !ok {
		err = fmt.Errorf("Extension (%s) does not support", srcExt)
		return
	}
	if _, ok := extMap[destExt]; !ok {
		err = fmt.Errorf("Extension (%s) does not support", destExt)
		return
	}
	ic = &ImageConverter{
		srcExt:  extMap[srcExt],
		destExt: extMap[destExt],
	}
	return
}

func (ic *ImageConverter) ConvertImageExt(rootPath string) error {

	err := filepath.Walk(rootPath, func(path string, fi os.FileInfo, err error) error {

		fmt.Printf("target path: %s\n", path)

		fileExt := filepath.Ext(fi.Name())

		if !fi.IsDir() && extMap[fileExt[1:]] == ic.srcExt {

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			img, _, err := image.Decode(file)
			if err != nil {
				return err
			}

			newFile, err := os.Create(path[:len(path)-len(filepath.Ext(path))] + "." + ic.destExt)
			if err != nil {
				return err
			}
			defer newFile.Close()

			switch ic.destExt {
			case "png":
				err = png.Encode(newFile, img)
			case "jpg":
				err = jpeg.Encode(newFile, img, nil)
			}
		}
		return err
	})
	return err
}
