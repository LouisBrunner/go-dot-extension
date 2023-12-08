// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Translation struct {
  obj gdc.ObjectPtr
}

func createTranslation(obj gdc.ObjectPtr) *Translation {
  return &Translation{
    obj: obj,
  }
}

func (me *Translation) BaseClass() string {
  return "Translation"
}
