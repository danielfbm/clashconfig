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

	fmt.Println("file", file, "patchFile", patchFile)
}
