// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Node3DGizmo struct {
  obj gdc.ObjectPtr
}

func (me *Node3DGizmo) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Node3DGizmo) BaseClass() string {
  return "Node3DGizmo"
}
