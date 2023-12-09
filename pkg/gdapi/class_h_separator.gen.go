// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HSeparator struct {
  obj gdc.ObjectPtr
}

func (me *HSeparator) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HSeparator) BaseClass() string {
  return "HSeparator"
}



// Enums

func (me *HSeparator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HSeparator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HSeparator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
