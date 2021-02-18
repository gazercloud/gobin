package main

import (
	"flag"
	"github.com/gazercloud/gobin/gobin"
	"io/ioutil"
	"log"
)

var configPackageName string
var configOutputFile string
var configPath string

func main() {
	flag.StringVar(&configPackageName, "pkg", "resources", "Package name to use in the generated code.")
	flag.StringVar(&configOutputFile, "o", "resources.go", "Optional name of the output file to be generated.")
	flag.StringVar(&configPath, "p", "resources", "Path to files")
	flag.Parse()

	files, err := gobin.GetDir(configPath)
	if err != nil {
		log.Panic(err)
		return
	}

	code := "package " + configPackageName + "\r\n\r\n"
	for _, file := range files {
		code += gobin.GenerateCode(file, "d:/555")
		code += "\r\n"
	}

	err = ioutil.WriteFile(configOutputFile, []byte(code), 0666)
	if err != nil {
		log.Panic(err)
	}
}
