// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Theme struct {
  obj gdc.ObjectPtr
}

func createTheme(obj gdc.ObjectPtr) *Theme {
  return &Theme{
    obj: obj,
  }
}

func (me *Theme) BaseClass() string {
  return "Theme"
}
