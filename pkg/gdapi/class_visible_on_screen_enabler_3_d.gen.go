// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenEnabler3D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenEnabler3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenEnabler3D) BaseClass() string {
  return "VisibleOnScreenEnabler3D"
}
