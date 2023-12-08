// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Texture2D struct {
  obj gdc.ObjectPtr
}

func createTexture2D(obj gdc.ObjectPtr) *Texture2D {
  return &Texture2D{
    obj: obj,
  }
}

func (me *Texture2D) BaseClass() string {
  return "Texture2D"
}
