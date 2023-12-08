// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Texture2DArray struct {
  obj gdc.ObjectPtr
}

func createTexture2DArray(obj gdc.ObjectPtr) *Texture2DArray {
  return &Texture2DArray{
    obj: obj,
  }
}

func (me *Texture2DArray) BaseClass() string {
  return "Texture2DArray"
}
