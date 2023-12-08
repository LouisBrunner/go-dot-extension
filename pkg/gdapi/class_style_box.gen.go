// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StyleBox struct {
  obj gdc.ObjectPtr
}

func createStyleBox(obj gdc.ObjectPtr) *StyleBox {
  return &StyleBox{
    obj: obj,
  }
}

func (me *StyleBox) BaseClass() string {
  return "StyleBox"
}
