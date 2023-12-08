// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MissingNode struct {
  obj gdc.ObjectPtr
}

func createMissingNode(obj gdc.ObjectPtr) *MissingNode {
  return &MissingNode{
    obj: obj,
  }
}

func (me *MissingNode) BaseClass() string {
  return "MissingNode"
}
