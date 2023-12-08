// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextServerManager struct {
  obj gdc.ObjectPtr
}

func createTextServerManager(obj gdc.ObjectPtr) *TextServerManager {
  return &TextServerManager{
    obj: obj,
  }
}

func (me *TextServerManager) BaseClass() string {
  return "TextServerManager"
}
