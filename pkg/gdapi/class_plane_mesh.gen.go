// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaneMesh struct {
  obj gdc.ObjectPtr
}

func (me *PlaneMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaneMesh) BaseClass() string {
  return "PlaneMesh"
}
