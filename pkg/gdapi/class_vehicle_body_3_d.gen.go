// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VehicleBody3D struct {
  obj gdc.ObjectPtr
}

func (me *VehicleBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VehicleBody3D) BaseClass() string {
  return "VehicleBody3D"
}
