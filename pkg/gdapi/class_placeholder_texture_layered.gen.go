// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderTextureLayered struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderTextureLayered) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderTextureLayered) BaseClass() string {
  return "PlaceholderTextureLayered"
}



// Enums

func (me *PlaceholderTextureLayered) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderTextureLayered) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTextureLayered) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PlaceholderTextureLayered) SetSize(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *PlaceholderTextureLayered) GetSize()  {
  panic("TODO: implement")
}

func  (me *PlaceholderTextureLayered) SetLayers(layers int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
