// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedCubemapArray struct {
  obj gdc.ObjectPtr
}

func (me *CompressedCubemapArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CompressedCubemapArray) BaseClass() string {
  return "CompressedCubemapArray"
}



// Enums

func (me *CompressedCubemapArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CompressedCubemapArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedCubemapArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
