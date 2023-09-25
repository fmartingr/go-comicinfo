package comicinfo

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteRead(t *testing.T) {
	tmpDir := t.TempDir()
	ci := ComicInfo{
		Title: "Test",
	}

	path := filepath.Join(tmpDir, "test.xml")

	err := Write(ci, path)
	if err != nil {
		t.Errorf("error writing ComicInfo: %v", err)
	}

	ci2, err := Open(path)
	if err != nil {
		t.Errorf("error reading ComicInfo: %v", err)
	}

	require.Equal(t, ci, ci2)
}
