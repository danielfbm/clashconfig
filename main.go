package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var file string
	var patchFile string
	// flag.StringVar(&file, `config`, ``, `file to be watched`)
	flag.StringVar(&patchFile, `patch`, ``, `file with patch configuration`)

	if len(os.Args) > 1 {
		fmt.Println("args", os.Args)
		file = os.Args[len(os.Args)-1]
	}

	flag.Parse()

	// validation
	if !isFile(file) {
		os.Exit(1)
		return
	}
	if !isFile(patchFile) {
		os.Exit(1)
		return
	}

	fmt.Println("got:", "configuration file", file, "patchFile", patchFile)
}

func isFile(file string) (ok bool) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println("file", file, " stat returned error:", err)
		return
	}
	if fileInfo.IsDir() {
		fmt.Println(file, "is a directory")
		return
	}
	ok = true
	return
}
