// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Texture struct {
  obj gdc.ObjectPtr
}

func (me *Texture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture) BaseClass() string {
  return "Texture"
}



// Enums

func (me *Texture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Texture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
