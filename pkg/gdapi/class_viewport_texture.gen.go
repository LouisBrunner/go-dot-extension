// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ViewportTexture struct {
  obj gdc.ObjectPtr
}

func createViewportTexture(obj gdc.ObjectPtr) *ViewportTexture {
  return &ViewportTexture{
    obj: obj,
  }
}

func (me *ViewportTexture) BaseClass() string {
  return "ViewportTexture"
}
