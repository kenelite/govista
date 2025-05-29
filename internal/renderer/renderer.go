package renderer

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// RenderHTML parses raw HTML and returns a fyne.CanvasObject to display the content.
func RenderHTML(htmlSrc string) fyne.CanvasObject {
	doc, err := html.Parse(strings.NewReader(htmlSrc))
	if err != nil {
		return widget.NewLabel("Failed to parse HTML: " + err.Error())
	}

	box := container.NewVBox()
	walkNode(doc, box)
	return container.NewScroll(box)
}

func walkNode(n *html.Node, box *fyne.Container) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "h1", "h2", "h3", "h4", "h5", "h6":
			text := extractText(n)
			txt := canvas.NewText(text, color.Black)
			txt.TextStyle = fyne.TextStyle{Bold: true}
			txt.TextSize = headerSize(n.Data)
			box.Add(txt)

		case "p":
			text := extractText(n)
			box.Add(widget.NewLabel(text))

		case "a":
			href := ""
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href = attr.Val
				}
			}
			text := extractText(n)
			if u, err := url.Parse(href); err == nil {
				box.Add(widget.NewHyperlink(text, u))
			} else {
				box.Add(widget.NewLabel(text))
			}

		case "img":
			src := ""
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					src = attr.Val
				}
			}
			if img := fetchImage(src); img != nil {
				box.Add(canvas.NewImageFromImage(img))
			}

		case "ul", "ol":
			list := container.NewVBox()
			for li := n.FirstChild; li != nil; li = li.NextSibling {
				if li.Type == html.ElementNode && li.Data == "li" {
					text := extractText(li)
					list.Add(widget.NewLabel("• " + text))
				}
			}
			box.Add(list)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkNode(c, box)
	}
}

func extractText(n *html.Node) string {
	var buf bytes.Buffer
	var f func(*html.Node)
	f = func(node *html.Node) {
		if node.Type == html.TextNode {
			buf.WriteString(node.Data)
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return strings.TrimSpace(buf.String())
}

func headerSize(tag string) float32 {
	switch tag {
	case "h1":
		return 24
	case "h2":
		return 20
	case "h3":
		return 18
	default:
		return 16
	}
}

func fetchImage(src string) image.Image {
	if !strings.HasPrefix(src, "http") {
		return nil
	}
	resp, err := http.Get(src)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		img2, _, err2 := image.Decode(bytes.NewReader(data))
		if err2 != nil {
			return nil
		}
		return img2
	}
	return img
}
