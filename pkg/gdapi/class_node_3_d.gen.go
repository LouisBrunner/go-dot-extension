// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Node3D struct {
  obj gdc.ObjectPtr
}

func createNode3D(obj gdc.ObjectPtr) *Node3D {
  return &Node3D{
    obj: obj,
  }
}

func (me *Node3D) BaseClass() string {
  return "Node3D"
}
