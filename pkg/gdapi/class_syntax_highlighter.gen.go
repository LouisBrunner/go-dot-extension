// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SyntaxHighlighter struct {
  obj gdc.ObjectPtr
}

func createSyntaxHighlighter(obj gdc.ObjectPtr) *SyntaxHighlighter {
  return &SyntaxHighlighter{
    obj: obj,
  }
}

func (me *SyntaxHighlighter) BaseClass() string {
  return "SyntaxHighlighter"
}
