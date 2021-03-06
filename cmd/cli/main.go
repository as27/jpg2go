package main

import (
	"fmt"

	"path/filepath"

	"strings"

	"flag"

	"github.com/as27/jpg2go"
	"github.com/as27/jpg2go/jpg"
	"github.com/as27/jpg2go/nfntresizer"
	"github.com/as27/jpg2go/osfiles"
	"github.com/as27/jpg2go/resize"
)

var srcFolder = ""
var relResizeFolder = "resize"
var resizeFolder = ""

func main() {
	flag.StringVar(&srcFolder, "src", "./", "Source folder which is scanned")
	flag.StringVar(&relResizeFolder, "dst", "resize", "Subfolder inside the src where the resized pictures are saved.")
	width := flag.Int("w", 200, "Max width of the resize picture")
	height := flag.Int("h", 200, "Max height of the resize picture")
	flag.Parse()
	resizeFolder = filepath.Join(srcFolder, relResizeFolder)
	// FilePathsGetter using osfiles
	fpg := osfiles.NewFilewalk()
	fpg.Filter = fileFilter

	imager := &jpg.Imager{}
	writer := &jpg.Writer{}
	writer.Path = makePath

	c := jpg2go.FilesComponents{
		FilePathsGetter: fpg,
		Imager:          imager,
		Resizer:         &nfntresizer.Resizer{},
		Writer:          writer,
	}
	s := jpg2go.Size{
		Width:  *width,
		Height: *height,
	}
	err := resize.Files(c, srcFolder, s)
	if err != nil {
		fmt.Println(err)
	}
}

// fileFilter returns ture, when I file should be used
// is needed for the osfiles package
func fileFilter(path string) bool {
	// Skips all files of the dst folder
	// TODO: closure of resizeFolder is not clean code
	if strings.HasPrefix(path, resizeFolder) {
		return false
	}
	ext := filepath.Ext(path)
	if strings.EqualFold(ext, ".jpg") ||
		strings.EqualFold(ext, ".jpeg") {
		return true
	}
	return false
}

// makePath function is needed for the jpg package
func makePath(srcPath string) string {
	relPath, _ := filepath.Rel(
		srcFolder,
		srcPath,
	)
	out := filepath.Join(
		srcFolder,
		relResizeFolder,
		relPath,
	)
	fmt.Println("Writing: ", out)
	return out
}
