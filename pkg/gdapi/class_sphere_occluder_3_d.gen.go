// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SphereOccluder3D struct {
  obj gdc.ObjectPtr
}

func (me *SphereOccluder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SphereOccluder3D) BaseClass() string {
  return "SphereOccluder3D"
}
