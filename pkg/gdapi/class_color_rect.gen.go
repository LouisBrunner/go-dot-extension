// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ColorRect struct {
  obj gdc.ObjectPtr
}

func createColorRect(obj gdc.ObjectPtr) *ColorRect {
  return &ColorRect{
    obj: obj,
  }
}

func (me *ColorRect) BaseClass() string {
  return "ColorRect"
}
