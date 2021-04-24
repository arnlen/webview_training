package main

import (
	"fmt"
	"io/ioutil"

	"github.com/markbates/pkger"
	"github.com/webview/webview"
)

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	htmlFile, err := pkger.Open("/views/index.html")
	if err != nil {
		panic(err)
	}
	defer htmlFile.Close()

	htmlFileContent, err := ioutil.ReadAll(htmlFile)
	if err != nil {
		panic(err)
	}

	htmlPageString := string(htmlFileContent)
	fmt.Println(htmlPageString)

	navigateTarget := fmt.Sprintf("data:text/html,%s", htmlPageString)

	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(navigateTarget)
	w.Run()
}
