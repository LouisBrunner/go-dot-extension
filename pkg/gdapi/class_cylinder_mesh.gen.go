// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CylinderMesh struct {
  obj gdc.ObjectPtr
}

func (me *CylinderMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CylinderMesh) BaseClass() string {
  return "CylinderMesh"
}
