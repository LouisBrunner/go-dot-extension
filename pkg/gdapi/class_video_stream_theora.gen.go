// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VideoStreamTheora struct {
  obj gdc.ObjectPtr
}

func (me *VideoStreamTheora) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VideoStreamTheora) BaseClass() string {
  return "VideoStreamTheora"
}



// Enums

func (me *VideoStreamTheora) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VideoStreamTheora) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VideoStreamTheora) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
