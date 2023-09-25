package comicinfo

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Write writes the ComicInfo spec to the specified path.
func Save(ci *ComicInfo, path string) error {
	contents, err := ci.Bytes()
	if err != nil {
		return fmt.Errorf("error marshalling ComicInfo: %w", err)
	}

	err = os.WriteFile(path, contents, 0644)
	if err != nil {
		return fmt.Errorf("error writing ComicInfo to file: %w", err)
	}
	return nil
}

func Write(ci *ComicInfo, w io.Writer) error {
	contents, err := ci.Bytes()
	if err != nil {
		return fmt.Errorf("error marshalling ComicInfo: %w", err)
	}

	_, err = w.Write(contents)
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

	var ci *ComicInfo
	err = xml.Unmarshal(f, &ci)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling ComicInfo: %w", err)
	}

	return ci, nil
}
