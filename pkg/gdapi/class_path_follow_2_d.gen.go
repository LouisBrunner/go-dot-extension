// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PathFollow2D struct {
  obj gdc.ObjectPtr
}

func (me *PathFollow2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PathFollow2D) BaseClass() string {
  return "PathFollow2D"
}



// Enums

func (me *PathFollow2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PathFollow2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PathFollow2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PathFollow2D) SetProgress(progress float32, )  {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&progress), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow2D) GetProgress() float32 {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow2D) SetHOffset(h_offset float32, )  {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&h_offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow2D) GetHOffset() float32 {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow2D) SetVOffset(v_offset float32, )  {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&v_offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow2D) GetVOffset() float32 {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow2D) SetProgressRatio(ratio float32, )  {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_progress_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow2D) GetProgressRatio() float32 {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_progress_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow2D) SetRotates(enabled bool, )  {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow2D) IsRotating() bool {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_rotating")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow2D) SetCubicInterpolation(enabled bool, )  {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cubic_interpolation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow2D) GetCubicInterpolation() bool {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cubic_interpolation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow2D) SetLoop(loop bool, )  {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow2D) HasLoop() bool {
  classNameV := StringNameFromStr("PathFollow2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PathFollow2D) GetPropProgress() float32 {
  panic("TODO: implement")
}

func (me *PathFollow2D) SetPropProgress(value float32) {
  panic("TODO: implement")
}

func (me *PathFollow2D) GetPropProgressRatio() float32 {
  panic("TODO: implement")
}

func (me *PathFollow2D) SetPropProgressRatio(value float32) {
  panic("TODO: implement")
}

func (me *PathFollow2D) GetPropHOffset() float32 {
  panic("TODO: implement")
}

func (me *PathFollow2D) SetPropHOffset(value float32) {
  panic("TODO: implement")
}

func (me *PathFollow2D) GetPropVOffset() float32 {
  panic("TODO: implement")
}

func (me *PathFollow2D) SetPropVOffset(value float32) {
  panic("TODO: implement")
}

func (me *PathFollow2D) GetPropRotates() bool {
  panic("TODO: implement")
}

func (me *PathFollow2D) SetPropRotates(value bool) {
  panic("TODO: implement")
}

func (me *PathFollow2D) GetPropCubicInterp() bool {
  panic("TODO: implement")
}

func (me *PathFollow2D) SetPropCubicInterp(value bool) {
  panic("TODO: implement")
}

func (me *PathFollow2D) GetPropLoop() bool {
  panic("TODO: implement")
}

func (me *PathFollow2D) SetPropLoop(value bool) {
  panic("TODO: implement")
}