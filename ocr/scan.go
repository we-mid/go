package ocr

import (
	"fmt"
	"log"
	"strings"

	"gitee.com/we-mid/go/util"
)

type scanOpts struct {
	isFromBytes bool
	bs          []byte
	filePath    string
	languages   []string
	// whitelist    string
}

func _scan(opts scanOpts) (string, error) {
	client, err := pool.Get()
	if err != nil {
		return "", err
	}
	defer pool.Put(client)

	if opts.isFromBytes {
		log.Printf("[ocr] reading from bytes: %d...\n", len(opts.bs))
		client.SetImageFromBytes(opts.bs)
	} else {
		filePath, err := util.Untildify(opts.filePath)
		if err != nil {
			return "", err
		}
		log.Printf("[ocr] reading from path: %q ...\n", filePath)
		client.SetImage(filePath)
	}
	langs := opts.languages
	if len(langs) == 0 {
		langs = []string{"eng"}
	}
	client.SetLanguage(langs...)
	log.Printf("[ocr] languages: %v\n", langs)

	// client.SetWhitelist(opts.whitelist)
	// log.Printf("[ocr] whitelist: %q\n", opts.whitelist)

	text, err := client.Text()
	if err != nil {
		if strings.Contains(err.Error(), "PixImage is not set") {
			err = fmt.Errorf("Empty content provided")
		}
		return "", err
	}
	return text, nil
}
