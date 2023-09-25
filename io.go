package comicinfo

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
)

// Write writes the ComicInfo spec to the specified path.
func Write(ci ComicInfo, path string) error {
	contents, err := xml.Marshal(ci)
	if err != nil {
		return fmt.Errorf("error marshalling ComicInfo: %w", err)
	}

	err = os.WriteFile(path, bytes.Join([][]byte{xmlHeader, contents}, []byte("")), 0644)
	if err != nil {
		return fmt.Errorf("error writing ComicInfo to file: %w", err)
	}
	return nil
}

// Read reads the ComicInfo spec from the specified path.
func Open(path string) (ComicInfo, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return ComicInfo{}, fmt.Errorf("error reading file: %w", err)
	}

	var ci ComicInfo
	err = xml.Unmarshal(f, &ci)
	if err != nil {
		return ComicInfo{}, fmt.Errorf("error unmarshalling ComicInfo: %w", err)
	}

	return ci, nil
}
