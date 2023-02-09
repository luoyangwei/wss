package main

import "github.com/luoyangwei/wss"

func main() {
	boot := wss.New(&wss.Options{})

	boot.AddResponder(&ReportAction{})
	boot.Bind(":9080")
}
