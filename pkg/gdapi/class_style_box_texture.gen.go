// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StyleBoxTexture struct {
  obj gdc.ObjectPtr
}

func createStyleBoxTexture(obj gdc.ObjectPtr) *StyleBoxTexture {
  return &StyleBoxTexture{
    obj: obj,
  }
}

func (me *StyleBoxTexture) BaseClass() string {
  return "StyleBoxTexture"
}
