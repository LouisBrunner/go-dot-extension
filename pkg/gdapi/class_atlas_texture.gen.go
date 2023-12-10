// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_atlas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AtlasTexture) GetAtlas() Texture2D {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_atlas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AtlasTexture) SetRegion(region Rect2, )  {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AtlasTexture) GetRegion() Rect2 {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AtlasTexture) SetMargin(margin Rect2, )  {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(margin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AtlasTexture) GetMargin() Rect2 {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AtlasTexture) SetFilterClip(enable bool, )  {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_clip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AtlasTexture) HasFilterClip() bool {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_filter_clip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AtlasTexture) GetPropAtlas() Texture2D {
  panic("TODO: implement")
}

func (me *AtlasTexture) SetPropAtlas(value Texture2D) {
  panic("TODO: implement")
}

func (me *AtlasTexture) GetPropRegion() Rect2 {
  panic("TODO: implement")
}

func (me *AtlasTexture) SetPropRegion(value Rect2) {
  panic("TODO: implement")
}

func (me *AtlasTexture) GetPropMargin() Rect2 {
  panic("TODO: implement")
}

func (me *AtlasTexture) SetPropMargin(value Rect2) {
  panic("TODO: implement")
}

func (me *AtlasTexture) GetPropFilterClip() bool {
  panic("TODO: implement")
}

func (me *AtlasTexture) SetPropFilterClip(value bool) {
  panic("TODO: implement")
}