// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VBoxContainer struct {
  obj gdc.ObjectPtr
}

func (me *VBoxContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VBoxContainer) BaseClass() string {
  return "VBoxContainer"
}



// Enums

func (me *VBoxContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VBoxContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VBoxContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
