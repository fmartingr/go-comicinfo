package comicinfo

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Read reads the ComicInfo spec from the specified reader
func Read(r io.Reader) (*ComicInfo, error) {
	var ci ComicInfo
	err := xml.NewDecoder(r).Decode(&ci)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling ComicInfo: %w", err)
	}

	return &ci, nil
}

// Write writes the ComicInfo spec to the specified writter
func Write(ci *ComicInfo, w io.Writer) error {
	contents, err := xml.Marshal(ci)
	if err != nil {
		return fmt.Errorf("error marshalling ComicInfo: %w", err)
	}

	_, err = w.Write(bytes.Join([][]byte{xmlHeader, contents}, []byte("")))
	if err != nil {
		return fmt.Errorf("error writing ComicInfo to file: %w", err)
	}
	return nil
}

// Read reads the ComicInfo spec from the specified path.
func Open(path string) (*ComicInfo, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var ci ComicInfo
	err = xml.Unmarshal(f, &ci)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling ComicInfo: %w", err)
	}

	return &ci, nil
}

// Save writes the ComicInfo spec to the specified path.
func Save(ci *ComicInfo, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	err = Write(ci, f)
	if err != nil {
		return fmt.Errorf("error writing ComicInfo to file: %w", err)
	}

	return nil
}
