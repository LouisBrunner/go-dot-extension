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
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_aperture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&aperture), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetAperture() float64 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_aperture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetShutterSpeed(shutter_speed float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shutter_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shutter_speed), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetShutterSpeed() float64 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shutter_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetFocalLength(focal_length float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_focal_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&focal_length), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetFocalLength() float64 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_focal_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetFocusDistance(focus_distance float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_focus_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&focus_distance), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetFocusDistance() float64 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_focus_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetNear(near float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&near), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetNear() float64 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetFar(far float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&far), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetFar() float64 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) GetFov() float64 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetAutoExposureMaxExposureValue(exposure_value_max float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_max_exposure_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_value_max), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetAutoExposureMaxExposureValue() float64 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_max_exposure_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPhysical) SetAutoExposureMinExposureValue(exposure_value_min float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_min_exposure_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_value_min), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPhysical) GetAutoExposureMinExposureValue() float64 {
  classNameV := StringNameFromStr("CameraAttributesPhysical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_min_exposure_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
