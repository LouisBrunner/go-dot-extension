// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CodeEdit struct {
  obj gdc.ObjectPtr
}

func createCodeEdit(obj gdc.ObjectPtr) *CodeEdit {
  return &CodeEdit{
    obj: obj,
  }
}

func (me *CodeEdit) BaseClass() string {
  return "CodeEdit"
}
