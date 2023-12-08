// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TubeTrailMesh struct {
  obj gdc.ObjectPtr
}

func (me *TubeTrailMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TubeTrailMesh) BaseClass() string {
  return "TubeTrailMesh"
}
