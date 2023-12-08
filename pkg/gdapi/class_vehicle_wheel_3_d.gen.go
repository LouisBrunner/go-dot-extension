// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VehicleWheel3D struct {
  obj gdc.ObjectPtr
}

func (me *VehicleWheel3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VehicleWheel3D) BaseClass() string {
  return "VehicleWheel3D"
}
