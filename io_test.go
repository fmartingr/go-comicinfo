package comicinfo

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveOpen(t *testing.T) {
	tmpDir := t.TempDir()
	ci := &ComicInfo{
		Title: "Test",
	}

	path := filepath.Join(tmpDir, "comicinfo.xml")

	err := Save(ci, path)
	if err != nil {
		t.Fatalf("error writing ComicInfo: %v", err)
	}

	ci2, err := Open(path)
	if err != nil {
		t.Fatalf("error reading ComicInfo: %v", err)
	}

	require.Equal(t, ci, ci2)
}

func TestWriteOpen(t *testing.T) {
	ci := &ComicInfo{
		Title: "Test",
	}

	var buf bytes.Buffer

	err := Write(ci, &buf)
	if err != nil {
		t.Fatalf("error writing ComicInfo: %v", err)
	}

	ci2, err := Read(&buf)
	if err != nil {
		t.Fatalf("error reading ComicInfo: %v", err)
	}

	require.Equal(t, ci, ci2)
}
