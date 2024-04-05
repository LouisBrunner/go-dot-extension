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

type ptrsForCameraAttributesPhysicalList struct {
  fnSetAperture gdc.MethodBindPtr
  fnGetAperture gdc.MethodBindPtr
  fnSetShutterSpeed gdc.MethodBindPtr
  fnGetShutterSpeed gdc.MethodBindPtr
  fnSetFocalLength gdc.MethodBindPtr
  fnGetFocalLength gdc.MethodBindPtr
  fnSetFocusDistance gdc.MethodBindPtr
  fnGetFocusDistance gdc.MethodBindPtr
  fnSetNear gdc.MethodBindPtr
  fnGetNear gdc.MethodBindPtr
  fnSetFar gdc.MethodBindPtr
  fnGetFar gdc.MethodBindPtr
  fnGetFov gdc.MethodBindPtr
  fnSetAutoExposureMaxExposureValue gdc.MethodBindPtr
  fnGetAutoExposureMaxExposureValue gdc.MethodBindPtr
  fnSetAutoExposureMinExposureValue gdc.MethodBindPtr
  fnGetAutoExposureMinExposureValue gdc.MethodBindPtr
}

var ptrsForCameraAttributesPhysical ptrsForCameraAttributesPhysicalList

func initCameraAttributesPhysicalPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CameraAttributesPhysical")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_aperture")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnSetAperture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_aperture")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnGetAperture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_shutter_speed")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnSetShutterSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_shutter_speed")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnGetShutterSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_focal_length")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnSetFocalLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_focal_length")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnGetFocalLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_focus_distance")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnSetFocusDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_focus_distance")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnGetFocusDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_near")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnSetNear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_near")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnGetNear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_far")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnSetFar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_far")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnGetFar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_fov")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnGetFov = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_auto_exposure_max_exposure_value")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnSetAutoExposureMaxExposureValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_auto_exposure_max_exposure_value")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnGetAutoExposureMaxExposureValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_auto_exposure_min_exposure_value")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnSetAutoExposureMinExposureValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_auto_exposure_min_exposure_value")
    defer methodName.Destroy()
    ptrsForCameraAttributesPhysical.fnGetAutoExposureMinExposureValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type CameraAttributesPhysical struct {
  CameraAttributes
}

func (me *CameraAttributesPhysical) BaseClass() string {
  return "CameraAttributesPhysical"
}

func NewCameraAttributesPhysical() *CameraAttributesPhysical {
  str := StringNameFromStr("CameraAttributesPhysical") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CameraAttributesPhysical{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CameraAttributesPhysical) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CameraAttributesPhysical) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CameraAttributesPhysical) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CameraAttributesPhysical) SetAperture(aperture float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&aperture) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnSetAperture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetAperture() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnGetAperture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetShutterSpeed(shutter_speed float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shutter_speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnSetShutterSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetShutterSpeed() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnGetShutterSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetFocalLength(focal_length float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&focal_length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnSetFocalLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetFocalLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnGetFocalLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetFocusDistance(focus_distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&focus_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnSetFocusDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetFocusDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnGetFocusDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetNear(near float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&near) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnSetNear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetNear() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnGetNear), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetFar(far float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&far) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnSetFar), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetFar() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnGetFar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) GetFov() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnGetFov), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetAutoExposureMaxExposureValue(exposure_value_max float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_value_max) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnSetAutoExposureMaxExposureValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetAutoExposureMaxExposureValue() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnGetAutoExposureMaxExposureValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetAutoExposureMinExposureValue(exposure_value_min float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_value_min) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnSetAutoExposureMinExposureValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetAutoExposureMinExposureValue() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPhysical.fnGetAutoExposureMinExposureValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
