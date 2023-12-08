// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Node2D struct {
  obj gdc.ObjectPtr
}

func (me *Node2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Node2D) BaseClass() string {
  return "Node2D"
}
