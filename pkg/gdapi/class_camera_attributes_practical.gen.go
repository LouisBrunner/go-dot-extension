// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CameraAttributesPractical struct {
  obj gdc.ObjectPtr
}

func (me *CameraAttributesPractical) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CameraAttributesPractical) BaseClass() string {
  return "CameraAttributesPractical"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPractical) IsDofBlurFarEnabled() bool {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_dof_blur_far_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPractical) SetDofBlurFarDistance(distance float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_far_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPractical) GetDofBlurFarDistance() float32 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_far_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPractical) SetDofBlurFarTransition(distance float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_far_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPractical) GetDofBlurFarTransition() float32 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_far_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPractical) SetDofBlurNearEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_near_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPractical) IsDofBlurNearEnabled() bool {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_dof_blur_near_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPractical) SetDofBlurNearDistance(distance float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_near_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPractical) GetDofBlurNearDistance() float32 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_near_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPractical) SetDofBlurNearTransition(distance float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_near_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPractical) GetDofBlurNearTransition() float32 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_near_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPractical) SetDofBlurAmount(amount float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dof_blur_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPractical) GetDofBlurAmount() float32 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dof_blur_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPractical) SetAutoExposureMaxSensitivity(max_sensitivity float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_max_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_sensitivity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPractical) GetAutoExposureMaxSensitivity() float32 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_max_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraAttributesPractical) SetAutoExposureMinSensitivity(min_sensitivity float32, )  {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_exposure_min_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_sensitivity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraAttributesPractical) GetAutoExposureMinSensitivity() float32 {
  classNameV := StringNameFromStr("CameraAttributesPractical")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_exposure_min_sensitivity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CameraAttributesPractical) GetPropDofBlurFarEnabled() bool {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) SetPropDofBlurFarEnabled(value bool) {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) GetPropDofBlurFarDistance() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) SetPropDofBlurFarDistance(value float32) {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) GetPropDofBlurFarTransition() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) SetPropDofBlurFarTransition(value float32) {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) GetPropDofBlurNearEnabled() bool {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) SetPropDofBlurNearEnabled(value bool) {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) GetPropDofBlurNearDistance() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) SetPropDofBlurNearDistance(value float32) {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) GetPropDofBlurNearTransition() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) SetPropDofBlurNearTransition(value float32) {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) GetPropDofBlurAmount() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) SetPropDofBlurAmount(value float32) {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) GetPropAutoExposureMinSensitivity() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) SetPropAutoExposureMinSensitivity(value float32) {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) GetPropAutoExposureMaxSensitivity() float32 {
  panic("TODO: implement")
}

func (me *CameraAttributesPractical) SetPropAutoExposureMaxSensitivity(value float32) {
  panic("TODO: implement")
}