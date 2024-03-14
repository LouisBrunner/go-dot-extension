// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CameraAttributesPhysical struct {
  CameraAttributes
}

func (me *CameraAttributesPhysical) BaseClass() string {
  return "CameraAttributesPhysical"
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

func  (me *CameraAttributesPhysical) SetAperture(aperture float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_aperture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&aperture), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPhysical) GetAperture() float32 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_aperture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPhysical) SetShutterSpeed(shutter_speed float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shutter_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shutter_speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPhysical) GetShutterSpeed() float32 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shutter_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPhysical) SetFocalLength(focal_length float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_focal_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&focal_length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPhysical) GetFocalLength() float32 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_focal_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPhysical) SetFocusDistance(focus_distance float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_focus_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&focus_distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPhysical) GetFocusDistance() float32 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_focus_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPhysical) SetNear(near float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&near), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPhysical) GetNear() float32 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPhysical) SetFar(far float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&far), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPhysical) GetFar() float32 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPhysical) GetFov() float32 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPhysical) SetAutoExposureMaxExposureValue(exposure_value_max float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_max_exposure_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_value_max), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPhysical) GetAutoExposureMaxExposureValue() float32 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_max_exposure_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPhysical) SetAutoExposureMinExposureValue(exposure_value_min float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_min_exposure_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_value_min), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPhysical) GetAutoExposureMinExposureValue() float32 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_min_exposure_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
