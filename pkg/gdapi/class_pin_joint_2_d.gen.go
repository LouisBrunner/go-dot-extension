// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PinJoint2D struct {
  obj gdc.ObjectPtr
}

func (me *PinJoint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PinJoint2D) BaseClass() string {
  return "PinJoint2D"
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

func  (me *PinJoint2D) SetSoftness(softness float32, )  {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_softness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&softness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PinJoint2D) GetSoftness() float32 {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_softness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PinJoint2D) SetAngularLimitLower(angular_limit_lower float32, )  {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_limit_lower")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_limit_lower), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PinJoint2D) GetAngularLimitLower() float32 {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_limit_lower")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PinJoint2D) SetAngularLimitUpper(angular_limit_upper float32, )  {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_limit_upper")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_limit_upper), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PinJoint2D) GetAngularLimitUpper() float32 {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_limit_upper")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PinJoint2D) SetMotorTargetVelocity(motor_target_velocity float32, )  {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motor_target_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&motor_target_velocity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PinJoint2D) GetMotorTargetVelocity() float32 {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motor_target_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PinJoint2D) SetMotorEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motor_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PinJoint2D) IsMotorEnabled() bool {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_motor_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PinJoint2D) SetAngularLimitEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_limit_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PinJoint2D) IsAngularLimitEnabled() bool {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_angular_limit_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
