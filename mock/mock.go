package mock

import (
	"image"
	"io"

	"github.com/as27/jpg2go"
)

// Resizer mocks a Resizer implementation
type Resizer struct {
	R io.Reader
	W io.Writer
	S jpg2go.Size
}

func (re *Resizer) Write(w io.Writer) error {
	re.W = w
	return nil
}

func (re *Resizer) Resize(img image.Image, s jpg2go.Size) (image.Image, error) {
	re.S = s
	return image.Image{}, nil
}

var _ jpg2go.Resizer = &Resizer{}
