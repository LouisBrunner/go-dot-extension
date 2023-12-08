// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FontVariation struct {
  obj gdc.ObjectPtr
}

func createFontVariation(obj gdc.ObjectPtr) *FontVariation {
  return &FontVariation{
    obj: obj,
  }
}

func (me *FontVariation) BaseClass() string {
  return "FontVariation"
}
