// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GraphNode struct {
  obj gdc.ObjectPtr
}

func createGraphNode(obj gdc.ObjectPtr) *GraphNode {
  return &GraphNode{
    obj: obj,
  }
}

func (me *GraphNode) BaseClass() string {
  return "GraphNode"
}
