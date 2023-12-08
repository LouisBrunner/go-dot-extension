// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextParagraph struct {
  obj gdc.ObjectPtr
}

func createTextParagraph(obj gdc.ObjectPtr) *TextParagraph {
  return &TextParagraph{
    obj: obj,
  }
}

func (me *TextParagraph) BaseClass() string {
  return "TextParagraph"
}
