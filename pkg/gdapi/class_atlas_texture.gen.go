// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type AtlasTexture struct {
  Texture2D
}

func (me *AtlasTexture) BaseClass() string {
  return "AtlasTexture"
}

func NewAtlasTexture() *AtlasTexture {
  str := StringNameFromStr("AtlasTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AtlasTexture{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{atlas.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AtlasTexture) GetAtlas() Texture2D {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_atlas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AtlasTexture) SetRegion(region Rect2, )  {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AtlasTexture) GetRegion() Rect2 {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AtlasTexture) SetMargin(margin Rect2, )  {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{margin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AtlasTexture) GetMargin() Rect2 {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AtlasTexture) SetFilterClip(enable bool, )  {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_clip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AtlasTexture) HasFilterClip() bool {
  classNameV := StringNameFromStr("AtlasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_filter_clip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
