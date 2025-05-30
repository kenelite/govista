# Govista

**Govista**  is a lightweight graphical web browser built with Fyne in Go. It supports loading and rendering local or remote HTML files with basic CSS and image support.



## Features

* Render HTML from remote URLs or local files

* Basic CSS support (inline` <style> `and `<link rel="stylesheet">)`

* Loads remote and local images

* Simple toolbar with address bar and navigation

* Chrome-like User-Agent


## Getting Started

### Prerequisites

* Go 1.20 or higher (tested with 1.24.3)

* Fyne v2.5.5

### Install dependencies

```bash
go mod tidy
```

### Run the browser

```bash
go run main.go
```

### Build the app

```bash
go build -o govista
```

## Usage

* Launch govista

* Enter a URL (https://example.com) or local HTML file path (/Users/you/test.html) in the address bar

* Press Enter or click "Go"

## Project Structure
```

govista/
├── main.go
├── internal/
│   ├── browser/        # Handles main UI and navigation
│   ├── layout/         # Manages toolbar and layout containers
│   ├── renderer/       # HTML parsing and rendering
│   ├── cssparser/      # Basic CSS parser
│   └── resourceloader/ # Loads image resources
└── go.mod
```

## Roadmap



## License

MIT License

Made with ❤️ in Go by [@kenelite](https://github.com/kenelite)