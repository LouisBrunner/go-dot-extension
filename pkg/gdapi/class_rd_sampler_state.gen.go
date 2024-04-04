// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type RDSamplerState struct {
  RefCounted
}

func (me *RDSamplerState) BaseClass() string {
  return "RDSamplerState"
}

func NewRDSamplerState() *RDSamplerState {
  str := StringNameFromStr("RDSamplerState") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDSamplerState{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetMagFilter() RenderingDeviceSamplerFilter {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mag_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2209202801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceSamplerFilter

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDSamplerState) SetMinFilter(p_member RenderingDeviceSamplerFilter, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1493420382) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetMinFilter() RenderingDeviceSamplerFilter {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2209202801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceSamplerFilter

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDSamplerState) SetMipFilter(p_member RenderingDeviceSamplerFilter, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mip_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1493420382) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetMipFilter() RenderingDeviceSamplerFilter {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mip_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2209202801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceSamplerFilter

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDSamplerState) SetRepeatU(p_member RenderingDeviceSamplerRepeatMode, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_repeat_u")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 246127626) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetRepeatU() RenderingDeviceSamplerRepeatMode {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_repeat_u")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227895872) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceSamplerRepeatMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDSamplerState) SetRepeatV(p_member RenderingDeviceSamplerRepeatMode, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_repeat_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 246127626) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetRepeatV() RenderingDeviceSamplerRepeatMode {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_repeat_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227895872) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceSamplerRepeatMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDSamplerState) SetRepeatW(p_member RenderingDeviceSamplerRepeatMode, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_repeat_w")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 246127626) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetRepeatW() RenderingDeviceSamplerRepeatMode {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_repeat_w")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227895872) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceSamplerRepeatMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDSamplerState) SetLodBias(p_member float64, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lod_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetLodBias() float64 {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lod_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDSamplerState) SetUseAnisotropy(p_member bool, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetUseAnisotropy() bool {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDSamplerState) SetAnisotropyMax(p_member float64, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anisotropy_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetAnisotropyMax() float64 {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_anisotropy_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDSamplerState) SetEnableCompare(p_member bool, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_compare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetEnableCompare() bool {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_compare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDSamplerState) SetCompareOp(p_member RenderingDeviceCompareOperator, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_compare_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2573711505) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetCompareOp() RenderingDeviceCompareOperator {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_compare_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 269730778) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceCompareOperator

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDSamplerState) SetMinLod(p_member float64, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_lod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetMinLod() float64 {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_lod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDSamplerState) SetMaxLod(p_member float64, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_lod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetMaxLod() float64 {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_lod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDSamplerState) SetBorderColor(p_member RenderingDeviceSamplerBorderColor, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1115869595) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetBorderColor() RenderingDeviceSamplerBorderColor {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3514246478) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceSamplerBorderColor

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDSamplerState) SetUnnormalizedUvw(p_member bool, )  {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unnormalized_uvw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDSamplerState) GetUnnormalizedUvw() bool {
  classNameV := StringNameFromStr("RDSamplerState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unnormalized_uvw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
