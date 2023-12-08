// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FontFile struct {
  obj gdc.ObjectPtr
}

func createFontFile(obj gdc.ObjectPtr) *FontFile {
  return &FontFile{
    obj: obj,
  }
}

func (me *FontFile) BaseClass() string {
  return "FontFile"
}
