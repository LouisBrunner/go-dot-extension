// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDVertexAttribute struct {
  obj gdc.ObjectPtr
}

func createRDVertexAttribute(obj gdc.ObjectPtr) *RDVertexAttribute {
  return &RDVertexAttribute{
    obj: obj,
  }
}

func (me *RDVertexAttribute) BaseClass() string {
  return "RDVertexAttribute"
}
