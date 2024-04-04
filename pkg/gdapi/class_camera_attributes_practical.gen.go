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

type CameraAttributesPractical struct {
  CameraAttributes
}

func (me *CameraAttributesPractical) BaseClass() string {
  return "CameraAttributesPractical"
}

func NewCameraAttributesPractical() *CameraAttributesPractical {
  str := StringNameFromStr("CameraAttributesPractical") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CameraAttributesPractical{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CameraAttributesPractical) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CameraAttributesPractical) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CameraAttributesPractical) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CameraAttributesPractical) SetDofBlurFarEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_far_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) IsDofBlurFarEnabled() bool {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_dof_blur_far_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurFarDistance(distance float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_far_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurFarDistance() float64 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_far_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurFarTransition(distance float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_far_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurFarTransition() float64 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_far_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurNearEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_near_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) IsDofBlurNearEnabled() bool {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_dof_blur_near_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurNearDistance(distance float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_near_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurNearDistance() float64 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_near_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurNearTransition(distance float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_near_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurNearTransition() float64 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_near_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurAmount(amount float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurAmount() float64 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetAutoExposureMaxSensitivity(max_sensitivity float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_max_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_sensitivity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetAutoExposureMaxSensitivity() float64 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_max_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetAutoExposureMinSensitivity(min_sensitivity float64, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_min_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_sensitivity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetAutoExposureMinSensitivity() float64 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_min_sensitivity")
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
