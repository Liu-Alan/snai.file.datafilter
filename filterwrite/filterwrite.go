package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

func ConvertEncoding(srcStr string, encoding string) (dstStr string) {
	// 创建编码处理器
	enc := mahonia.NewDecoder(encoding)
	// 编码器处理字符串为utf8的字符串
	utfStr := enc.ConvertString(srcStr)
	dstStr = utfStr
	return
}

func main() {
	// 1. 打开文件
	file, _ := os.Open("./kaifang.txt")
	defer file.Close()

	// 创建优质文件
	goodFile, _ := os.OpenFile("./kaifang_good.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer goodFile.Close()

	// 创建劣质文件
	badFile, _ := os.OpenFile("./kaifang_bad.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer badFile.Close()

	// 2.缓冲读取
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		// 转str，转utf
		enStr := string(lineBytes)
		lineStr := ConvertEncoding(enStr, "utf-8")

		// 3.根据行数据，取身份证
		fields := strings.Split(lineStr, ",")
		// 判断长度大于等于2，下标1的位置长度=18
		if len(fields) >= 2 && len(fields[1]) == 18 {
			goodFile.WriteString(lineStr + "\n")
			fmt.Println("Good:", lineStr)
		} else {
			badFile.WriteString(lineStr + "\n")
			fmt.Println("Bad:", lineStr)
		}
	}
}
