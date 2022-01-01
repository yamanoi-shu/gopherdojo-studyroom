package image_converter

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
)

type ImageConverter struct{}

func NewImageConverter() *ImageConverter {
	return &ImageConverter{}
}

func (ic *ImageConverter) ConvertImageExt(dirPath string, fromExt string, toExt string) error {
	err := walkDir(dirPath, fromExt, toExt)
	return err
}

func walkDir(dirPath string, fromExt string, toExt string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			err := walkDir(path.Join(dirPath, file.Name()), fromExt, toExt)
			if err != nil {
				return err
			}
		} else if path.Ext(file.Name()) == fromExt {
			switch toExt {
			case ".png":
				err = convertToPNG(path.Join(dirPath, file.Name()))
			case ".jpg":
				err = convertToJPG(path.Join(dirPath, file.Name()))
			}
		}
		if err != nil {
			return err
		}
	}
	return nil

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

	newFile, err := os.Create(imagePath[:len(imagePath)-len(path.Ext(imagePath))] + ".jpg")
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

	newFile, err := os.Create(imagePath[:len(imagePath)-len(path.Ext(imagePath))] + ".png")
	if err != nil {
		return err
	}

	defer newFile.Close()

	if err := png.Encode(newFile, img); err != nil {
		return err
	}

	return nil
}
