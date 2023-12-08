// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SystemFont struct {
  obj gdc.ObjectPtr
}

func createSystemFont(obj gdc.ObjectPtr) *SystemFont {
  return &SystemFont{
    obj: obj,
  }
}

func (me *SystemFont) BaseClass() string {
  return "SystemFont"
}
