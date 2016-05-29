package dto

import "bytes"

// File TBD
type File struct {
	URL string
}

// FileList TBD
type FileList []*File

// PrettyStr TBD
func (fl FileList) PrettyStr() string {
	var ls bytes.Buffer

	for _, file := range fl {
		ls.WriteString(file.URL + "\n")
	}

	return ls.String()
}

// ToURLMap TBD
func (fl FileList) ToURLMap() map[string]*File {
	m := make(map[string]*File)
	for _, file := range fl {
		m[file.URL] = file
	}
	return m
}
