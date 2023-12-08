// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AStar3D struct {
  obj gdc.ObjectPtr
}

func (me *AStar3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AStar3D) BaseClass() string {
  return "AStar3D"
}
