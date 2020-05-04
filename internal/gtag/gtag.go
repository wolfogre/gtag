package gtag

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

type GenerateResult struct {
	PkgPath    string
	OutputPath string
	Content    []byte
	Errs       []error
}

func Generate(ctx context.Context, wd string) ([]GenerateResult, error) {
	cfg := &packages.Config{
		Context: ctx,
		Mode:    packages.LoadAllSyntax,
		Dir:     wd,
		Env:     os.Environ(),
	}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		return nil, fmt.Errorf("load packages: %w", err)
	}

	for _, p := range pkgs {
		for _, err := range p.Errors {
			return nil, fmt.Errorf("%v: %w", p.Name, err)
		}
	}

	generated := make([]GenerateResult, len(pkgs))
	for i, pkg := range pkgs {
		generated[i].PkgPath = pkg.PkgPath
		outDir, err := detectOutputDir(pkg.GoFiles)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	}

}

func detectOutputDir(paths []string) (string, error) {
	if len(paths) == 0 {
		return "", errors.New("no files to derive output directory from")
	}
	dir := filepath.Dir(paths[0])
	for _, p := range paths[1:] {
		if dir2 := filepath.Dir(p); dir2 != dir {
			return "", fmt.Errorf("found conflicting directories %q and %q", dir, dir2)
		}
	}
	return dir, nil
}
