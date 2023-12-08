// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkinReference struct {
  obj gdc.ObjectPtr
}

func createSkinReference(obj gdc.ObjectPtr) *SkinReference {
  return &SkinReference{
    obj: obj,
  }
}

func (me *SkinReference) BaseClass() string {
  return "SkinReference"
}
