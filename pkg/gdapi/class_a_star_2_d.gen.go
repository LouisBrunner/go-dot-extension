// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AStar2D struct {
  obj gdc.ObjectPtr
}

func createAStar2D(obj gdc.ObjectPtr) *AStar2D {
  return &AStar2D{
    obj: obj,
  }
}

func (me *AStar2D) BaseClass() string {
  return "AStar2D"
}
