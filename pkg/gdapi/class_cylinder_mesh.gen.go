// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CylinderMesh struct {
  obj gdc.ObjectPtr
}

func createCylinderMesh(obj gdc.ObjectPtr) *CylinderMesh {
  return &CylinderMesh{
    obj: obj,
  }
}

func (me *CylinderMesh) BaseClass() string {
  return "CylinderMesh"
}
