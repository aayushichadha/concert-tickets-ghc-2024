// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// mkstdlib generates the zstdlib.go file, containing the Go standard
// library API symbols. It's baked into the binary to avoid scanning
// GOPATH in the common case.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

func mustOpen(name string) io.Reader {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func api(base string) string {
	return filepath.Join(runtime.GOROOT(), "api", base)
}

var sym = regexp.MustCompile(`^pkg (\S+).*?, (?:var|func|type|const) ([A-Z]\w*)`)

func main() {
	var buf bytes.Buffer
	outf := func(format string, args ...interface{}) {
		fmt.Fprintf(&buf, format, args...)
	}
	outf(`// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

`)
	outf("// Code generated by mkstdlib.go. DO NOT EDIT.\n\n")
	outf("package imports\n")
	outf("var stdlib = map[string][]string{\n")
	f := readAPI()
	sc := bufio.NewScanner(f)

	// The APIs of the syscall/js and unsafe packages need to be computed explicitly,
	// because they're not included in the GOROOT/api/go1.*.txt files at this time.
	pkgs := map[string]map[string]bool{
		"syscall/js": syms("syscall/js", "GOOS=js", "GOARCH=wasm"),
		"unsafe":     syms("unsafe"),
	}
	paths := []string{"syscall/js", "unsafe"}

	for sc.Scan() {
		l := sc.Text()
		if m := sym.FindStringSubmatch(l); m != nil {
			path, sym := m[1], m[2]

			if _, ok := pkgs[path]; !ok {
				pkgs[path] = map[string]bool{}
				paths = append(paths, path)
			}
			pkgs[path][sym] = true
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Strings(paths)
	for _, path := range paths {
		outf("\t%q: {\n", path)
		pkg := pkgs[path]
		var syms []string
		for sym := range pkg {
			syms = append(syms, sym)
		}
		sort.Strings(syms)
		for _, sym := range syms {
			outf("\t\t%q,\n", sym)
		}
		outf("},\n")
	}
	outf("}\n")
	fmtbuf, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("zstdlib.go", fmtbuf, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// readAPI opens an io.Reader that reads all stdlib API content.
func readAPI() io.Reader {
	entries, err := os.ReadDir(filepath.Join(runtime.GOROOT(), "api"))
	if err != nil {
		log.Fatal(err)
	}
	var readers []io.Reader
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, "go") && strings.HasSuffix(name, ".txt") {
			readers = append(readers, mustOpen(api(name)))
		}
	}
	return io.MultiReader(readers...)
}

// syms computes the exported symbols in the specified package.
func syms(pkg string, extraEnv ...string) map[string]bool {
	var env []string
	if len(extraEnv) != 0 {
		env = append(os.Environ(), extraEnv...)
	}
	pkgs, err := packages.Load(&packages.Config{Mode: packages.NeedTypes, Env: env}, pkg)
	if err != nil {
		log.Fatalln(err)
	} else if len(pkgs) != 1 {
		log.Fatalf("got %d packages, want one package %q", len(pkgs), pkg)
	}
	syms := make(map[string]bool)
	for _, name := range pkgs[0].Types.Scope().Names() {
		if token.IsExported(name) {
			syms[name] = true
		}
	}
	return syms
}
