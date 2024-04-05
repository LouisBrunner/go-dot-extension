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

type ptrsForPinJoint2DList struct {
  fnSetSoftness gdc.MethodBindPtr
  fnGetSoftness gdc.MethodBindPtr
  fnSetAngularLimitLower gdc.MethodBindPtr
  fnGetAngularLimitLower gdc.MethodBindPtr
  fnSetAngularLimitUpper gdc.MethodBindPtr
  fnGetAngularLimitUpper gdc.MethodBindPtr
  fnSetMotorTargetVelocity gdc.MethodBindPtr
  fnGetMotorTargetVelocity gdc.MethodBindPtr
  fnSetMotorEnabled gdc.MethodBindPtr
  fnIsMotorEnabled gdc.MethodBindPtr
  fnSetAngularLimitEnabled gdc.MethodBindPtr
  fnIsAngularLimitEnabled gdc.MethodBindPtr
}

var ptrsForPinJoint2D ptrsForPinJoint2DList

func initPinJoint2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PinJoint2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_softness")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnSetSoftness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_softness")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnGetSoftness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_angular_limit_lower")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnSetAngularLimitLower = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_angular_limit_lower")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnGetAngularLimitLower = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_angular_limit_upper")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnSetAngularLimitUpper = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_angular_limit_upper")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnGetAngularLimitUpper = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_motor_target_velocity")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnSetMotorTargetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_motor_target_velocity")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnGetMotorTargetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_motor_enabled")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnSetMotorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_motor_enabled")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnIsMotorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_angular_limit_enabled")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnSetAngularLimitEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_angular_limit_enabled")
    defer methodName.Destroy()
    ptrsForPinJoint2D.fnIsAngularLimitEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type PinJoint2D struct {
  Joint2D
}

func (me *PinJoint2D) BaseClass() string {
  return "PinJoint2D"
}

func NewPinJoint2D() *PinJoint2D {
  str := StringNameFromStr("PinJoint2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PinJoint2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PinJoint2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PinJoint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PinJoint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PinJoint2D) SetSoftness(softness float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&softness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnSetSoftness), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PinJoint2D) GetSoftness() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnGetSoftness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PinJoint2D) SetAngularLimitLower(angular_limit_lower float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_limit_lower) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnSetAngularLimitLower), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PinJoint2D) GetAngularLimitLower() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnGetAngularLimitLower), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PinJoint2D) SetAngularLimitUpper(angular_limit_upper float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_limit_upper) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnSetAngularLimitUpper), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PinJoint2D) GetAngularLimitUpper() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnGetAngularLimitUpper), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PinJoint2D) SetMotorTargetVelocity(motor_target_velocity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&motor_target_velocity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnSetMotorTargetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PinJoint2D) GetMotorTargetVelocity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnGetMotorTargetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PinJoint2D) SetMotorEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnSetMotorEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PinJoint2D) IsMotorEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnIsMotorEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PinJoint2D) SetAngularLimitEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnSetAngularLimitEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PinJoint2D) IsAngularLimitEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint2D.fnIsAngularLimitEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
