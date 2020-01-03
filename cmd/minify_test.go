package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestMinifyInvalidMediaType(t *testing.T) {
	rescueStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	defer func() {
		os.Stderr = rescueStderr
	}()

	got := runMinify([]string{"invalid"})
	expected := ""
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	w.Close()
	out, _ := ioutil.ReadAll(r)

	expected = fmt.Sprintf("An invalid media type was specified. %s\n", supportedMediaTypes)
	if string(out) != expected {
		t.Errorf("Expected %s, got %s", expected, out)
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
    <p>Hello world!</p>
  </body>
</html>`)

	got, err := testMinify(mediaTypeJSON, input)
	if err != nil {
		t.Error(err)
	}
	expected := `<!DOCTYPE html><html><head><title>This is a title</title></head><body><p>Hello world!</p></body></html>`
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

	return runMinify([]string{mediaType}), nil
}
