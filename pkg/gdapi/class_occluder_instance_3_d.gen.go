// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OccluderInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *OccluderInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OccluderInstance3D) BaseClass() string {
  return "OccluderInstance3D"
}
