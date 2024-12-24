package main

import "fmt"

type JSONDocument struct {
	info string
}

func (d JSONDocument) GetData() {
	fmt.Printf("Fetched info from document: %v\n", d.info)
}

type HTMLDocument struct {
	info string
}

func (d HTMLDocument) GetData() {
	fmt.Printf("Fetched info from document: %v\n", d.info)
}

func (d HTMLDocument) accept(v DocumentVisitor) {
	v.visitHTMLDocument(d)
}

func (d JSONDocument) accept(v DocumentVisitor) {
	v.visitJSONDocument(d)
}

type DocumentVisitor interface {
	visitHTMLDocument(d HTMLDocument)
	visitJSONDocument(d JSONDocument)
}

type DocumentExporter struct {
}

func (d DocumentExporter) visitHTMLDocument(doc HTMLDocument) {
	fmt.Println("Exporting HTML document data...")
	doc.GetData()
}

func (d DocumentExporter) visitJSONDocument(doc JSONDocument) {
	fmt.Println("Exporting JSON document data...")
	doc.GetData()
}

func main() {
	htmlDocument := HTMLDocument{info: "data from HTML document"}
	jsonDocument := JSONDocument{info: "data from JSON document"}

	exporter := DocumentExporter{}

	jsonDocument.accept(exporter)
	fmt.Println()
	htmlDocument.accept(exporter)
}
