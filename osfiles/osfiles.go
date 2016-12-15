// Package osfiles is an implementation for the FilePathGetter and the
// Writer interface.
package osfiles

import (
	"os"
	"path/filepath"
)

// Filewalk implements the FilePathGetter interface.
type Filewalk struct {
	Filter FilterFunc
}

// FilterFunc is used to filter out not wanted filepaths. If that
// func returns false the path won't be added to the slice.
type FilterFunc func(path string) bool

// NewFilewalk creates a Filewalk and returns a pointer. By changing the
// filter some filepaths can be filtered out.
func NewFilewalk() *Filewalk {
	return &Filewalk{
		Filter: func(s string) bool {
			// No filter function by default
			return true
		},
	}
}

// Get implements the jpg2go FilePathsGetter interface
func (f *Filewalk) Get(rootPath string) []string {
	var filepaths []string
	filepath.Walk(
		rootPath,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && f.Filter(path) {
				filepaths = append(filepaths, path)
			}
			return err
		},
	)
	return filepaths
}
