package ocr

import (
	"fmt"
	"log"
	"strings"

	"gitee.com/we-mid/go/util"
	"github.com/otiai10/gosseract/v2"

	// "github.com/atotto/clipboard"
	"golang.design/x/clipboard"
)

const (
// Available Languages
// https://github.com/tesseract-ocr/tessdoc/blob/main/Data-Files-in-different-versions.md
// language = "eng"
// language = "hans"
// language = "chi_sim"

// whiteList = "*Xx×х+%-.0123456789"
)

var (
	pool *util.Pool[*gosseract.Client]
)

func Setup() func() {
	pool = util.NewPool[*gosseract.Client](2, func() (*gosseract.Client, error) {
		return newClient(), nil
	}, func(client *gosseract.Client) error {
		return client.Close()
	})
	return func() {
		if err := pool.Destroy(); err != nil {
			log.Println("[ocr] pool.Destroy:", err)
		}
	}
}

func Scan(languages []string, filePath string) (string, error) {
	return _scan(languages, filePath, false)
}
func ScanClipboard(languages []string) (string, error) {
	return _scan(languages, "", true)
}

func _scan(languages []string, filePath string, isClipboard bool) (string, error) {
	var err error

	// todo
	// filePath, err = handleTildy(filePath)
	filePath, err = util.Untildify(filePath)

	if err != nil {
		return "", err
	}
	log.Printf("[ocr] languages: %v\n", languages)

	client, err := pool.Get()
	if err != nil {
		return "", err
	}
	defer pool.Put(client)

	if len(languages) == 0 {
		languages = []string{"eng"}
	}
	client.Languages = languages

	if isClipboard {
		log.Println("[ocr] reading from clipboard...")
		if err := clipboard.Init(); err != nil {
			log.Println("[ocr] error:", err)
		}
		bs := clipboard.Read(clipboard.FmtImage)
		client.SetImageFromBytes(bs)
	} else {
		log.Printf("[ocr] reading from path: %q ...\n", filePath)
		client.SetImage(filePath)
	}
	text, err := client.Text()
	if err != nil {
		if strings.Contains(err.Error(), "PixImage is not set") {
			err = fmt.Errorf("Empty content provided")
		}
		return "", err
	}
	return text, nil
}

func newClient() *gosseract.Client {
	client := gosseract.NewClient()
	// client.Languages = []string{language}
	// client.SetWhitelist(whiteList)
	return client
}
