package comicinfo

var xmlHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>`)

// ComicInfo ...
type ComicInfo struct {
	Title               string    `xml:"Title,omitempty"`
	Series              string    `xml:"Series,omitempty"`
	Number              string    `xml:"Number,omitempty"`
	Count               int       `xml:"Count,omitempty"`
	Volume              int       `xml:"Volume,omitempty"`
	AlternateSeries     string    `xml:"AlternateSeries,omitempty"`
	AlternateNumber     string    `xml:"AlternateNumber,omitempty"`
	AlternateCount      int       `xml:"AlternateCount,omitempty"`
	Summary             string    `xml:"Summary,omitempty"`
	Notes               string    `xml:"Notes,omitempty"`
	Year                int       `xml:"Year,omitempty"`
	Month               int       `xml:"Month,omitempty"`
	Day                 int       `xml:"Day,omitempty"`
	Writer              string    `xml:"Writer,omitempty"`
	Penciller           string    `xml:"Penciller,omitempty"`
	Inker               string    `xml:"Inker,omitempty"`
	Colorist            string    `xml:"Colorist,omitempty"`
	Letterer            string    `xml:"Letterer,omitempty"`
	CoverArtist         string    `xml:"CoverArtist,omitempty"`
	Editor              string    `xml:"Editor,omitempty"`
	Publisher           string    `xml:"Publisher,omitempty"`
	Imprint             string    `xml:"Imprint,omitempty"`
	Genre               string    `xml:"Genre,omitempty"`
	Web                 string    `xml:"Web,omitempty"`
	PageCount           int       `xml:"PageCount,omitempty"`
	LanguageISO         string    `xml:"LanguageISO,omitempty"`
	Format              string    `xml:"Format,omitempty"`
	BlackAndWhite       YesNo     `xml:"BlackAndWhite,omitempty"`
	Manga               Manga     `xml:"Manga,omitempty"`
	Characters          string    `xml:"Characters,omitempty"`
	Teams               string    `xml:"Teams,omitempty"`
	Locations           string    `xml:"Locations,omitempty"`
	ScanInformation     string    `xml:"ScanInformation,omitempty"`
	StoryArc            string    `xml:"StoryArc,omitempty"`
	SeriesGroup         string    `xml:"SeriesGroup,omitempty"`
	AgeRating           AgeRating `xml:"AgeRating,omitempty"`
	Pages               Pages     `xml:"Pages,omitempty"`
	CommunityRating     Rating    `xml:"CommunityRating,omitempty"`
	MainCharacterOrTeam string    `xml:"MainCharacterOrTeam,omitempty"`
	Review              string    `xml:"Review,omitempty"`

	// Internal
	XmlnsXsd string `xml:"xmlns:xsd,attr"`
	XmlNsXsi string `xml:"xmlns:xsi,attr"`
	// XsiSchemaLocation string `xml:"xsi:schemaLocation,attr,omitempty"`
}

func (ci *ComicInfo) SetXMLAttributes() {
	ci.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	ci.XmlNsXsi = "http://www.w3.org/2001/XMLSchema-instance"
}

// New provides a new ComicInfo struct with the XML attributes set
func NewComicInfo() *ComicInfo {
	ci := ComicInfo{}
	ci.SetXMLAttributes()
	return &ci
}

// YesNo defines the YesNo type
type YesNo string

var (
	YesNoUnknown YesNo = "Unknown"
	Yes          YesNo = "Yes"
	No           YesNo = "No"
)

// Manga defines the Manga type
type Manga string

var (
	MangaUnknown           Manga = "Unknown"
	MangaYes               Manga = "Yes"
	MangaNo                Manga = "No"
	MangeYesAndRightToLeft Manga = "YesAndRightToLeft"
)

// Rating defines the Rating type
type Rating float64

// AgeRating defines the AgeRating type
type AgeRating string

var (
	AgeRatingUnknown          AgeRating = "Unknown"
	AgeRatingAdultsOnlyPlus18 AgeRating = "Adults Only 18+"
	AgeRatingEarlyChildhood   AgeRating = "Early Childhood"
	AgeRatingEveryone         AgeRating = "Everyone"
	AgeRatingEveryone10Plus   AgeRating = "Everyone 10+"
	AgeRatingG                AgeRating = "G"
	AgeRatingKidsToAdults     AgeRating = "Kids to Adults"
	AgeRatingM                AgeRating = "M"
	AgeRatingMAPlus15         AgeRating = "MA15+"
	AgeRatingMaturePlus17     AgeRating = "Mature 17+"
	AgeRatingPG               AgeRating = "PG"
	AgeRatingRPlus18          AgeRating = "R18+"
	AgeRatingPending          AgeRating = "Rating Pending"
	AgeRatingTeen             AgeRating = "Teen"
	AgeRatingXPlus18          AgeRating = "X18+"
)

// Pages defines the Pages type (slice of ComicPageInfo for proper XML marshalling)
type Pages struct {
	Pages []ComicPageInfo `xml:"Page,omitempty"`
}

// ComicPageInfo defines the ComicPageInfo type
type ComicPageInfo struct {
	Image       int           `xml:"Image,attr"`
	Type        ComicPageType `xml:"Type,omitempty,attr"`
	DoublePage  bool          `xml:"DoublePage,omitempty,attr"`
	ImageSize   int64         `xml:"ImageSize,omitempty,attr"`
	Key         string        `xml:"Key,omitempty,attr"`
	Bookmark    string        `xml:"Bookmark,omitempty,attr"`
	ImageWidth  int           `xml:"ImageWidth,omitempty,attr"`
	ImageHeight int           `xml:"ImageHeight,omitempty,attr"`
}

// NewComicPageInfo provides a new ComicPageInfo struct with the XML attributes set
func NewComicPageInfo() ComicPageInfo {
	return ComicPageInfo{
		Type: ComicPageTypeStory,
	}
}

// ComicPageType defines the ComicPageType type
type ComicPageType string

var (
	ComicPageTypeFrontCover    ComicPageType = "FrontCover"
	ComicPageTypeInnerCover    ComicPageType = "InnerCover"
	ComicPageTypeRoundup       ComicPageType = "Roundup"
	ComicPageTypeStory         ComicPageType = "Story"
	ComicPageTypeAdvertisement ComicPageType = "Advertisement"
	ComicPageTypeEditorial     ComicPageType = "Editorial"
	ComicPageTypeLetters       ComicPageType = "Letters"
	ComicPageTypePreview       ComicPageType = "Preview"
	ComicPageTypeBackCover     ComicPageType = "BackCover"
	ComicPageTypeOther         ComicPageType = "Other"
	ComicPageTypeDeleted       ComicPageType = "Deleted"
)
