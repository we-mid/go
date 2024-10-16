package ocr

import (
	"log"

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

var pool *util.Pool[*gosseract.Client]

func Setup(poolSize int) func() {
	pool = util.NewPool[*gosseract.Client](poolSize, func() (*gosseract.Client, error) {
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
func newClient() *gosseract.Client {
	client := gosseract.NewClient()
	// client.Languages = []string{language}
	// client.SetWhitelist(whiteList)
	return client
}

func Scan(languages []string, filePath string) (string, error) {
	return _scan(scanOpts{false, nil, filePath, languages})
}
func ScanBytes(languages []string, bs []byte) (string, error) {
	return _scan(scanOpts{true, bs, "", languages})
}
func ScanClipboard(languages []string) (string, error) {
	log.Println("[ocr] reading from clipboard...")
	if err := clipboard.Init(); err != nil {
		return "", err
	}
	bs := clipboard.Read(clipboard.FmtImage)
	return ScanBytes(languages, bs)
}
