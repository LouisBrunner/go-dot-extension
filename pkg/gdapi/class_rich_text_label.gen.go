// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RichTextLabel struct {
  obj gdc.ObjectPtr
}

func createRichTextLabel(obj gdc.ObjectPtr) *RichTextLabel {
  return &RichTextLabel{
    obj: obj,
  }
}

func (me *RichTextLabel) BaseClass() string {
  return "RichTextLabel"
}
