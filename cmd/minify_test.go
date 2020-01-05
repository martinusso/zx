package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestMinifyInvalidMediaType(t *testing.T) {
	got, err := runMinify([]string{"invalid"})
	expectedErr := fmt.Sprintf("An invalid media type was specified. %s", supportedMediaTypes)
	if err.Error() != expectedErr {
		t.Errorf("Expected '%s', got '%s'", expectedErr, err)
	}
	expected := ""
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestGetMediaTypes(t *testing.T) {
	values := map[string]string{
		mediaTypeCSS:  "text/css",
		mediaTypeHTML: "text/html",
		mediaTypeJS:   "application/javascript",
		mediaTypeJSON: "application/json",
		mediaTypeSVG:  "image/svg+xml",
		mediaTypeXML:  "text/xml",
		"invalid":     ""}
	for k, v := range values {
		inputMinify = ""
		got := getMediaType([]string{k})
		if got != v {
			t.Errorf("Expected '%s', got '%s'", v, got)
		}
	}
	inputMinify = ""
}

func TestMinifyCSS(t *testing.T) {
	input := []byte(`/* comment */
body {
   overflow: hidden;
   background-color: #000000;
   background-image: url(images/bg.gif);
   background-repeat: no-repeat;
}`)

	got, err := testMinify(mediaTypeCSS, input)
	if err != nil {
		t.Error(err)
	}

	expected := `body{overflow:hidden;background-color:#000;background-image:url(images/bg.gif);background-repeat:no-repeat}`
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestMinifyHTML(t *testing.T) {
	input := []byte(`<!DOCTYPE html>
<html>
  <head>
    <title>This is a title</title>
  </head>
  <body>
  <!-- comment here -->
    <p class="first program">Hello world!</p>
  </body>
</html>`)

	got, err := testMinify(mediaTypeHTML, input)
	if err != nil {
		t.Error(err)
	}
	expected := `<!doctype html><html><head><title>This is a title</title></head><body><p class="first program">Hello world!</p></body></html>`
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestMinifyEmptyInput(t *testing.T) {
	got, err := testMinify(mediaTypeHTML, nil)
	if err.Error() != emptyInput {
		t.Errorf("Expected '%s', got '%s'", emptyInput, err.Error())
	}
	expected := ""
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func testMinify(mediaType string, input []byte) (string, error) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name()) // clean up
	if _, err := tmpfile.Write(input); err != nil {
		return "", err
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		return "", err
	}

	rescueStdin := os.Stdin
	defer func() { os.Stdin = rescueStdin }()
	os.Stdin = tmpfile

	inputMinify = ""
	return runMinify([]string{mediaType})
}
