// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GridMap struct {
  obj gdc.ObjectPtr
}

func (me *GridMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GridMap) BaseClass() string {
  return "GridMap"
}

const (
  GridMapINVALID_CELL_ITEM = -1
)
