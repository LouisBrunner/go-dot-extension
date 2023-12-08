// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *VisualInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualInstance3D) BaseClass() string {
  return "VisualInstance3D"
}
