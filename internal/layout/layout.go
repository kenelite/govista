package layout

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Toolbar holds navigation controls and URL entry for GoVista.
type Toolbar struct {
	Back    *widget.Button
	Forward *widget.Button
	Refresh *widget.Button
	URL     *widget.Entry
	Go      *widget.Button
}

// NewToolbar creates a toolbar with callbacks for each action.
func NewToolbar(
	onBack func(),
	onForward func(),
	onRefresh func(),
	onGo func(string),
) *Toolbar {
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("https://...")
	urlEntry.Resize(fyne.NewSize(500, 3))
	urlEntry.Wrapping = fyne.TextTruncate // manually enlarge

	backBtn := widget.NewButton("◀", func() { onBack() })
	forwardBtn := widget.NewButton("▶", func() { onForward() })
	refreshBtn := widget.NewButton("🏡", func() { onRefresh() })
	goBtn := widget.NewButton("Go", func() { onGo(urlEntry.Text) })

	return &Toolbar{
		Back:    backBtn,
		Forward: forwardBtn,
		Refresh: refreshBtn,
		URL:     urlEntry,
		Go:      goBtn,
	}
}

// Container returns a horizontal layout of the toolbar elements.
func (t *Toolbar) Container() fyne.CanvasObject {
	// Arrange: Back, Forward, Refresh, URL entry, Go button, spacer
	return container.NewHBox(
		t.Back,
		t.Forward,
		t.Refresh,
		t.URL,
		t.Go,
		layout.NewSpacer(),
	)
}
