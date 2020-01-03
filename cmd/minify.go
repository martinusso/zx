package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/martinusso/zx/internal/clipboard"
	"github.com/spf13/cobra"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

const (
	mediaTypeCSS  = "css"
	mediaTypeHTML = "html"
	mediaTypeJS   = "js"
	mediaTypeJSON = "json"
	mediaTypeSVG  = "svg"
	mediaTypeXML  = "xml"

	supportedMediaTypes = `Supported media types: css, html, js, json, svg, xml`
)

var (
	inputMinify string

	minifyCmd = &cobra.Command{
		Use:   "minify",
		Short: "Minify HTML, CSS, JS, JSON, XML and SVG",
		Long: `minify removes whitespace, strips comments, combines files, and optimizes/shortens a few common programming patterns.

` + supportedMediaTypes + `

To confirm the input, press Ctrl+] ENTER`,
		Run: func(cmd *cobra.Command, args []string) {
			s := runMinify(args)
			clipboard.Write(s)
			fmt.Println(s)
		},
	}
)

func init() {
	minifyCmd.Flags().StringVarP(&inputMinify, "type", "t", "", supportedMediaTypes)
	rootCmd.AddCommand(minifyCmd)
}

func runMinify(args []string) string {
	mediaType := getMediaType(args)
	if mediaType == "" {
		err := errors.New(fmt.Sprintf("An invalid media type was specified. %s", supportedMediaTypes))
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	fmt.Println("Type (or paste) here:")
	lines := getMinifyInput()
	if len(lines) == 0 {
		return ""
	}

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.Add("text/html", &html.Minifier{
		KeepConditionalComments: true,
		KeepDefaultAttrVals:     true,
		KeepDocumentTags:        true,
		KeepEndTags:             true,
	})
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	s, err := m.String(mediaType, strings.Join(lines, ""))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	return s
}

func getMediaType(args []string) string {
	if inputMinify == "" {
		if len(args) > 0 {
			inputMinify = args[0]
		}
	}
	switch inputMinify {
	case mediaTypeCSS:
		return "text/css"
	case mediaTypeHTML:
		return "text/html"
	case mediaTypeJS:
		return "application/javascript"
	case mediaTypeJSON:
		return "application/json"
	case mediaTypeSVG:
		return "image/svg+xml"
	case mediaTypeXML:
		return "text/xml"
	}
	return ""
}

func getMinifyInput() (lines []string) {
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 1 {
			// Group Separator (GS ^]): ctrl-]
			if line[0] == '\x1D' {
				break
			}
		}
		lines = append(lines, line)
	}
	if err := scn.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return
}
