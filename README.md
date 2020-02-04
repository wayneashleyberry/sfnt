[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/sfnt)](https://goreportcard.com/report/github.com/wayneashleyberry/sfnt)

### Installation

```sh
go get github.com/wayneashleyberry/sfnt
```

### Usage

```sh
sfnt PaperCuteBold.otf | prettyjson
```

```sh
{
  "sfnt.NameIDCopyright": "Copyright (c) 2012 by Fanny Coulez, Julien Saurin. All rights reserved.",
  "sfnt.NameIDDescription": "Copyright (c) 2012 by Fanny Coulez, Julien Saurin. All rights reserved.",
  "sfnt.NameIDDesigner": "Fanny Coulez, Julien Saurin",
  "sfnt.NameIDDesignerURL": "https://www.facebook.com/LaGoupilParis",
  "sfnt.NameIDFamily": "PaperCute",
  "sfnt.NameIDFull": "PaperCute Bold",
  "sfnt.NameIDManufacturer": "Fanny Coulez, Julien Saurin",
  "sfnt.NameIDPostScript": "PaperCuteBold",
  "sfnt.NameIDSubfamily": "Bold",
  "sfnt.NameIDTrademark": "PaperCute Bold is a trademark of Fanny Coulez, Julien Saurin.",
  "sfnt.NameIDTypographicFamily": "PaperCute",
  "sfnt.NameIDTypographicSubfamily": "Bold",
  "sfnt.NameIDUniqueIdentifier": "FannyCoulez,JulienSaurin: PaperCute Bold: 2012",
  "sfnt.NameIDVendorURL": "https://www.myfonts.com/foundry/La_Goupil/",
  "sfnt.NameIDVersion": "4.000"
}
```
