package main

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func main() {
	log.Print("Hello! This script programmed with GopherJS usage")
	js.Global.Set("OnBodyLoad", OnBodyLoad)
}

func OnBodyLoad() {

	doc := dom.GetWindow().Document()
	body := doc.GetElementsByTagName("body")[0]
	body.AppendChild(GenPage())

}
