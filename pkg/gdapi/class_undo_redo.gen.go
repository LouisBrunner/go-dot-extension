// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UndoRedo struct {
  obj gdc.ObjectPtr
}

func createUndoRedo(obj gdc.ObjectPtr) *UndoRedo {
  return &UndoRedo{
    obj: obj,
  }
}

func (me *UndoRedo) BaseClass() string {
  return "UndoRedo"
}
