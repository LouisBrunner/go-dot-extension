// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImageTexture3D struct {
  obj gdc.ObjectPtr
}

func (me *ImageTexture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImageTexture3D) BaseClass() string {
  return "ImageTexture3D"
}



// Enums

func (me *ImageTexture3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImageTexture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImageTexture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ImageTexture3D) Create(format ImageFormat, width int, height int, depth int, use_mipmaps bool, data Image, )  {
  panic("TODO: implement")
}

func  (me *ImageTexture3D) Update(data Image, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
