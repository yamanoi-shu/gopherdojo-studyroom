package image_converter

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type ImageConverter struct{}

func NewImageConverter() *ImageConverter {
	return &ImageConverter{}
}

func (ic *ImageConverter) ConvertJPGToPNG(dirPath string) error {
	err := walkDir(dirPath)
	return err
}

func walkDir(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			err := walkDir(path.Join(dirPath, file.Name()))
			if err != nil {
				return err
			}
		} else if path.Ext(file.Name()) == ".jpg" {
			err := convertJPGToPNG(path.Join(dirPath, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil

}

func convertJPGToPNG(jpgPath string) error {
	file, err := os.Open(jpgPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	newFile, err := os.Create(strings.TrimRight(jpgPath, ".jpg") + ".png")
	if err != nil {
		return err
	}

	defer newFile.Close()

	if err := png.Encode(newFile, img); err != nil {
		return err
	}

	return nil
}
