package browser

import (
	"io"
	"net/http"
	"strings"

	"github.com/kenelite/govista/internal/layout"
	"github.com/kenelite/govista/internal/renderer"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Browser struct {
	window  fyne.Window
	toolbar *layout.Toolbar
	content *widget.Label
	current string
	history []string
	forward []string
}

func NewBrowser(win fyne.Window) *Browser {
	b := &Browser{
		window:  win,
		content: widget.NewLabel("Welcome to GoVista!"),
	}

	tb := layout.NewToolbar(b.GoBack, b.GoForward, b.Refresh, b.NavigateTo)
	b.toolbar = tb

	scroll := container.NewScroll(b.content)
	ui := container.NewBorder(tb.Container(), nil, nil, nil, scroll)
	win.SetContent(ui)

	return b
}

func (b *Browser) NavigateTo(url string) {
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		b.content.SetText("Failed to create request: " + err.Error())
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		b.content.SetText("Failed to load page: " + err.Error())
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		b.content.SetText("Error reading content")
		return
	}
	html := string(data)

	b.history = append(b.history, b.current)
	b.current = url
	b.forward = nil
	b.window.SetContent(container.NewBorder(b.toolbar.Container(), nil, nil, nil, renderer.RenderHTML(html)))
}

func (b *Browser) GoBack() {
	if len(b.history) == 0 {
		return
	}
	b.forward = append([]string{b.current}, b.forward...)
	last := b.history[len(b.history)-1]
	b.history = b.history[:len(b.history)-1]
	b.NavigateTo(last)
}

func (b *Browser) GoForward() {
	if len(b.forward) == 0 {
		return
	}
	next := b.forward[0]
	b.forward = b.forward[1:]
	b.history = append(b.history, b.current)
	b.NavigateTo(next)
}

func (b *Browser) Refresh() {
	b.NavigateTo(b.current)
}
