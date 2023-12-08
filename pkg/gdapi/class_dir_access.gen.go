// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DirAccess struct {
  obj gdc.ObjectPtr
}

func createDirAccess(obj gdc.ObjectPtr) *DirAccess {
  return &DirAccess{
    obj: obj,
  }
}

func (me *DirAccess) BaseClass() string {
  return "DirAccess"
}
