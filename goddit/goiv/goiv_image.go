// +build !imagick

package goiv

import (
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/blezek/tga"
	"github.com/nfnt/resize"
)

func init() {
	tga.RegisterFormat()
}

// decode decodes image.
func decode(filename string, width, height int) (img image.Image, err error) {
	img, err = decodeURL(filename)

	if err != nil {
		return nil, err
	}

	return scale(img, width, height)
}

// decodeFile decodes image file.
func decodeFile(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", filename, err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", filename, err)
	}

	return img, nil
}

// decodeURL decodes image from URL.
func decodeURL(url string) (image.Image, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", url, err)
	}

	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", url, err)
	}

	return img, nil
}

// downloadURL returns bytes from URL.
func downloadURL(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", url, err)
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

// scale scales image keeping aspect ratio.
func scale(img image.Image, width, height int) (image.Image, error) {
	return resize.Resize(0, uint(height), img, resize.NearestNeighbor), nil
}
