package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"gitee.com/we-mid/go/ocr"
)

func main() {
	var exitCode int
	defer func() {
		log.Printf("[ocr] teardown completed. exitCode: %d\n", exitCode)
		os.Exit(exitCode)
	}()
	// Note: It's weird that exitCode is not correct if writen inline or even using args in a closure
	// defer func(c int) {
	// 	log.Printf("[ocr] teardown completed. exitCode: %d\n", c)
	// }(exitCode)
	// defer log.Printf("[ocr] teardown completed. exitCode: %d\n", exitCode)

	// 优雅关闭：信号处理，监听中断信号（如Ctrl+C）
	// 创建一个用于接收信号的通道
	sigCh := make(chan os.Signal, 1)
	// 使用 signal.Notify 注册我们关心的信号
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	teardown := ocr.Setup(1) // poolSize
	defer func() {
		// time for teardown to get ready, otherwise will cause errors like:
		// Error in pixGetWidth: pix not defined
		// Error in pixGetHeight: pix not defined
		// SIGSEGV: segmentation violation
		// PC=0x100b0c50b m=7 sigcode=1
		// signal arrived during cgo execution
		time.Sleep(1 * time.Second)
		teardown()
	}()

	doneCh := make(chan struct{}, 1)
	errorCh := make(chan error, 1)
	go mainLogic(doneCh, errorCh)

	// 阻塞等待信号
	select {
	case sig := <-sigCh:
		// 将 os.Signal 断言为 syscall.Signal 并获取其数值
		if s, ok := sig.(syscall.Signal); ok {
			log.Printf("[ocr] received signal: %q (0x%x)\n", sig, int(s))
			exitCode = 128 + int(s)
		} else {
			log.Printf("[ocr] unknown signal: %q\n", sig)
			exitCode = 128 + int(syscall.SIGINT)
		}
	case err := <-errorCh:
		log.Println("[ocr] error:", err)
		exitCode = 1
	case <-doneCh:
		exitCode = 0
	}
}

func mainLogic(doneCh chan<- struct{}, errorCh chan<- error) {
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
		errorCh <- err
		return
	}
	fmt.Println() // padding
	fmt.Println(text)
	fmt.Println() // padding
	doneCh <- struct{}{}
}
