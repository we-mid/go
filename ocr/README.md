# ocr

> OCR-图像识别 Go语言SDK+命令行工具

**Tesseract安装**

Tesseract Installation:
https://github.com/otiai10/gosseract#installation<br>
https://github.com/tesseract-ocr/tessdoc/blob/main/Installation.md

Troubleshootings: See [we-mid/bec-grpc#troubleshootings](https://github.com/we-mid/bec-grpc?tab=readme-ov-file#troubleshootings)

**下载简体中文等语言的最新训练数据**

```sh
# 通过gitmirror加速下载gitraw
gitrawdown() {
    if [ $# -lt 2 ]; then
        echo "缺少参数"
        return 1
    fi
    local url=$(echo $1 | \
        sed 's|https://github.com/\(.*\)/\(.*\)/blob/\(.*\)|https://raw.githubusercontent.com/\1/\2/refs/heads/\3|g' | \
        sed 's|raw.githubusercontent.com|raw.gitmirror.com|')
    echo "Downloading... $url"
    echo "=> $2"
    curl -fL $url > $2
}
gitrawdown https://raw... ~/Downloads/testdata_fast/chi_sim.traineddata
```

```sh
mkdir -p ~/Downloads/tessdata_fast
cd ~/Downloads/tessdata_fast
gitrawdown https://raw.githubusercontent.com/tesseract-ocr/tessdata_fast/refs/heads/main/chi_sim.traineddata \
    chi_sim.traineddata
gitrawdown https://raw.githubusercontent.com/tesseract-ocr/tessdata_fast/refs/heads/main/eng.traineddata \
    eng.traineddata
```

**用法：作为命令行工具**

```sh
go install gitee.com/we-mid/go/ocr/cmd/ocrscan@latest

# 如果指定为自己下载的训练数据
export TESSDATA_PREFIX=~/Downloads/tessdata_fast

# 从文件路径读取图片
ocrscan -l chi_sim,eng '~/Desktop/截屏2024-10-16 10.23.27.png'

# 从剪切板读取图片
ocrscan -l chi_sim,eng -c

>> 访问来源Top10                            自然月    4
用户类型   全部用户
指标筛选   访问人数    打开次数
任务栏 ee 3,199
手机端搜索 Se 953
Android系统 mm = 320
发现 小程序 = 190
单聊分享 p 385
PC端 ， 41
小程序功能 ) 31
收藏 | 25
群聊分享 8
公众号菜单 4
```

**用法：作为Go语言SDK**

```go
// ...
import "gitee.com/we-mid/go/ocr"

func main() {
	// ...
	var text string
	var err error

	// languages can be:
	// - []
	// - ["eng"]
	// - ["chi_sim", "eng"]
	if *isClipboard {
		text, err = ocr.ScanClipboard(languages)
	} else {
		// ...
		text, err = ocr.Scan(languages, filePath)
	}
	if err != nil {
		log.Println("[ocr] error:", err)
	}
	fmt.Println(text)
}
```

**相关资料**

Available Languages:
https://github.com/tesseract-ocr/tessdoc/blob/main/Data-Files-in-different-versions.md

Trained Data: Hans => chi_sim
https://github.com/tesseract-ocr/tessdata_best/blob/main/chi_sim.traineddata
https://github.com/tesseract-ocr/tessdata_fast/blob/main/chi_sim.traineddata

GitHub Raw Mirror:
https://raw.githubusercontent.com/tesseract-ocr/tessdata_best/refs/heads/main/chi_sim.traineddata
=>
https://raw.gitmirror.com/tesseract-ocr/tessdata_best/refs/heads/main/chi_sim.traineddata
