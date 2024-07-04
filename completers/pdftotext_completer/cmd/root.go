package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdftotext",
	Short: "Portable Document Format (PDF) to text converter",
	Long:  "https://linux.die.net/man/1/pdftotext",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("nopgbrk", "nopgbrk", false, "don't insert a page break at the end of each page")
	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().StringS("f", "f", "", "first page to convert")
	rootCmd.Flags().BoolS("simple", "simple", false, "simple one-column page layout")
	rootCmd.Flags().BoolS("simple2", "simple2", false, "simple one-column page layout, version 2")
	rootCmd.Flags().BoolS("table", "table", false, "similar to -layout, but optimized for tables")
	rootCmd.Flags().BoolS("lineprinter", "lineprinter", false, "use strict fixed-pitch/height layout")
	rootCmd.Flags().BoolS("raw", "raw", false, "keep strings in content stream order")
	rootCmd.Flags().StringS("fixed", "fixed", "", "assume fixed-pitch (or tabular) text")
	rootCmd.Flags().StringS("linespacing", "linespacing", "", "fixed line spacing for LinePrinter mode")
	rootCmd.Flags().BoolS("clip", "clip", false, "separate clipped text")
	rootCmd.Flags().BoolS("nodiag", "nodiag", false, "discard diagonal text")
	rootCmd.Flags().StringS("enc", "enc", "", "output text encoding name")
	rootCmd.Flags().StringS("eol", "eol", "", "output end-of-line convention (unix, dos, or mac)")
	rootCmd.Flags().BoolS("layout", "layout", false, "maintain original physical layout")
	rootCmd.Flags().StringS("marginl", "marginl", "", "left page margin")
	rootCmd.Flags().StringS("l", "l", "", "last page to convert")
	rootCmd.Flags().StringS("marginr", "marginr", "", "right page margin")
	rootCmd.Flags().StringS("margint", "margint", "", "top page margin")
	rootCmd.Flags().StringS("marginb", "marginb", "", "bottom page margin")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("verbose", "verbose", false, "print per-page status information")
	rootCmd.Flags().BoolS("q", "q", false, "don't print any messages or errors")
	rootCmd.Flags().StringS("cfg", "cfg", "", "configuration file to use in place of .xpdfrc")
	rootCmd.Flags().BoolS("listencodings", "listencodings", false, "list all available output text encodings")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolS("help", "help", false, "print usage information")
	rootCmd.Flags().BoolS("bom", "bom", false, "insert a Unicode BOM at the start of the text file")
	rootCmd.Flags().Bool("help", false, "print usage information")
}
