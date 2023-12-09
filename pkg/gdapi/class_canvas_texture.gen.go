// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasTexture struct {
  obj gdc.ObjectPtr
}

func (me *CanvasTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasTexture) BaseClass() string {
  return "CanvasTexture"
}



// Enums

func (me *CanvasTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CanvasTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CanvasTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CanvasTexture) SetDiffuseTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) GetDiffuseTexture()  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) SetNormalTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) GetNormalTexture()  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) SetSpecularTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) GetSpecularTexture()  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) SetSpecularColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) GetSpecularColor()  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) SetSpecularShininess(shininess float32, )  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) GetSpecularShininess()  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) SetTextureFilter(filter CanvasItemTextureFilter, )  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) GetTextureFilter()  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) SetTextureRepeat(repeat CanvasItemTextureRepeat, )  {
  panic("TODO: implement")
}

func  (me *CanvasTexture) GetTextureRepeat()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
