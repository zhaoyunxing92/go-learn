package fmt

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"strconv"
	"strings"
)

func Process(filename string, src []byte) ([]byte, error) {
	var (
		file *ast.File
		err  error
	)
	fileSet := token.NewFileSet()

	if file, err = parser.ParseFile(fileSet, filename, src, parser.ParseComments); err != nil {
		return nil, err
	}

	return formatFile(fileSet, file)
}

//groupImports group imports
func groupImports(f *ast.File) {
	if len(f.Decls) <= 1 {
		return
	}
	var sysSpec, internalSpec, dubboSpec []*ast.ImportSpec
	for _, spec := range f.Imports {
		name, _ := strconv.Unquote(spec.Path.Value)
		if strings.HasPrefix(name, "dubbo") {
			dubboSpec = append(dubboSpec, spec)
		} else if strings.ContainsAny(name, "/") {
			internalSpec = append(internalSpec, spec)
		} else {
			sysSpec = append(sysSpec, spec)
		}
	}
	idx := getImportsIndex(f)
	fmt.Println(idx)
	var sys = f.Decls[idx]
	gen, _ := sys.(*ast.GenDecl)
	for _, spec := range sysSpec {
		spec.Path.ValuePos = gen.Pos()
		gen.Specs = append(gen.Specs, spec)
	}
	//pos
	//f.Decls = append(f.Decls, sys)
	//var first *ast.GenDecl
	//for _, decl := range f.Decls {
	//	gen, ok := decl.(*ast.GenDecl)
	//	if !ok || gen.Tok != token.IMPORT || declImports(gen, "C") {
	//		continue
	//	}
	//	if first == nil {
	//		first = gen
	//		continue // Don't touch the first one.
	//	}
	//}
	//first.End()
	// Merge all the import declarations into the first one.
	//var first *ast.GenDecl
	//for i := 0; i < len(f.Decls); i++ {
	//	decl := f.Decls[i]
	//	gen, ok := decl.(*ast.GenDecl)
	//	if !ok || gen.Tok != token.IMPORT || declImports(gen, "C") {
	//		continue
	//	}
	//
	//	if first == nil {
	//		first = gen
	//		continue // Don't touch the first one.
	//	}
	//	/**
	//	&ast.GenDecl{
	//		Doc:    doc,
	//		TokPos: pos,
	//		Tok:    keyword,
	//		Lparen: lparen,
	//		Specs:  list,
	//		Rparen: rparen,
	//	}
	//	*/
	//	//gd := &ast.GenDecl{
	//	//	Doc:    nil,
	//	//	TokPos: pos,
	//	//	Tok:    keyword,
	//	//	Lparen: lparen,
	//	//	Specs:  make(ast.Spec, 0),
	//	//	Rparen: rparen,
	//	//}
	//	// We now know there is more than one package in this import
	//	// declaration. Ensure that it ends up parenthesized.
	//	first.Lparen = first.Pos()
	//	// Move the imports of the other import declaration to the first one.
	//	for _, spec := range gen.Specs {
	//		spec.(*ast.ImportSpec).Path.ValuePos = first.Pos()
	//		first.Specs = append(first.Specs, spec)
	//	}
	//
	//	f.Decls = append(f.Decls[:i], f.Decls[i+1:]...)
	//	i--
	//	fmt.Println("")
	//}
}
func getImportsIndex(f *ast.File) int {
	for i, decl := range f.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.IMPORT || declImports(gen, "C") {
			continue
		}
		return i
	}
	return 0
}

//mergeImports merge imports
func mergeImports(fileSet *token.FileSet, f *ast.File) {
	if len(f.Decls) <= 1 {
		return
	}

	// Merge all the import declarations into the first one.
	var first *ast.GenDecl
	for i := 0; i < len(f.Decls); i++ {
		decl := f.Decls[i]
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.IMPORT || declImports(gen, "C") {
			continue
		}
		fmt.Printf("decl=%v, pos=%d, lparen=%d ,rparen=%d \n", gen, gen.Pos(), gen.Lparen, gen.Rparen)
		if first == nil {
			first = gen
			continue // Don't touch the first one.
		}
		//position := fileSet.Position(first.Pos())

		// We now know there is more than one package in this import
		// declaration. Ensure that it ends up parenthesized.
		first.Lparen = first.Pos()
		// Move the imports of the other import declaration to the first one.
		for _, spec := range gen.Specs {
			spec.(*ast.ImportSpec).Path.ValuePos = first.Pos()
			first.Specs = append(first.Specs, spec)
		}

		f.Decls = append(f.Decls[:i], f.Decls[i+1:]...)
		i--
	}
}

// declImports reports whether gen contains an import of path.
// Taken from golang.org/x/tools/ast/astutil.
func declImports(gen *ast.GenDecl, path string) bool {
	if gen.Tok != token.IMPORT {
		return false
	}
	for _, spec := range gen.Specs {
		impspec := spec.(*ast.ImportSpec)
		if importPath(impspec) == path {
			return true
		}
	}
	return false
}

func importPath(s ast.Spec) string {
	t, err := strconv.Unquote(s.(*ast.ImportSpec).Path.Value)
	if err == nil {
		return t
	}
	return ""
}

func formatFile(fileSet *token.FileSet, file *ast.File) ([]byte, error) {
	mergeImports(fileSet, file)
	//groupImports(file)
	printConfig := &printer.Config{Mode: printer.TabIndent, Tabwidth: 8}
	var buf bytes.Buffer

	err := printConfig.Fprint(&buf, fileSet, file)
	if err != nil {
		return nil, err
	}
	out := buf.Bytes()
	if out, err = format.Source(out); err != nil {
		return nil, err
	}

	return out, nil
}
