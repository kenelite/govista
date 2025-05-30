package main

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kenelite/govista/internal/browser"
)

func main() {
	a := app.New()
	win := a.NewWindow("GoVista")
	win.Resize(fyne.NewSize(1000, 700))

	b := browser.NewBrowser(win)

	// Optional: load default homepage
	go func() {
		b.NavigateTo("google.com")
	}()

	win.ShowAndRun()
}
