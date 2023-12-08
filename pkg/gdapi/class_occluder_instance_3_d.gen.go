// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OccluderInstance3D struct {
  obj gdc.ObjectPtr
}

func createOccluderInstance3D(obj gdc.ObjectPtr) *OccluderInstance3D {
  return &OccluderInstance3D{
    obj: obj,
  }
}

func (me *OccluderInstance3D) BaseClass() string {
  return "OccluderInstance3D"
}
