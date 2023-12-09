// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PointLight2D struct {
  obj gdc.ObjectPtr
}

func (me *PointLight2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PointLight2D) BaseClass() string {
  return "PointLight2D"
}



// Enums

func (me *PointLight2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PointLight2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PointLight2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PointLight2D) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *PointLight2D) GetTexture()  {
  panic("TODO: implement")
}

func  (me *PointLight2D) SetTextureOffset(texture_offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *PointLight2D) GetTextureOffset()  {
  panic("TODO: implement")
}

func  (me *PointLight2D) SetTextureScale(texture_scale float32, )  {
  panic("TODO: implement")
}

func  (me *PointLight2D) GetTextureScale()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
