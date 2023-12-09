// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderCubemap struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderCubemap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderCubemap) BaseClass() string {
  return "PlaceholderCubemap"
}



// Enums

func (me *PlaceholderCubemap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderCubemap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderCubemap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
