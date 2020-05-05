package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gochore/tag/internal/gtag"
)

var (
	files = flag.String("files", "", "source files")
	types = flag.String("name", "", "struct types")
)

func main() {
	flag.Parse()
	_, err := gtag.Generate(context.Background(), strings.Split(*files, ","), strings.Split(*types, ","))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
