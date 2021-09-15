package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cmdSearch := flag.NewFlagSet("search", flag.ExitOnError)
	pkgName := cmdSearch.String("n", "", "package name")
	isMore := cmdSearch.Bool("m", false, "more packages")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n", "gopm search [-m] -n package_name")
		cmdSearch.PrintDefaults()
		os.Exit(0)
	}

	switch os.Args[1] {
	case cmdSearch.Name():
		cmdSearch.Parse(os.Args[2:])
		url := &urlBase{"https://pkg.go.dev/search", *pkgName, 1}
		url.search(*isMore)
	}
}
