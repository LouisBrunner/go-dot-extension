// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GraphEdit struct {
  obj gdc.ObjectPtr
}

func (me *GraphEdit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GraphEdit) BaseClass() string {
  return "GraphEdit"
}

type GraphEditPanningScheme int
const (
  GraphEditScrollZooms GraphEditPanningScheme = 0
  GraphEditScrollPans GraphEditPanningScheme = 1
)
