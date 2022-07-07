package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func If(condition bool, trueVal, falseVal any) any {
	if condition {
		return trueVal
	}
	return falseVal
}

func GetFileDirectoryToCaller(opts ...int) (directory string, ok bool) {
	var filename string
	directory = ""
	skip := 1
	if opts != nil {
		skip = opts[0]
	}
	if _, filename, _, ok = runtime.Caller(skip); ok {
		directory = path.Dir(filename)
	}
	return
}

func GetRunPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(currentPath, "/goApp")
}

//TrimLeftChars удаляет в строке символы начиная с левой части
func TrimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}
