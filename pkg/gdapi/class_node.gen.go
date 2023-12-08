// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Node struct {
  obj gdc.ObjectPtr
}

func createNode(obj gdc.ObjectPtr) *Node {
  return &Node{
    obj: obj,
  }
}

func (me *Node) BaseClass() string {
  return "Node"
}
