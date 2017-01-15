// Package resize uses the jpg2go interfaces to resize alle the pictures
package resize

import "github.com/as27/jpg2go"

// Files resizes all files of the FilePathsGetter
func Files(c jpg2go.FilesComponents, root string, size jpg2go.Size) error {
	filePaths := c.Get(root)
	for _, filePath := range filePaths {
		img, err := c.Image(filePath)
		if err != nil {
			return err
		}
		img, err = c.Resize(img, size)
		if err != nil {
			return err
		}
		err = c.Write(img, filePath)
		if err != nil {
			return err
		}
	}
	return nil
}
