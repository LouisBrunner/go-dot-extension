// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PanelContainer struct {
  Container
}

func (me *PanelContainer) BaseClass() string {
  return "PanelContainer"
}



// Enums

func (me *PanelContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PanelContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PanelContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
