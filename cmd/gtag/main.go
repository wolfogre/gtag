package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/gochore/tag/internal/gtag"
)

var (
	file = flag.String("file", "", "source file")
	name = flag.String("name", "", "struct type name")
)

func main() {
	flag.Parse()
	_, err := gtag.Generate(context.Background(), *file, *name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
