// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ParallaxLayer struct {
  obj gdc.ObjectPtr
}

func (me *ParallaxLayer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ParallaxLayer) BaseClass() string {
  return "ParallaxLayer"
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
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ParallaxLayer) GetPropMotionScale() Vector2 {
  panic("TODO: implement")
}

func (me *ParallaxLayer) SetPropMotionScale(value Vector2) {
  panic("TODO: implement")
}

func (me *ParallaxLayer) GetPropMotionOffset() Vector2 {
  panic("TODO: implement")
}

func (me *ParallaxLayer) SetPropMotionOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *ParallaxLayer) GetPropMotionMirroring() Vector2 {
  panic("TODO: implement")
}

func (me *ParallaxLayer) SetPropMotionMirroring(value Vector2) {
  panic("TODO: implement")
}