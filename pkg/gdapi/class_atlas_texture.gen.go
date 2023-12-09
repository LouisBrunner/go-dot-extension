// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AtlasTexture struct {
  obj gdc.ObjectPtr
}

func (me *AtlasTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AtlasTexture) BaseClass() string {
  return "AtlasTexture"
}



// Enums

func (me *AtlasTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AtlasTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AtlasTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AtlasTexture) SetAtlas(atlas Texture2D, )  {
  panic("TODO: implement")
}

func  (me *AtlasTexture) GetAtlas()  {
  panic("TODO: implement")
}

func  (me *AtlasTexture) SetRegion(region Rect2, )  {
  panic("TODO: implement")
}

func  (me *AtlasTexture) GetRegion()  {
  panic("TODO: implement")
}

func  (me *AtlasTexture) SetMargin(margin Rect2, )  {
  panic("TODO: implement")
}

func  (me *AtlasTexture) GetMargin()  {
  panic("TODO: implement")
}

func  (me *AtlasTexture) SetFilterClip(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AtlasTexture) HasFilterClip()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
