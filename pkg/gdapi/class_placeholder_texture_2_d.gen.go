// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderTexture2D struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderTexture2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderTexture2D) BaseClass() string {
  return "PlaceholderTexture2D"
}



// Enums

func (me *PlaceholderTexture2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderTexture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTexture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PlaceholderTexture2D) SetSize(size Vector2, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
