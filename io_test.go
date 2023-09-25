package comicinfo

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveRead(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.xml")

	ci := &ComicInfo{
		Title: "Test",
	}

	err := Save(ci, path)
	if err != nil {
		t.Errorf("error writing ComicInfo: %v", err)
	}

	ci2, err := Open(path)
	if err != nil {
		t.Errorf("error reading ComicInfo: %v", err)
	}

	require.Equal(t, ci, ci2)
}

func TestWriteRead(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.xml")

	ci := &ComicInfo{
		Title: "Test",
	}

	contents, err := ci.Bytes()
	if err != nil {
		t.Errorf("error writing ComicInfo: %v", err)
	}

	if err := os.WriteFile(path, contents, 0644); err != nil {
		t.Errorf("error writing ComicInfo: %v", err)
	}

	ci2, err := Open(path)
	if err != nil {
		t.Errorf("error reading ComicInfo: %v", err)
	}

	require.Equal(t, ci, ci2)
}
