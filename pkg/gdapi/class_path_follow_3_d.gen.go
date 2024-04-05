// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForPathFollow3DList struct {
  fnSetProgress gdc.MethodBindPtr
  fnGetProgress gdc.MethodBindPtr
  fnSetHOffset gdc.MethodBindPtr
  fnGetHOffset gdc.MethodBindPtr
  fnSetVOffset gdc.MethodBindPtr
  fnGetVOffset gdc.MethodBindPtr
  fnSetProgressRatio gdc.MethodBindPtr
  fnGetProgressRatio gdc.MethodBindPtr
  fnSetRotationMode gdc.MethodBindPtr
  fnGetRotationMode gdc.MethodBindPtr
  fnSetCubicInterpolation gdc.MethodBindPtr
  fnGetCubicInterpolation gdc.MethodBindPtr
  fnSetUseModelFront gdc.MethodBindPtr
  fnIsUsingModelFront gdc.MethodBindPtr
  fnSetLoop gdc.MethodBindPtr
  fnHasLoop gdc.MethodBindPtr
  fnSetTiltEnabled gdc.MethodBindPtr
  fnIsTiltEnabled gdc.MethodBindPtr
  fnCorrectPosture gdc.MethodBindPtr
}

var ptrsForPathFollow3D ptrsForPathFollow3DList

func initPathFollow3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PathFollow3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_progress")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnSetProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_progress")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnGetProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_h_offset")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnSetHOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_h_offset")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnGetHOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_v_offset")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnSetVOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_v_offset")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnGetVOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_progress_ratio")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnSetProgressRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_progress_ratio")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnGetProgressRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_rotation_mode")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnSetRotationMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1640311967))
  }
  {
    methodName := StringNameFromStr("get_rotation_mode")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnGetRotationMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814010545))
  }
  {
    methodName := StringNameFromStr("set_cubic_interpolation")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnSetCubicInterpolation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_cubic_interpolation")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnGetCubicInterpolation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_use_model_front")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnSetUseModelFront = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_using_model_front")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnIsUsingModelFront = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_loop")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnSetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("has_loop")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnHasLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_tilt_enabled")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnSetTiltEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_tilt_enabled")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnIsTiltEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("correct_posture")
    defer methodName.Destroy()
    ptrsForPathFollow3D.fnCorrectPosture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2686588690))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&progress) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnSetProgress), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetProgress() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnGetProgress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetHOffset(h_offset float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&h_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnSetHOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetHOffset() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnGetHOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetVOffset(v_offset float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&v_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnSetVOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetVOffset() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnGetVOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetProgressRatio(ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnSetProgressRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetProgressRatio() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnGetProgressRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetRotationMode(rotation_mode PathFollow3DRotationMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnSetRotationMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetRotationMode() PathFollow3DRotationMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret PathFollow3DRotationMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnGetRotationMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PathFollow3D) SetCubicInterpolation(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnSetCubicInterpolation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) GetCubicInterpolation() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnGetCubicInterpolation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetUseModelFront(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnSetUseModelFront), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) IsUsingModelFront() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnIsUsingModelFront), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetLoop(loop bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnSetLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) HasLoop() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnHasLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PathFollow3D) SetTiltEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnSetTiltEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PathFollow3D) IsTiltEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnIsTiltEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  PathFollow3DCorrectPosture(transform Transform3D, rotation_mode PathFollow3DRotationMode, ) Transform3D {
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), gdc.ConstTypePtr(&rotation_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&rotation_mode)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow3D.fnCorrectPosture), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
