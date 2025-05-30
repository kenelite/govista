package resourceloader

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

// LoadImage attempts to load an image from a given source URL or local path.
func LoadImage(src string) fyne.CanvasObject {
	if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
		return loadRemoteImage(src)
	} else {
		return loadLocalImage(src)
	}
}

// loadRemoteImage downloads the image and returns a fyne Image.
func loadRemoteImage(imgURL string) fyne.CanvasObject {
	resp, err := http.Get(imgURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return canvas.NewText("[Image not found]", nil)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return canvas.NewText("[Image load error]", nil)
	}

	r := fyne.NewStaticResource("remote", data)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillContain
	return img
}

// loadLocalImage loads image from local filesystem.
func loadLocalImage(path string) fyne.CanvasObject {
	u, err := url.Parse(path)
	if err == nil && u.Scheme != "" && u.Scheme != "file" {
		src := storage.NewURI(path)
		img := canvas.NewImageFromURI(src)
		img.FillMode = canvas.ImageFillContain
		return img
	}

	file, err := os.Open(path)
	if err != nil {
		return canvas.NewText("[Image missing]", nil)
	}
	defer file.Close()

	img := canvas.NewImageFromReader(file, path)
	img.FillMode = canvas.ImageFillContain
	return img
}
