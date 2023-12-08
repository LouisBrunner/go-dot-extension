// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenNotifier3D struct {
  obj gdc.ObjectPtr
}

func createVisibleOnScreenNotifier3D(obj gdc.ObjectPtr) *VisibleOnScreenNotifier3D {
  return &VisibleOnScreenNotifier3D{
    obj: obj,
  }
}

func (me *VisibleOnScreenNotifier3D) BaseClass() string {
  return "VisibleOnScreenNotifier3D"
}
