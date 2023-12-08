// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasGroup struct {
  obj gdc.ObjectPtr
}

func createCanvasGroup(obj gdc.ObjectPtr) *CanvasGroup {
  return &CanvasGroup{
    obj: obj,
  }
}

func (me *CanvasGroup) BaseClass() string {
  return "CanvasGroup"
}
