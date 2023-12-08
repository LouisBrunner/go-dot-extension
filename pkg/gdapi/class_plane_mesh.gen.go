// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaneMesh struct {
  obj gdc.ObjectPtr
}

func createPlaneMesh(obj gdc.ObjectPtr) *PlaneMesh {
  return &PlaneMesh{
    obj: obj,
  }
}

func (me *PlaneMesh) BaseClass() string {
  return "PlaneMesh"
}
