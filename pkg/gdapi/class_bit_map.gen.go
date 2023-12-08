// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BitMap struct {
  obj gdc.ObjectPtr
}

func createBitMap(obj gdc.ObjectPtr) *BitMap {
  return &BitMap{
    obj: obj,
  }
}

func (me *BitMap) BaseClass() string {
  return "BitMap"
}
