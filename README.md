# go-comicinfo

[![Go Reference](https://pkg.go.dev/badge/github.com/fmartingr/go-comicinfo.svg)](https://pkg.go.dev/github.com/fmartingr/go-comicinfo)
[![Go Report Card](https://goreportcard.com/badge/github.com/fmartingr/go-comicinfo)](https://goreportcard.com/report/github.com/fmartingr/go-comicinfo)
[![codecov](https://codecov.io/gh/fmartingr/go-comicinfo/branch/master/graph/badge.svg?token=ZQZQZQZQZQ)](https://codecov.io/gh/fmartingr/go-comicinfo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Golang implementation of the [ComicInfo.xml specification](https://anansi-project.github.io/docs/category/comicinfo) to read and write `comicinfo.xml` files.


## Contributing

Contributions are welcome, please open an issue or a pull request.

## Versioning

This project uses [Semantic Versioning](https://semver.org/) and will follow the major and minor versions of the [ComicInfo.xml specification](https://anansi-project.github.io/docs/category/comicinfo), while using the hotfix version for any changes to the library itself.

## Usage example

```go
package main

import "github.com/fmartingr/go-comicinfo/v2"

func main() {
	ci := comicinfo.NewComicInfo()
	ci.Series = "One Piece"
	ci.Number = "3093"
	ci.Summary = "Luffy and the Straw Hat Pirates return in an all-new Epic Voyage!"
	ci.AgeRating = comicinfo.AgeRatingEveryone
	ci.LanguageISO = "en"

	for i := 1; i <= 17; i++ {
		ci.Pages.Pages = append(ci.Pages.Pages, comicinfo.ComicPageInfo{
			Image: i,
		})
	}

	_ = comicinfo.Write(ci, "volume/comicinfo.xml")
}
```

Result

```xml
<?xml version="1.0" encoding="UTF-8"?>
<ComicInfo xmlns:xsd="http://www.w3.org/2001/XMLSchema"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <Series>One Piece</Series>
    <Number>3093</Number>
    <Summary>Luffy and the Straw Hat Pirates return in an all-new Epic Voyage!</Summary>
    <LanguageISO>en</LanguageISO>
    <AgeRating>Everyone</AgeRating>
    <Pages>
        <Page Image="1"></Page>
        <Page Image="2"></Page>
        <Page Image="3"></Page>
        <Page Image="4"></Page>
        <Page Image="5"></Page>
        <Page Image="6"></Page>
        <Page Image="7"></Page>
        <Page Image="8"></Page>
        <Page Image="9"></Page>
        <Page Image="10"></Page>
        <Page Image="11"></Page>
        <Page Image="12"></Page>
        <Page Image="13"></Page>
        <Page Image="14"></Page>
        <Page Image="15"></Page>
        <Page Image="16"></Page>
        <Page Image="17"></Page>
    </Pages>
</ComicInfo>

```

## License

[MIT LICENSE](LICENSE)
