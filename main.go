package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/image/font/sfnt"
)

func main() {
	src, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	f, err := sfnt.Parse(src)
	if err != nil {
		panic(err)
	}

	nameIDs := []sfnt.NameID{
		sfnt.NameIDCopyright,
		sfnt.NameIDFamily,
		sfnt.NameIDSubfamily,
		sfnt.NameIDUniqueIdentifier,
		sfnt.NameIDFull,
		sfnt.NameIDVersion,
		sfnt.NameIDPostScript,
		sfnt.NameIDTrademark,
		sfnt.NameIDManufacturer,
		sfnt.NameIDDesigner,
		sfnt.NameIDDescription,
		sfnt.NameIDVendorURL,
		sfnt.NameIDDesignerURL,
		sfnt.NameIDLicense,
		sfnt.NameIDLicenseURL,
		sfnt.NameIDTypographicFamily,
		sfnt.NameIDTypographicSubfamily,
		sfnt.NameIDCompatibleFull,
		sfnt.NameIDSampleText,
		sfnt.NameIDPostScriptCID,
		sfnt.NameIDWWSFamily,
		sfnt.NameIDWWSSubfamily,
		sfnt.NameIDLightBackgroundPalette,
		sfnt.NameIDDarkBackgroundPalette,
		sfnt.NameIDVariationsPostScriptPrefix,
	}

	labels := []string{
		"sfnt.NameIDCopyright",
		"sfnt.NameIDFamily",
		"sfnt.NameIDSubfamily",
		"sfnt.NameIDUniqueIdentifier",
		"sfnt.NameIDFull",
		"sfnt.NameIDVersion",
		"sfnt.NameIDPostScript",
		"sfnt.NameIDTrademark",
		"sfnt.NameIDManufacturer",
		"sfnt.NameIDDesigner",
		"sfnt.NameIDDescription",
		"sfnt.NameIDVendorURL",
		"sfnt.NameIDDesignerURL",
		"sfnt.NameIDLicense",
		"sfnt.NameIDLicenseURL",
		"sfnt.NameIDTypographicFamily",
		"sfnt.NameIDTypographicSubfamily",
		"sfnt.NameIDCompatibleFull",
		"sfnt.NameIDSampleText",
		"sfnt.NameIDPostScriptCID",
		"sfnt.NameIDWWSFamily",
		"sfnt.NameIDWWSSubfamily",
		"sfnt.NameIDLightBackgroundPalette",
		"sfnt.NameIDDarkBackgroundPalette",
		"sfnt.NameIDVariationsPostScriptPrefix",
	}

	for i, nameID := range nameIDs {
		value, err := f.Name(nil, nameID)
		if err != nil {
			continue
		}

		label := labels[i]

		fmt.Printf("%s (%d): %s\n", label, nameID, value)
	}
}
