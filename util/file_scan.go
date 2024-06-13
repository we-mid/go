package util

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFile(filename string, readLine func([]byte) error) error {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}
	defer file.Close()
	// 创建一个新的 Scanner 来读取文件
	scanner := bufio.NewScanner(file)
	// 循环读取每一行
	for scanner.Scan() {
		if err := readLine(scanner.Bytes()); err != nil {
			return err
		}
	}
	// 检查扫描过程中是否有错误发生
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner.Err: %w", err)
	}
	return nil
}
