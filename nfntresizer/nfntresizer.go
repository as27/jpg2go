// Package nfntresizer implements the Resizer by using the Resizer
// package from nfnt
package nfntresizer

import (
	"image"

	"github.com/as27/jpg2go"
	"github.com/nfnt/resize"
)

// InterpolationFunction is used, when resizing.
// There are following options:
//	NearestNeighbor
//	Bilinear
//	Bicubic
//	MitchellNetravali
//	Lanczos2
//	Lanczos3
var InterpolationFunction = resize.Lanczos3

// Resizer is the type, which is implementing the jpg2go.Resizer
// interface. In the case of the nfntresizer the type is an empty
// struct.
type Resizer struct {
}

// Resize implements the jpg2go.Resizer
func (r *Resizer) Resize(img image.Image, s jpg2go.Size) (image.Image, error) {
	rimg := resize.Thumbnail(uint(s.Width), uint(s.Height), img, InterpolationFunction)
	return rimg, nil
}
