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

type ptrsForRDSamplerStateList struct {
	fnSetMagFilter       gdc.MethodBindPtr
	fnGetMagFilter       gdc.MethodBindPtr
	fnSetMinFilter       gdc.MethodBindPtr
	fnGetMinFilter       gdc.MethodBindPtr
	fnSetMipFilter       gdc.MethodBindPtr
	fnGetMipFilter       gdc.MethodBindPtr
	fnSetRepeatU         gdc.MethodBindPtr
	fnGetRepeatU         gdc.MethodBindPtr
	fnSetRepeatV         gdc.MethodBindPtr
	fnGetRepeatV         gdc.MethodBindPtr
	fnSetRepeatW         gdc.MethodBindPtr
	fnGetRepeatW         gdc.MethodBindPtr
	fnSetLodBias         gdc.MethodBindPtr
	fnGetLodBias         gdc.MethodBindPtr
	fnSetUseAnisotropy   gdc.MethodBindPtr
	fnGetUseAnisotropy   gdc.MethodBindPtr
	fnSetAnisotropyMax   gdc.MethodBindPtr
	fnGetAnisotropyMax   gdc.MethodBindPtr
	fnSetEnableCompare   gdc.MethodBindPtr
	fnGetEnableCompare   gdc.MethodBindPtr
	fnSetCompareOp       gdc.MethodBindPtr
	fnGetCompareOp       gdc.MethodBindPtr
	fnSetMinLod          gdc.MethodBindPtr
	fnGetMinLod          gdc.MethodBindPtr
	fnSetMaxLod          gdc.MethodBindPtr
	fnGetMaxLod          gdc.MethodBindPtr
	fnSetBorderColor     gdc.MethodBindPtr
	fnGetBorderColor     gdc.MethodBindPtr
	fnSetUnnormalizedUvw gdc.MethodBindPtr
	fnGetUnnormalizedUvw gdc.MethodBindPtr
}

var ptrsForRDSamplerState ptrsForRDSamplerStateList

func initRDSamplerStatePtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDSamplerState")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_mag_filter")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetMagFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1493420382))
	}
	{
		methodName := StringNameFromStr("get_mag_filter")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetMagFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2209202801))
	}
	{
		methodName := StringNameFromStr("set_min_filter")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetMinFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1493420382))
	}
	{
		methodName := StringNameFromStr("get_min_filter")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetMinFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2209202801))
	}
	{
		methodName := StringNameFromStr("set_mip_filter")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetMipFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1493420382))
	}
	{
		methodName := StringNameFromStr("get_mip_filter")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetMipFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2209202801))
	}
	{
		methodName := StringNameFromStr("set_repeat_u")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetRepeatU = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 246127626))
	}
	{
		methodName := StringNameFromStr("get_repeat_u")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetRepeatU = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227895872))
	}
	{
		methodName := StringNameFromStr("set_repeat_v")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetRepeatV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 246127626))
	}
	{
		methodName := StringNameFromStr("get_repeat_v")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetRepeatV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227895872))
	}
	{
		methodName := StringNameFromStr("set_repeat_w")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetRepeatW = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 246127626))
	}
	{
		methodName := StringNameFromStr("get_repeat_w")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetRepeatW = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227895872))
	}
	{
		methodName := StringNameFromStr("set_lod_bias")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetLodBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_lod_bias")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetLodBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_use_anisotropy")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetUseAnisotropy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_use_anisotropy")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetUseAnisotropy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_anisotropy_max")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetAnisotropyMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_anisotropy_max")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetAnisotropyMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_enable_compare")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetEnableCompare = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_compare")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetEnableCompare = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_compare_op")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetCompareOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2573711505))
	}
	{
		methodName := StringNameFromStr("get_compare_op")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetCompareOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 269730778))
	}
	{
		methodName := StringNameFromStr("set_min_lod")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetMinLod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_min_lod")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetMinLod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_max_lod")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetMaxLod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_max_lod")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetMaxLod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_border_color")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetBorderColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1115869595))
	}
	{
		methodName := StringNameFromStr("get_border_color")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetBorderColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3514246478))
	}
	{
		methodName := StringNameFromStr("set_unnormalized_uvw")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnSetUnnormalizedUvw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_unnormalized_uvw")
		defer methodName.Destroy()
		ptrsForRDSamplerState.fnGetUnnormalizedUvw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

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

func (me *RDSamplerState) SetMagFilter(p_member RenderingDeviceSamplerFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetMagFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetMagFilter() RenderingDeviceSamplerFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceSamplerFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetMagFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDSamplerState) SetMinFilter(p_member RenderingDeviceSamplerFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetMinFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetMinFilter() RenderingDeviceSamplerFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceSamplerFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetMinFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDSamplerState) SetMipFilter(p_member RenderingDeviceSamplerFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetMipFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetMipFilter() RenderingDeviceSamplerFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceSamplerFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetMipFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDSamplerState) SetRepeatU(p_member RenderingDeviceSamplerRepeatMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetRepeatU), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetRepeatU() RenderingDeviceSamplerRepeatMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceSamplerRepeatMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetRepeatU), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDSamplerState) SetRepeatV(p_member RenderingDeviceSamplerRepeatMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetRepeatV), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetRepeatV() RenderingDeviceSamplerRepeatMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceSamplerRepeatMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetRepeatV), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDSamplerState) SetRepeatW(p_member RenderingDeviceSamplerRepeatMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetRepeatW), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetRepeatW() RenderingDeviceSamplerRepeatMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceSamplerRepeatMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetRepeatW), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDSamplerState) SetLodBias(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetLodBias), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetLodBias() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetLodBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDSamplerState) SetUseAnisotropy(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetUseAnisotropy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetUseAnisotropy() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetUseAnisotropy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDSamplerState) SetAnisotropyMax(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetAnisotropyMax), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetAnisotropyMax() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetAnisotropyMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDSamplerState) SetEnableCompare(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetEnableCompare), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetEnableCompare() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetEnableCompare), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDSamplerState) SetCompareOp(p_member RenderingDeviceCompareOperator) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetCompareOp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetCompareOp() RenderingDeviceCompareOperator {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceCompareOperator

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetCompareOp), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDSamplerState) SetMinLod(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetMinLod), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetMinLod() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetMinLod), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDSamplerState) SetMaxLod(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetMaxLod), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetMaxLod() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetMaxLod), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDSamplerState) SetBorderColor(p_member RenderingDeviceSamplerBorderColor) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetBorderColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetBorderColor() RenderingDeviceSamplerBorderColor {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceSamplerBorderColor

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetBorderColor), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDSamplerState) SetUnnormalizedUvw(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnSetUnnormalizedUvw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDSamplerState) GetUnnormalizedUvw() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDSamplerState.fnGetUnnormalizedUvw), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
