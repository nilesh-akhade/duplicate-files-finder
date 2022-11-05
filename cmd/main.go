package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/inhies/go-bytesize"
	"github.com/nilesh-akhade/duplicate-files-finder/pkg/dupe"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "Directory to check for duplicates")
	flag.Parse()
	dupeFilesFinder := dupe.New(dir, true)
	dfi, err := dupeFilesFinder.Find()
	if err != nil {
		log.Fatalf("err:%v", err)
		return
	}
	// fmt.Printf("%+v\n", dfi)
	fmt.Println("==================================")
	fmt.Printf("Total Files Scanned            : %v\n", dfi.Total)
	fmt.Printf("Files which are duplicate      : %v\n", dfi.DuplicateFilesCount)
	fmt.Printf("Total Unique Files             : %v\n", dfi.UniqueFilesCount)
	sz := bytesize.New(float64(dfi.DuplicateSize))
	fmt.Printf("Space taken by the extra copies: %v\n", sz.String())
	fmt.Println("=================================")
}
