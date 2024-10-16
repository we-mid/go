package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"gitee.com/we-mid/go/ocr"
)

func main() {
	teardown := ocr.Setup()
	defer teardown()

	isClipboard := flag.Bool("c", false, "Whether to read from clipboard")
	langStr := flag.String("l", "", "Languages joined by `,`")
	flag.Parse()

	var languages []string
	if len(*langStr) > 0 {
		languages = strings.Split(*langStr, ",")
	}

	var text string
	var err error
	if *isClipboard {
		text, err = ocr.ScanClipboard(languages)
	} else {
		filePath := flag.Arg(0)
		text, err = ocr.Scan(languages, filePath)
	}
	if err != nil {
		log.Println("[ocr] error:", err)
	}
	fmt.Println(text)
}
