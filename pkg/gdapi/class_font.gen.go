// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Font struct {
  obj gdc.ObjectPtr
}

func createFont(obj gdc.ObjectPtr) *Font {
  return &Font{
    obj: obj,
  }
}

func (me *Font) BaseClass() string {
  return "Font"
}
