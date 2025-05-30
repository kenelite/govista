package renderer

import (
	"bytes"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/kenelite/govista/internal/cssparser"
	"github.com/kenelite/govista/internal/resourceloader"
)

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderHTML(htmlStr string, basePath string) fyne.CanvasObject {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return widget.NewLabel("Failed to parse HTML")
	}

	css := extractCSS(doc, basePath)
	rules := cssparser.ParseCSS(css)

	body := findNode(doc, "body")
	if body == nil {
		return widget.NewLabel("No <body> tag found")
	}

	return r.renderNode(body, basePath, rules)
}

func extractCSS(n *html.Node, basePath string) string {
	var css string
	var walker func(*html.Node)
	walker = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "style" && node.FirstChild != nil {
			css += node.FirstChild.Data
		}
		if node.Type == html.ElementNode && node.Data == "link" {
			isCSS := false
			href := ""
			for _, attr := range node.Attr {
				if attr.Key == "rel" && attr.Val == "stylesheet" {
					isCSS = true
				}
				if attr.Key == "href" {
					href = attr.Val
				}
			}
			if isCSS && href != "" {
				if strings.HasPrefix(href, "http") {
					resp, err := http.Get(href)
					if err == nil {
						defer resp.Body.Close()
						buf := new(bytes.Buffer)
						buf.ReadFrom(resp.Body)
						css += buf.String()
					}
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			walker(c)
		}
	}
	walker(n)
	return css
}

func (r *Renderer) renderNode(n *html.Node, basePath string, rules cssparser.RuleSet) fyne.CanvasObject {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			return widget.NewLabel(text)
		}
		return nil
	}
	if n.Type == html.ElementNode {
		switch n.Data {
		case "p", "div", "span", "h1", "h2", "h3":
			content := container.NewVBox()
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				obj := r.renderNode(c, basePath, rules)
				if obj != nil {
					content.Add(obj)
				}
			}
			return content
		case "img":
			src := ""
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					src = attr.Val
					break
				}
			}
			if src != "" {
				return resourceloader.LoadImage(src)
			}
		}
	}
	return nil
}

func findNode(n *html.Node, tag string) *html.Node {
	if n.Type == html.ElementNode && n.Data == tag {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if res := findNode(c, tag); res != nil {
			return res
		}
	}
	return nil
}
