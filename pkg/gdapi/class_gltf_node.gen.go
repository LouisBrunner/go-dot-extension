// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFNode struct {
  obj gdc.ObjectPtr
}

func createGLTFNode(obj gdc.ObjectPtr) *GLTFNode {
  return &GLTFNode{
    obj: obj,
  }
}

func (me *GLTFNode) BaseClass() string {
  return "GLTFNode"
}
