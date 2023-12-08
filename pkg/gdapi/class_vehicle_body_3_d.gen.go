// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VehicleBody3D struct {
  obj gdc.ObjectPtr
}

func createVehicleBody3D(obj gdc.ObjectPtr) *VehicleBody3D {
  return &VehicleBody3D{
    obj: obj,
  }
}

func (me *VehicleBody3D) BaseClass() string {
  return "VehicleBody3D"
}
