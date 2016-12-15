package resize

import (
	"testing"

	"bytes"

	"github.com/as27/jpg2go"
	"github.com/as27/jpg2go/mock"
)

func TestResize(t *testing.T) {
	mock := &mock.Resizer{}
	size := jpg2go.Size{700, 700}
	buf := bytes.NewBufferString("ImageBuffer")
	Resize(mock, size, buf, buf)

}
