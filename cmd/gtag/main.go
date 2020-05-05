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
	types = flag.String("types", "", "struct types")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if *types == "" || len(args) != 1 {
		printUsages()
		return
	}
	dir := args[0]

	_, err := gtag.Generate(context.Background(), dir, strings.Split(*types, ","))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printUsages() {
	fmt.Println(`gtag -types A,B dir`)
	flag.PrintDefaults()
}
