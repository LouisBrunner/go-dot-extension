// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PanelContainer struct {
  obj gdc.ObjectPtr
}

func (me *PanelContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

// TODO: properties (class)

// TODO: signals (class)
