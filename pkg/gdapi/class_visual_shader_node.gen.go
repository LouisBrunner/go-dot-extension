// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNode struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNode(obj gdc.ObjectPtr) *VisualShaderNode {
  return &VisualShaderNode{
    obj: obj,
  }
}

func (me *VisualShaderNode) BaseClass() string {
  return "VisualShaderNode"
}
