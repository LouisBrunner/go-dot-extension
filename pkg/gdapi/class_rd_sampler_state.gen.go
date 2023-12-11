// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDSamplerState struct {
  obj gdc.ObjectPtr
}

func (me *RDSamplerState) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDSamplerState) BaseClass() string {
  return "RDSamplerState"
}



// Enums

func (me *RDSamplerState) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDSamplerState) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDSamplerState) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDSamplerState) SetMagFilter(p_member RenderingDeviceSamplerFilter, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mag_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1493420382) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetMagFilter() RenderingDeviceSamplerFilter {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mag_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2209202801) // FIXME: should cache?
  var ret RenderingDeviceSamplerFilter
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetMinFilter(p_member RenderingDeviceSamplerFilter, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1493420382) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetMinFilter() RenderingDeviceSamplerFilter {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2209202801) // FIXME: should cache?
  var ret RenderingDeviceSamplerFilter
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetMipFilter(p_member RenderingDeviceSamplerFilter, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mip_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1493420382) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetMipFilter() RenderingDeviceSamplerFilter {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mip_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2209202801) // FIXME: should cache?
  var ret RenderingDeviceSamplerFilter
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetRepeatU(p_member RenderingDeviceSamplerRepeatMode, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_repeat_u")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 246127626) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetRepeatU() RenderingDeviceSamplerRepeatMode {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_repeat_u")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227895872) // FIXME: should cache?
  var ret RenderingDeviceSamplerRepeatMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetRepeatV(p_member RenderingDeviceSamplerRepeatMode, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_repeat_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 246127626) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetRepeatV() RenderingDeviceSamplerRepeatMode {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_repeat_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227895872) // FIXME: should cache?
  var ret RenderingDeviceSamplerRepeatMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetRepeatW(p_member RenderingDeviceSamplerRepeatMode, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_repeat_w")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 246127626) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetRepeatW() RenderingDeviceSamplerRepeatMode {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_repeat_w")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227895872) // FIXME: should cache?
  var ret RenderingDeviceSamplerRepeatMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetLodBias(p_member float32, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lod_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetLodBias() float32 {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lod_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetUseAnisotropy(p_member bool, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetUseAnisotropy() bool {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetAnisotropyMax(p_member float32, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anisotropy_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetAnisotropyMax() float32 {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_anisotropy_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetEnableCompare(p_member bool, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_compare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetEnableCompare() bool {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_compare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetCompareOp(p_member RenderingDeviceCompareOperator, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_compare_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2573711505) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetCompareOp() RenderingDeviceCompareOperator {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_compare_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 269730778) // FIXME: should cache?
  var ret RenderingDeviceCompareOperator
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetMinLod(p_member float32, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_lod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetMinLod() float32 {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_lod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetMaxLod(p_member float32, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_lod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetMaxLod() float32 {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_lod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetBorderColor(p_member RenderingDeviceSamplerBorderColor, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1115869595) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetBorderColor() RenderingDeviceSamplerBorderColor {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3514246478) // FIXME: should cache?
  var ret RenderingDeviceSamplerBorderColor
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDSamplerState) SetUnnormalizedUvw(p_member bool, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unnormalized_uvw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDSamplerState) GetUnnormalizedUvw() bool {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unnormalized_uvw")
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
