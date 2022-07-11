package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

// GetGolangVersion 获取Golang语言版本
func GetGolangVersion() string {
	return runtime.Version()
}

// GetCurrentDirectory 获取当前运行路径
func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// PathExists 判断指定路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetFilePathShortName 获取文件路径不带后缀的名称，比如/root/data/test.txt获取的是test
func GetFilePathShortName(filePath string) string {
	fileName := path.Base(filePath)
	ext := path.Ext(fileName)
	return strings.TrimSuffix(fileName, ext)
}

// ReadStringFromFile 从一个指定文件中读取数据并转换成字符串
func ReadStringFromFile(filePath string) (string, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(buf[:]), nil
}

// GetIndent 从一个指定字符串中获取前置空格数量
func GetIndent(str string) string {
	i := 0
	for ; i < len(str); i++ {
		c := str[i]
		if c == ' ' || c == '\t' {
			continue
		}
		break
	}
	return str[:i]
}

// GetIndentOf 获取包含某个字符串的那一行的前置空格数量
func GetIndentOf(str string, of string) string {
	line := GetTheLineContains(str, of)
	return GetIndent(line)
}

// GetTheLineContains 获取首次包含特定字符串的行
func GetTheLineContains(str string, contain string) string {
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if strings.Contains(line, contain) {
			return line
		}
	}
	return ""
}

// FillIndent 在字符串的每一行添加缩进
func FillIndent(str string, indent string) string {
	buf := &bytes.Buffer{}

	lines := strings.Split(str, "\n")
	for _, line := range lines {
		buf.WriteString(fmt.Sprintf("%v%v\n", indent, line))
	}
	return buf.String()
}

// Unindent 取消缩进
func Unindent(str string) string {
	i := 0
	for ; i < len(str); i++ {
		c := str[i]
		if c == ' ' || c == '\t' {
			continue
		}
		break
	}
	return str[i:]
}

// UnindentLines 取消前N行的缩进
func UnindentLines(str string, lineNum int) string {
	buf := &bytes.Buffer{}

	lines := strings.Split(str, "\n")
	for _, line := range lines {
		lineNum--
		if lineNum >= 0 {
			buf.WriteString(fmt.Sprintf("%v\n", Unindent(line)))
		} else {
			buf.WriteString(fmt.Sprintf("%v\n", line))
		}
	}
	return buf.String()
}

// // golang新版本的应该
// func PathExist(_path string) bool {
// 	_, err := os.Stat(_path)
// 	if err != nil && os.IsNotExist(err) {
// 		return false
// 	}
// 	return true
// }

func NowTimeMs() int64 {
	return int64(time.Now().UnixNano() / (1000 * 1000))
}

// EnsurePath 确保一条路径是存在的，如果不存在就创建
func EnsurePath(filePath string) {
	pathBase := path.Dir(filePath)
	os.MkdirAll(pathBase, os.ModePerm)
}
