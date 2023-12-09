// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderCubemapArray struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderCubemapArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderCubemapArray) BaseClass() string {
  return "PlaceholderCubemapArray"
}



// Enums

func (me *PlaceholderCubemapArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderCubemapArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderCubemapArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
