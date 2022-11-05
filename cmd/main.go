package main

import (
	"fmt"
	"log"

	"github.com/nilesh-akhade/duplicate-files-finder/pkg/dupe"
)

func main() {
	dupeFilesFinder := dupe.New("/home/nilesh/pproj/clean-code/duplicate-files-finder/input", true) // TODO: sys.argv
	dfi, err := dupeFilesFinder.Find()
	if err != nil {
		log.Fatalf("err:%v", err)
		return
	}
	fmt.Printf("%+v\n", dfi)
}
