package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/wolfogre/gtag/internal/gtag"
)

var (
	Types = flag.String("types", "", "struct types")
	Tags  = flag.String("tags", "", "preset tags")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if *Types == "" || len(args) != 1 {
		printUsages()
		return
	}
	dir := args[0]

	types := strings.Split(*Types, ",")
	var tags []string
	if *Tags != "" {
		tags = strings.Split(*Tags, ",")
	}

	_, err := gtag.Generate(context.Background(), dir, types, tags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printUsages() {
	fmt.Println(`gtag -types A,B dir`)
	flag.PrintDefaults()
}
