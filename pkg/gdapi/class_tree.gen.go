// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Tree struct {
  obj gdc.ObjectPtr
}

func createTree(obj gdc.ObjectPtr) *Tree {
  return &Tree{
    obj: obj,
  }
}

func (me *Tree) BaseClass() string {
  return "Tree"
}
