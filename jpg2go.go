package jpg2go

import "image"

// Size describes the size of an image in pixel
type Size struct {
	Width  int
	Height int
}

// FilePathsGetter takes a filepath and returns filenpaths in a slice
type FilePathsGetter interface {
	Get(root string) []string
}

// Imager takes a filePath and return an Image
type Imager interface {
	Image(filePath string) (image.Image, error)
}

// Resizer reads resizes and writes an image.
type Resizer interface {
	Resize(img image.Image, s Size) (image.Image, error)
}

// Writer takes an Image and the srcPath and writes
type Writer interface {
	Write(img image.Image, srcPath string) error
}

// FilesComponents are all the interfaces wich are needed for
// resizing files
type FilesComponents struct {
	FilePathsGetter
	Imager
	Resizer
	Writer
}
