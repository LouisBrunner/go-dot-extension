// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Separator struct {
  obj gdc.ObjectPtr
}

func (me *Separator) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Separator) BaseClass() string {
  return "Separator"
}



// Enums

func (me *Separator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Separator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Separator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
