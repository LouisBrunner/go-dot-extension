// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImageTexture struct {
  obj gdc.ObjectPtr
}

func (me *ImageTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImageTexture) BaseClass() string {
  return "ImageTexture"
}



// Enums

func (me *ImageTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImageTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImageTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  ImageTextureCreateFromImage(image Image, )  {
  panic("TODO: implement")
}

func  (me *ImageTexture) GetFormat()  {
  panic("TODO: implement")
}

func  (me *ImageTexture) SetImage(image Image, )  {
  panic("TODO: implement")
}

func  (me *ImageTexture) Update(image Image, )  {
  panic("TODO: implement")
}

func  (me *ImageTexture) SetSizeOverride(size Vector2i, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
