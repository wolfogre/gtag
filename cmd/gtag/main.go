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
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

var (
	Types   = flag.String("types", "", "struct types")
	Tags    = flag.String("tags", "", "preset tags")
	Version = flag.Bool("version", false, " show version")
)

func main() {
	flag.Parse()

	if *Version {
		fmt.Printf("gtag %s, commit %s, built at %s by %s\n", version, commit, date, builtBy)
		return
	}

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

	result, err := gtag.Generate(context.Background(), dir, types, tags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range result {
		fmt.Printf("generated %s -> %s\n", v.Source, v.Output)
	}
}

func printUsages() {
	fmt.Println(`gtag -types A,B dir`)
	flag.PrintDefaults()
}
