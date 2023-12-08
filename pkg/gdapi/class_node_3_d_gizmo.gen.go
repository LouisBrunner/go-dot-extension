// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Node3DGizmo struct {
  obj gdc.ObjectPtr
}

func createNode3DGizmo(obj gdc.ObjectPtr) *Node3DGizmo {
  return &Node3DGizmo{
    obj: obj,
  }
}

func (me *Node3DGizmo) BaseClass() string {
  return "Node3DGizmo"
}
