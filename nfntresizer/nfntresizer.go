// Package nfntresizer implements the Resizer by using the Resizer
// package from nfnt
package nfntresizer

import (
	"image"

	"github.com/as27/jpg2go"
	"github.com/nfnt/resize"
)

var InterpolationFunction = resize.Lanczos3

type Resizer struct {
}

func (r *Resizer) Resize(img image.Image, s jpg2go.Size) (image.Image, error) {
	rimg := resize.Thumbnail(uint(s.Width), uint(s.Height), img, InterpolationFunction)
	return rimg, nil
}
