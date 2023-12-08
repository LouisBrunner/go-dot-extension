// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoxOccluder3D struct {
  obj gdc.ObjectPtr
}

func (me *BoxOccluder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BoxOccluder3D) BaseClass() string {
  return "BoxOccluder3D"
}
