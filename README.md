# Go GUI training with webview

Simple training app I'm using to teach myself how to create GUI app with go.
This app opens a window and load an HTML page inside.

## Running the app

```
go run main.go
```

## Building for MacOS

First, you need to package static files (html pages).

```
pkger
```

>*Note: make sure to have the pkger [CLI installed](https://github.com/markbates/pkger#cli).*

Then build the app:

```
go build -o webview.app/Contents/MacOS/webview
```

If everything ran fine, you should be able to open the app:

```
open webview.app
```

## Dependencies

- [Webview](https://github.com/webview/webview) for GUI
- [pkger](https://github.com/markbates/pkger) to package static assets