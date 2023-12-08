// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VehicleWheel3D struct {
  obj gdc.ObjectPtr
}

func createVehicleWheel3D(obj gdc.ObjectPtr) *VehicleWheel3D {
  return &VehicleWheel3D{
    obj: obj,
  }
}

func (me *VehicleWheel3D) BaseClass() string {
  return "VehicleWheel3D"
}
