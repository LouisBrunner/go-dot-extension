// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CylinderShape3D struct {
  obj gdc.ObjectPtr
}

func (me *CylinderShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CylinderShape3D) BaseClass() string {
  return "CylinderShape3D"
}
