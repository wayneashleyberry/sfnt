package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/image/font/sfnt"
)

func main() {
	cmd := &cobra.Command{
		Use:   "sfnt [filename]",
		Short: "Display truetype font metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(args[0])
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(filename string) error {
	src, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return err
	}

	f, err := sfnt.Parse(src)
	if err != nil {
		return err
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

	out := map[string]string{}

	for i, nameID := range nameIDs {
		value, err := f.Name(nil, nameID)
		if err != nil {
			continue
		}

		label := labels[i]

		out[label] = value
	}

	b, err := json.Marshal(out)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}
