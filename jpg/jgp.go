// Package jpg implements the jpg2go Imager interface by using
// the jpg package from the standart lilbrary
package jpg

import (
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

// Imager implements the Imager interface
type Imager struct {
}

// Image implements the Imager interface
func (i *Imager) Image(filePath string) (image.Image, error) {
	var img image.Image
	fi, err := os.Open(filePath)
	defer fi.Close()
	if err != nil {
		return img, err
	}
	img, err = jpeg.Decode(fi)
	return img, err
}

// FileMode for ensure the filepath
var FileMode os.FileMode = 0777

// PathFunc takes the sourcePath and returns the dstPath
type PathFunc func(string) string

// Writer implements the Writer interface
type Writer struct {
	Path PathFunc
}

// Write implements the Writer interface
func (w *Writer) Write(img image.Image, srcPath string) error {
	dst := w.Path(srcPath)
	err := os.MkdirAll(filepath.Dir(dst), FileMode)
	if err != nil {
		return err
	}
	fi, err := os.Create(dst)
	defer fi.Close()
	if err != nil {
		return err
	}
	err = jpeg.Encode(fi, img, &jpeg.Options{Quality: 100})
	return err
}
