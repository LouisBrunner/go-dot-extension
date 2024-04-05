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

type ptrsForCameraAttributesPracticalList struct {
  fnSetDofBlurFarEnabled gdc.MethodBindPtr
  fnIsDofBlurFarEnabled gdc.MethodBindPtr
  fnSetDofBlurFarDistance gdc.MethodBindPtr
  fnGetDofBlurFarDistance gdc.MethodBindPtr
  fnSetDofBlurFarTransition gdc.MethodBindPtr
  fnGetDofBlurFarTransition gdc.MethodBindPtr
  fnSetDofBlurNearEnabled gdc.MethodBindPtr
  fnIsDofBlurNearEnabled gdc.MethodBindPtr
  fnSetDofBlurNearDistance gdc.MethodBindPtr
  fnGetDofBlurNearDistance gdc.MethodBindPtr
  fnSetDofBlurNearTransition gdc.MethodBindPtr
  fnGetDofBlurNearTransition gdc.MethodBindPtr
  fnSetDofBlurAmount gdc.MethodBindPtr
  fnGetDofBlurAmount gdc.MethodBindPtr
  fnSetAutoExposureMaxSensitivity gdc.MethodBindPtr
  fnGetAutoExposureMaxSensitivity gdc.MethodBindPtr
  fnSetAutoExposureMinSensitivity gdc.MethodBindPtr
  fnGetAutoExposureMinSensitivity gdc.MethodBindPtr
}

var ptrsForCameraAttributesPractical ptrsForCameraAttributesPracticalList

func initCameraAttributesPracticalPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CameraAttributesPractical")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_dof_blur_far_enabled")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnSetDofBlurFarEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_dof_blur_far_enabled")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnIsDofBlurFarEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_dof_blur_far_distance")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnSetDofBlurFarDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_dof_blur_far_distance")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnGetDofBlurFarDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_dof_blur_far_transition")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnSetDofBlurFarTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_dof_blur_far_transition")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnGetDofBlurFarTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_dof_blur_near_enabled")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnSetDofBlurNearEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_dof_blur_near_enabled")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnIsDofBlurNearEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_dof_blur_near_distance")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnSetDofBlurNearDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_dof_blur_near_distance")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnGetDofBlurNearDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_dof_blur_near_transition")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnSetDofBlurNearTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_dof_blur_near_transition")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnGetDofBlurNearTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_dof_blur_amount")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnSetDofBlurAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_dof_blur_amount")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnGetDofBlurAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_auto_exposure_max_sensitivity")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnSetAutoExposureMaxSensitivity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_auto_exposure_max_sensitivity")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnGetAutoExposureMaxSensitivity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_auto_exposure_min_sensitivity")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnSetAutoExposureMinSensitivity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_auto_exposure_min_sensitivity")
    defer methodName.Destroy()
    ptrsForCameraAttributesPractical.fnGetAutoExposureMinSensitivity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnSetDofBlurFarEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) IsDofBlurFarEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnIsDofBlurFarEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurFarDistance(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnSetDofBlurFarDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurFarDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnGetDofBlurFarDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurFarTransition(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnSetDofBlurFarTransition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurFarTransition() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnGetDofBlurFarTransition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurNearEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnSetDofBlurNearEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) IsDofBlurNearEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnIsDofBlurNearEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurNearDistance(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnSetDofBlurNearDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurNearDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnGetDofBlurNearDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurNearTransition(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnSetDofBlurNearTransition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurNearTransition() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnGetDofBlurNearTransition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetDofBlurAmount(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnSetDofBlurAmount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetDofBlurAmount() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnGetDofBlurAmount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetAutoExposureMaxSensitivity(max_sensitivity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_sensitivity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnSetAutoExposureMaxSensitivity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetAutoExposureMaxSensitivity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnGetAutoExposureMaxSensitivity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraAttributesPractical) SetAutoExposureMinSensitivity(min_sensitivity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_sensitivity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnSetAutoExposureMinSensitivity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraAttributesPractical) GetAutoExposureMinSensitivity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributesPractical.fnGetAutoExposureMinSensitivity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
