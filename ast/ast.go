package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	path := "/Users/docker/code/github/go/go-learn/ast/service"

	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, path, func(fi os.FileInfo) bool {
		return true
	}, parser.ParseComments)

	if err != nil {
		panic(err)
	}
	for _, pk := range packages {
		//全部文件
		for _, f := range pk.Files {
			var first *ast.GenDecl
			for i := 0; i < len(f.Decls); i++ {
				decl := f.Decls[i]
				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.IMPORT {
					continue
				}
				if first == nil {
					first = gen
					continue // Don't touch the first one.
				}
				// We now know there is more than one package in this fmt
				// declaration. Ensure that it ends up parenthesized.
				first.Lparen = first.Pos()
				// Move the imports of the other fmt declaration to the first one.
				for _, spec := range gen.Specs {
					spec.(*ast.ImportSpec).Path.ValuePos = first.Pos()
					first.Specs = append(first.Specs, spec)
				}
				f.Decls = append(f.Decls[:i], f.Decls[i+1:]...)
				i--
			}
		}
	}
}
