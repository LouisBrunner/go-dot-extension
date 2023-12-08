// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenNotifier2D struct {
  obj gdc.ObjectPtr
}

func createVisibleOnScreenNotifier2D(obj gdc.ObjectPtr) *VisibleOnScreenNotifier2D {
  return &VisibleOnScreenNotifier2D{
    obj: obj,
  }
}

func (me *VisibleOnScreenNotifier2D) BaseClass() string {
  return "VisibleOnScreenNotifier2D"
}
