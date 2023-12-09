// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextureLayered struct {
  obj gdc.ObjectPtr
}

func (me *TextureLayered) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureLayered) BaseClass() string {
  return "TextureLayered"
}



// Enums

type TextureLayeredLayeredType int
const (
  TextureLayeredLayeredTypeLayeredType2DArray TextureLayeredLayeredType = 0
  TextureLayeredLayeredTypeLayeredTypeCubemap TextureLayeredLayeredType = 1
  TextureLayeredLayeredTypeLayeredTypeCubemapArray TextureLayeredLayeredType = 2
)

func (me *TextureLayered) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureLayered) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureLayered) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TextureLayered) XGetFormat()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) XGetLayeredType()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) XGetWidth()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) XGetHeight()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) XGetLayers()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) XHasMipmaps()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) XGetLayerData(layer_index int, )  {
  panic("TODO: implement")
}

func  (me *TextureLayered) GetFormat()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) GetLayeredType()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) GetWidth()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) GetHeight()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) GetLayers()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) HasMipmaps()  {
  panic("TODO: implement")
}

func  (me *TextureLayered) GetLayerData(layer int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
