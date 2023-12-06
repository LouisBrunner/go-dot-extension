package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"modernc.org/cc/v4"
)

const (
	prefixEnum           = "GDEXTENSION_"
	prefixTypes          = "GDExtension"
	prefixShort          = "GD"
	prefixInterfaceFuncs = "Interface"
)

func walkAST(ast *cc.AST, fn func(decl *cc.Declarator) error) error {
	unit := ast.TranslationUnit
	for unit != nil {
		decl := unit.ExternalDeclaration
		switch decl.Case {
		case cc.ExternalDeclarationDecl:
			decl := decl.Declaration
			switch {
			case decl.InitDeclaratorList != nil:
				entry := decl.InitDeclaratorList
				for entry != nil {
					err := fn(entry.InitDeclarator.Declarator)
					if err != nil {
						return err
					}
					entry = entry.InitDeclaratorList
				}
			default:
				return fmt.Errorf("unsupported declaration case %d", decl.Case)
			}

		default:
			return fmt.Errorf("unsupported external declaration case %d", decl.Case)
		}
		unit = unit.TranslationUnit
	}

	return nil
}

//go:embed templates/*
var templates embed.FS

func writeTemplate[T any](w io.Writer, templateName string, data T) error {
	templateFilename := fmt.Sprintf("%s.tmpl", templateName)
	tmpl, err := template.New(templateFilename).Funcs(map[string]interface{}{
		//
	}).ParseFS(templates, filepath.Join("templates", templateFilename))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

const (
	fileTypes      = "types.gen.go"
	fileInterface  = "interface.gen.go"
	fileInterfaceH = "interface.gen.h"
	fileInterfaceC = "interface.gen.c"
)

type outputFiles struct {
	files      map[string]*os.File
	ifaceFuncs map[string]string
	allFuncs   map[string]*funcType
}

func output(ast *cc.AST, ifaceFuncs map[string]string, outputDir string) error {
	out := &outputFiles{
		files:      make(map[string]*os.File),
		ifaceFuncs: ifaceFuncs,
		allFuncs:   make(map[string]*funcType),
	}
	for _, fileName := range []string{
		fileTypes,
		fileInterface,
		fileInterfaceH,
		fileInterfaceC,
	} {
		file, err := os.Create(filepath.Join(outputDir, fileName))
		if err != nil {
			return err
		}
		defer file.Close()
		err = writeTemplate[any](file, "preamble.all", nil)
		if err != nil {
			return err
		}
		if strings.HasSuffix(fileName, ".go") {
			err = writeTemplate[any](file, "preamble.go", map[string]interface{}{
				"File": fileName,
			})
			if err != nil {
				return err
			}
		}
		out.files[fileName] = file
	}

	err := walkAST(ast, func(decl *cc.Declarator) error {
		name, err := identifierOf(decl.DirectDeclarator)
		if err != nil {
			return err
		}
		if !strings.HasPrefix(name, prefixTypes) {
			return nil
		}
		typ, err := transformType(name, decl.Type(), decl.IsConst())
		if err != nil {
			return err
		}
		return handleType(out, typ)
	})
	if err != nil {
		return err
	}

	err = generateInterface(out)
	if err != nil {
		return err
	}

	return nil
}
