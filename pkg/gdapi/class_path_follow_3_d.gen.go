// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PathFollow3D struct {
  obj gdc.ObjectPtr
}

func (me *PathFollow3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PathFollow3D) BaseClass() string {
  return "PathFollow3D"
}



// Enums

type PathFollow3DRotationMode int
const (
  PathFollow3DRotationModeRotationNone PathFollow3DRotationMode = 0
  PathFollow3DRotationModeRotationY PathFollow3DRotationMode = 1
  PathFollow3DRotationModeRotationXy PathFollow3DRotationMode = 2
  PathFollow3DRotationModeRotationXyz PathFollow3DRotationMode = 3
  PathFollow3DRotationModeRotationOriented PathFollow3DRotationMode = 4
)

func (me *PathFollow3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PathFollow3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PathFollow3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PathFollow3D) SetProgress(progress float32, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&progress), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow3D) GetProgress() float32 {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow3D) SetHOffset(h_offset float32, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&h_offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow3D) GetHOffset() float32 {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow3D) SetVOffset(v_offset float32, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&v_offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow3D) GetVOffset() float32 {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow3D) SetProgressRatio(ratio float32, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_progress_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow3D) GetProgressRatio() float32 {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_progress_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow3D) SetRotationMode(rotation_mode PathFollow3DRotationMode, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1640311967) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow3D) GetRotationMode() PathFollow3DRotationMode {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814010545) // FIXME: should cache?
  var ret PathFollow3DRotationMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow3D) SetCubicInterpolation(enabled bool, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cubic_interpolation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow3D) GetCubicInterpolation() bool {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cubic_interpolation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow3D) SetUseModelFront(enabled bool, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_model_front")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow3D) IsUsingModelFront() bool {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_model_front")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow3D) SetLoop(loop bool, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow3D) HasLoop() bool {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PathFollow3D) SetTiltEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tilt_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PathFollow3D) IsTiltEnabled() bool {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_tilt_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  PathFollow3DCorrectPosture(transform Transform3D, rotation_mode PathFollow3DRotationMode, ) Transform3D {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("correct_posture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2686588690) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(&rotation_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PathFollow3D) GetPropProgress() float32 {
  panic("TODO: implement")
}

func (me *PathFollow3D) SetPropProgress(value float32) {
  panic("TODO: implement")
}

func (me *PathFollow3D) GetPropProgressRatio() float32 {
  panic("TODO: implement")
}

func (me *PathFollow3D) SetPropProgressRatio(value float32) {
  panic("TODO: implement")
}

func (me *PathFollow3D) GetPropHOffset() float32 {
  panic("TODO: implement")
}

func (me *PathFollow3D) SetPropHOffset(value float32) {
  panic("TODO: implement")
}

func (me *PathFollow3D) GetPropVOffset() float32 {
  panic("TODO: implement")
}

func (me *PathFollow3D) SetPropVOffset(value float32) {
  panic("TODO: implement")
}

func (me *PathFollow3D) GetPropRotationMode() int {
  panic("TODO: implement")
}

func (me *PathFollow3D) SetPropRotationMode(value int) {
  panic("TODO: implement")
}

func (me *PathFollow3D) GetPropUseModelFront() bool {
  panic("TODO: implement")
}

func (me *PathFollow3D) SetPropUseModelFront(value bool) {
  panic("TODO: implement")
}

func (me *PathFollow3D) GetPropCubicInterp() bool {
  panic("TODO: implement")
}

func (me *PathFollow3D) SetPropCubicInterp(value bool) {
  panic("TODO: implement")
}

func (me *PathFollow3D) GetPropLoop() bool {
  panic("TODO: implement")
}

func (me *PathFollow3D) SetPropLoop(value bool) {
  panic("TODO: implement")
}

func (me *PathFollow3D) GetPropTiltEnabled() bool {
  panic("TODO: implement")
}

func (me *PathFollow3D) SetPropTiltEnabled(value bool) {
  panic("TODO: implement")
}