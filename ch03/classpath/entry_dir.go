package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry is one of the 4 entries
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func (entry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(entry.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, entry, err
}
func (entry *DirEntry) String() string {
	return entry.absDir
}
