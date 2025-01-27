package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	//防止中文乱码的一个库
	"github.com/axgle/mahonia"
)

// 读取数据
func read() {
	contentBytes, err := os.ReadFile("./kaifang.txt")
	if err != nil {
		fmt.Println("读入失败，", err)
	}

	contentStr := string(contentBytes)
	// 逐行打印，并处理乱码
	lineStrs := strings.Split(contentStr, "\n\r")
	for _, lineStr := range lineStrs {
		newStr := ConvertEncoding(lineStr, "utf-8")
		fmt.Println(newStr)
	}
}

// 方法2：缓冲读取（如果文件比较大的情况下建议是缓冲读取）
func read2() {
	file, _ := os.Open("./kaifang.txt")
	defer file.Close()
	// 建缓冲区
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		// 转str，转utf
		enStr := string(lineBytes)
		// 如中文编码用gbk
		uftStr := ConvertEncoding(enStr, "utf-8")
		fmt.Println(uftStr)
	}
}

// 处理乱码
// 参数1：处理的数据
// 参数2：数据目前的编码
// 参数3：返回的正常数据
func ConvertEncoding(srcStr string, encoding string) (dstStr string) {
	// 创建编码处理器
	enc := mahonia.NewDecoder(encoding)
	// 编码器处理字符串为utf8的字符串
	utfStr := enc.ConvertString(srcStr)
	dstStr = utfStr
	return
}

func main() {
	read2()
}
