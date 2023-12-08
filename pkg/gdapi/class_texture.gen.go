// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Texture struct {
  obj gdc.ObjectPtr
}

func createTexture(obj gdc.ObjectPtr) *Texture {
  return &Texture{
    obj: obj,
  }
}

func (me *Texture) BaseClass() string {
  return "Texture"
}
