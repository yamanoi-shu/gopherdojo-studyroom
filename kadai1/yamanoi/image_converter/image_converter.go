package image_converter

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type ImageConverter struct{}

func NewImageConverter() *ImageConverter {
	return &ImageConverter{}
}

func (ic *ImageConverter) ConvertImageExt(dirPath string, fromExt string, toExt string) error {
	err := filepath.Walk(dirPath, func(path string, fi os.FileInfo, err error) error {
		if filepath.Ext(fi.Name()) == fromExt {
			switch toExt {
			case ".png":
				err = convertToPNG(path)
			case ".jpg":
				err = convertToJPG(path)
			}
		}
		return err
	})
	return err
}

func convertToJPG(imagePath string) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	newFile, err := os.Create(imagePath[:len(imagePath)-len(filepath.Ext(imagePath))] + ".jpg")
	if err != nil {
		return err
	}

	defer newFile.Close()

	if err := jpeg.Encode(newFile, img, nil); err != nil {
		return err
	}

	return nil
}

func convertToPNG(imagePath string) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	newFile, err := os.Create(imagePath[:len(imagePath)-len(filepath.Ext(imagePath))] + ".png")
	if err != nil {
		return err
	}

	defer newFile.Close()

	if err := png.Encode(newFile, img); err != nil {
		return err
	}

	return nil
}
