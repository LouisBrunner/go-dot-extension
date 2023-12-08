// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CodeHighlighter struct {
  obj gdc.ObjectPtr
}

func createCodeHighlighter(obj gdc.ObjectPtr) *CodeHighlighter {
  return &CodeHighlighter{
    obj: obj,
  }
}

func (me *CodeHighlighter) BaseClass() string {
  return "CodeHighlighter"
}
