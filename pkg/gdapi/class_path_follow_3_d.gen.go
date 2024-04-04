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

type PathFollow3D struct {
  Node3D
}

func (me *PathFollow3D) BaseClass() string {
  return "PathFollow3D"
}

func NewPathFollow3D() *PathFollow3D {
  str := StringNameFromStr("PathFollow3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PathFollow3D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *PathFollow3D) SetProgress(progress float64, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&progress) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetProgress() float64 {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetHOffset(h_offset float64, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&h_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetHOffset() float64 {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetVOffset(v_offset float64, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&v_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetVOffset() float64 {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetProgressRatio(ratio float64, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_progress_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetProgressRatio() float64 {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_progress_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetRotationMode(rotation_mode PathFollow3DRotationMode, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1640311967) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetRotationMode() PathFollow3DRotationMode {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814010545) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret PathFollow3DRotationMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PathFollow3D) SetCubicInterpolation(enabled bool, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cubic_interpolation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetCubicInterpolation() bool {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cubic_interpolation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetUseModelFront(enabled bool, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_model_front")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) IsUsingModelFront() bool {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_model_front")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetLoop(loop bool, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) HasLoop() bool {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetTiltEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tilt_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) IsTiltEnabled() bool {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_tilt_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  PathFollow3DCorrectPosture(transform Transform3D, rotation_mode PathFollow3DRotationMode, ) Transform3D {
  classNameV := StringNameFromStr("PathFollow3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("correct_posture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2686588690) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), gdc.ConstTypePtr(&rotation_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&rotation_mode)

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
