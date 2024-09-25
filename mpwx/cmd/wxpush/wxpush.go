package main

import (
	"flag"

	"gitee.com/we-mid/go/mpwx"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	defer mpwx.WG.Wait()
	flag.Parse()
	text := flag.Arg(0)
	mpwx.Default.PushToAdmin(text)
}
