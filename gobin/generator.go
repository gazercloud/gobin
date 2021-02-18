package gobin

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GenerateResourceName(filePath string) string {
	filePath = strings.TrimPrefix(filePath, "/")
	filePath = strings.ReplaceAll(filePath, "/", "_")
	filePath = strings.ReplaceAll(filePath, ".", "_")
	return filePath
}

func GenerateCode(filePath string, rootPath string) string {
	result := ""

	resName := GenerateResourceName(strings.TrimPrefix(filePath, rootPath))

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}

	result = "var R_" + resName + " = []byte {"

	for i, b := range fileContent {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprint(b)
	}
	result += "}\r\n"

	return result
}
