// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorSyntaxHighlighter struct {
  obj gdc.ObjectPtr
}

func createEditorSyntaxHighlighter(obj gdc.ObjectPtr) *EditorSyntaxHighlighter {
  return &EditorSyntaxHighlighter{
    obj: obj,
  }
}

func (me *EditorSyntaxHighlighter) BaseClass() string {
  return "EditorSyntaxHighlighter"
}
