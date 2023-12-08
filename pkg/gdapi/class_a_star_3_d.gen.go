// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AStar3D struct {
  obj gdc.ObjectPtr
}

func createAStar3D(obj gdc.ObjectPtr) *AStar3D {
  return &AStar3D{
    obj: obj,
  }
}

func (me *AStar3D) BaseClass() string {
  return "AStar3D"
}
