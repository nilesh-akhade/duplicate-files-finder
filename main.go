package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/inhies/go-bytesize"
)

const waitMsg = "Finding duplicate files..."

func main() {
	var dir string
	var recursive bool
	flag.StringVar(&dir, "dir", ".", "Directory to check for duplicates")
	flag.BoolVar(&recursive, "r", false, "Check duplicates recursively in all sub-directories")
	flag.Parse()

	complete := make(chan bool)
	defer close(complete)
	waitMsgEraseComplete := make(chan bool)
	defer close(waitMsgEraseComplete)
	go func() {
		fmt.Print(waitMsg)
		<-complete
		fmt.Print(strings.Repeat("\b", len(waitMsg)))
		waitMsgEraseComplete <- true
	}()

	dupeFilesFinder := New(dir, recursive)
	dfi, err := dupeFilesFinder.Find()
	if err != nil {
		complete <- true
		<-waitMsgEraseComplete
		log.Printf("err: %v", err)
		return
	}

	complete <- true
	<-waitMsgEraseComplete

	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Directory                    : %s\n", dir)
	fmt.Printf("Recursive                    : %v\n", recursive)
	fmt.Printf("Total files                  : %d\n", dfi.Total)
	fmt.Printf("Unique files                 : %d\n", dfi.UniqueFilesCount)
	fmt.Printf("Files which are duplicated   : %d\n", dfi.DuplicateFilesCount)
	sz := bytesize.New(float64(dfi.DuplicateSize))
	fmt.Printf("Space taken by the duplicates: %s\n", sz.String())
	fmt.Println(strings.Repeat("-", 50))
}
