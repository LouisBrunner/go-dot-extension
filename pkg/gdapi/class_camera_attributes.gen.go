// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CameraAttributes struct {
  obj gdc.ObjectPtr
}

func (me *CameraAttributes) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CameraAttributes) BaseClass() string {
  return "CameraAttributes"
}



// Enums

func (me *CameraAttributes) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CameraAttributes) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CameraAttributes) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CameraAttributes) SetExposureMultiplier(multiplier float32, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exposure_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributes) GetExposureMultiplier() float32 {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exposure_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributes) SetExposureSensitivity(sensitivity float32, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exposure_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sensitivity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributes) GetExposureSensitivity() float32 {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exposure_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributes) SetAutoExposureEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributes) IsAutoExposureEnabled() bool {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_auto_exposure_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributes) SetAutoExposureSpeed(exposure_speed float32, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributes) GetAutoExposureSpeed() float32 {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributes) SetAutoExposureScale(exposure_grey float32, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_grey), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributes) GetAutoExposureScale() float32 {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CameraAttributes) GetPropExposureSensitivity() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributes) SetPropExposureSensitivity(value float32) {
  panic("TODO: implement")
}

func (me *CameraAttributes) GetPropExposureMultiplier() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributes) SetPropExposureMultiplier(value float32) {
  panic("TODO: implement")
}

func (me *CameraAttributes) GetPropAutoExposureEnabled() bool {
  panic("TODO: implement")
}

func (me *CameraAttributes) SetPropAutoExposureEnabled(value bool) {
  panic("TODO: implement")
}

func (me *CameraAttributes) GetPropAutoExposureScale() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributes) SetPropAutoExposureScale(value float32) {
  panic("TODO: implement")
}

func (me *CameraAttributes) GetPropAutoExposureSpeed() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributes) SetPropAutoExposureSpeed(value float32) {
  panic("TODO: implement")
}