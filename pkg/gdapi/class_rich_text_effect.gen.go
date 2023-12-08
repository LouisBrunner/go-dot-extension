// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RichTextEffect struct {
  obj gdc.ObjectPtr
}

func createRichTextEffect(obj gdc.ObjectPtr) *RichTextEffect {
  return &RichTextEffect{
    obj: obj,
  }
}

func (me *RichTextEffect) BaseClass() string {
  return "RichTextEffect"
}
