// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenNotifier3D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenNotifier3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenNotifier3D) BaseClass() string {
  return "VisibleOnScreenNotifier3D"
}
