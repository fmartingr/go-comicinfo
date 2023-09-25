package comicinfo

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const schemaURL = "https://raw.githubusercontent.com/anansi-project/comicinfo/main/schema/v2.0/ComicInfo.xsd"

func TestFullSchema(t *testing.T) {
	// Download the schema from the URL into a temporary file
	tmpDir := t.TempDir()
	schemaPath := filepath.Join(tmpDir, "schema.xsd")

	resp, err := http.Get(schemaURL)
	if err != nil {
		t.Errorf("error downloading schema: %s", err)
	}

	defer resp.Body.Close()

	out, err := os.Create(schemaPath)
	if err != nil {
		t.Errorf("error creating schema file: %s", err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		t.Errorf("error writing schema file: %s", err)
	}

	// Create a new ComicInfo struct with all the fields filled
	ci := NewComicInfo()
	ci.Title = "Test"
	ci.Series = "Test"
	ci.Number = "1"
	ci.Count = 1
	ci.Volume = 1
	ci.AlternateSeries = "Test"
	ci.AlternateNumber = "1"
	ci.AlternateCount = 1
	ci.Summary = "Test"
	ci.Notes = "Test"
	ci.Year = 2021
	ci.Month = 1
	ci.Day = 1
	ci.Writer = "Test"
	ci.Penciller = "Test"
	ci.Inker = "Test"
	ci.Colorist = "Test"
	ci.Letterer = "Test"
	ci.CoverArtist = "Test"
	ci.Editor = "Test"
	ci.Publisher = "Test"
	ci.Imprint = "Test"
	ci.Genre = "Test"
	ci.Web = "Test"
	ci.PageCount = 1
	ci.LanguageISO = "en"
	ci.Format = "Test"
	ci.BlackAndWhite = Yes
	ci.Manga = MangaYes
	ci.Characters = "Test"
	ci.Teams = "Test"
	ci.Locations = "Test"
	ci.ScanInformation = "Test"
	ci.StoryArc = "Test"
	ci.SeriesGroup = "Test"
	ci.AgeRating = AgeRatingEveryone
	ci.MainCharacterOrTeam = "Test"
	ci.Review = "Test"
	ci.Pages.Pages = append(ci.Pages.Pages, ComicPageInfo{
		Image:       1,
		Type:        ComicPageTypeEditorial,
		DoublePage:  false,
		ImageSize:   999,
		Key:         "?",
		Bookmark:    "yes",
		ImageWidth:  20,
		ImageHeight: 21,
	})
	ci.CommunityRating = 5.0
	ci.MainCharacterOrTeam = "Test"
	ci.Review = "Test"

	// Write the ComicInfo struct to a temporary file
	ciPath := filepath.Join(tmpDir, "comicinfo.xml")

	err = Write(ci, ciPath)
	if err != nil {
		t.Errorf("error writing ComicInfo: %v", err)
	}

	// Validate the schema running xmllint in a process for comicinfo.xml
	cmd := exec.Command("xmllint", "--noout", "--schema", schemaPath, ciPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("error validating ComicInfo: %s\n\nValidator error: %s", err, output)
	}

	require.NoError(t, cmd.Err, output)
}
