package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	conf := pdfcpu.NewAESConfiguration("", "102610261026", 256)
	api.DecryptFile("test.pdf", "output.pdf", conf)
}
