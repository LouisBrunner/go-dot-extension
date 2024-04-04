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

type CameraAttributes struct {
  Resource
}

func (me *CameraAttributes) BaseClass() string {
  return "CameraAttributes"
}

func NewCameraAttributes() *CameraAttributes {
  str := StringNameFromStr("CameraAttributes") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CameraAttributes{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *CameraAttributes) SetExposureMultiplier(multiplier float64, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exposure_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributes) GetExposureMultiplier() float64 {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exposure_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributes) SetExposureSensitivity(sensitivity float64, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exposure_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sensitivity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributes) GetExposureSensitivity() float64 {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exposure_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributes) SetAutoExposureEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributes) IsAutoExposureEnabled() bool {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_auto_exposure_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributes) SetAutoExposureSpeed(exposure_speed float64, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributes) GetAutoExposureSpeed() float64 {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributes) SetAutoExposureScale(exposure_grey float64, )  {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_grey) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributes) GetAutoExposureScale() float64 {
  classNameV := StringNameFromStr("CameraAttributes")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
