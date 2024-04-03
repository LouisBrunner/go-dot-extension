// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ParallaxLayer struct {
  Node2D
}

func (me *ParallaxLayer) BaseClass() string {
  return "ParallaxLayer"
}

func NewParallaxLayer() *ParallaxLayer {
  str := StringNameFromStr("ParallaxLayer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ParallaxLayer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ParallaxLayer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ParallaxLayer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ParallaxLayer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ParallaxLayer) SetMotionScale(scale Vector2, )  {
  classNameV := StringNameFromStr("ParallaxLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParallaxLayer) GetMotionScale() Vector2 {
  classNameV := StringNameFromStr("ParallaxLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParallaxLayer) SetMotionOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("ParallaxLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParallaxLayer) GetMotionOffset() Vector2 {
  classNameV := StringNameFromStr("ParallaxLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParallaxLayer) SetMirroring(mirror Vector2, )  {
  classNameV := StringNameFromStr("ParallaxLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mirroring")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mirror.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParallaxLayer) GetMirroring() Vector2 {
  classNameV := StringNameFromStr("ParallaxLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mirroring")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
