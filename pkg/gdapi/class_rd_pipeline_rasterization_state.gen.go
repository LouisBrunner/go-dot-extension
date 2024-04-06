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

type ptrsForRDPipelineRasterizationStateList struct {
	fnSetEnableDepthClamp        gdc.MethodBindPtr
	fnGetEnableDepthClamp        gdc.MethodBindPtr
	fnSetDiscardPrimitives       gdc.MethodBindPtr
	fnGetDiscardPrimitives       gdc.MethodBindPtr
	fnSetWireframe               gdc.MethodBindPtr
	fnGetWireframe               gdc.MethodBindPtr
	fnSetCullMode                gdc.MethodBindPtr
	fnGetCullMode                gdc.MethodBindPtr
	fnSetFrontFace               gdc.MethodBindPtr
	fnGetFrontFace               gdc.MethodBindPtr
	fnSetDepthBiasEnabled        gdc.MethodBindPtr
	fnGetDepthBiasEnabled        gdc.MethodBindPtr
	fnSetDepthBiasConstantFactor gdc.MethodBindPtr
	fnGetDepthBiasConstantFactor gdc.MethodBindPtr
	fnSetDepthBiasClamp          gdc.MethodBindPtr
	fnGetDepthBiasClamp          gdc.MethodBindPtr
	fnSetDepthBiasSlopeFactor    gdc.MethodBindPtr
	fnGetDepthBiasSlopeFactor    gdc.MethodBindPtr
	fnSetLineWidth               gdc.MethodBindPtr
	fnGetLineWidth               gdc.MethodBindPtr
	fnSetPatchControlPoints      gdc.MethodBindPtr
	fnGetPatchControlPoints      gdc.MethodBindPtr
}

var ptrsForRDPipelineRasterizationState ptrsForRDPipelineRasterizationStateList

func initRDPipelineRasterizationStatePtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDPipelineRasterizationState")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_enable_depth_clamp")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetEnableDepthClamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_depth_clamp")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetEnableDepthClamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_discard_primitives")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetDiscardPrimitives = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_discard_primitives")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetDiscardPrimitives = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_wireframe")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetWireframe = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_wireframe")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetWireframe = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_cull_mode")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetCullMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2662586502))
	}
	{
		methodName := StringNameFromStr("get_cull_mode")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetCullMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2192484313))
	}
	{
		methodName := StringNameFromStr("set_front_face")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetFrontFace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2637251213))
	}
	{
		methodName := StringNameFromStr("get_front_face")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetFrontFace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 708793786))
	}
	{
		methodName := StringNameFromStr("set_depth_bias_enabled")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetDepthBiasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_depth_bias_enabled")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetDepthBiasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_depth_bias_constant_factor")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetDepthBiasConstantFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth_bias_constant_factor")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetDepthBiasConstantFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_depth_bias_clamp")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetDepthBiasClamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth_bias_clamp")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetDepthBiasClamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_depth_bias_slope_factor")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetDepthBiasSlopeFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth_bias_slope_factor")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetDepthBiasSlopeFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_line_width")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetLineWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_line_width")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetLineWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_patch_control_points")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnSetPatchControlPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_patch_control_points")
		defer methodName.Destroy()
		ptrsForRDPipelineRasterizationState.fnGetPatchControlPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type RDPipelineRasterizationState struct {
	RefCounted
}

func (me *RDPipelineRasterizationState) BaseClass() string {
	return "RDPipelineRasterizationState"
}

func NewRDPipelineRasterizationState() *RDPipelineRasterizationState {
	str := StringNameFromStr("RDPipelineRasterizationState") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDPipelineRasterizationState{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDPipelineRasterizationState) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDPipelineRasterizationState) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDPipelineRasterizationState) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDPipelineRasterizationState) SetEnableDepthClamp(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetEnableDepthClamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetEnableDepthClamp() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetEnableDepthClamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineRasterizationState) SetDiscardPrimitives(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetDiscardPrimitives), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetDiscardPrimitives() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetDiscardPrimitives), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineRasterizationState) SetWireframe(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetWireframe), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetWireframe() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetWireframe), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineRasterizationState) SetCullMode(p_member RenderingDevicePolygonCullMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetCullMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetCullMode() RenderingDevicePolygonCullMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDevicePolygonCullMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetCullMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineRasterizationState) SetFrontFace(p_member RenderingDevicePolygonFrontFace) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetFrontFace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetFrontFace() RenderingDevicePolygonFrontFace {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDevicePolygonFrontFace

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetFrontFace), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineRasterizationState) SetDepthBiasEnabled(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetDepthBiasEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetDepthBiasEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetDepthBiasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineRasterizationState) SetDepthBiasConstantFactor(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetDepthBiasConstantFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetDepthBiasConstantFactor() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetDepthBiasConstantFactor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineRasterizationState) SetDepthBiasClamp(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetDepthBiasClamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetDepthBiasClamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetDepthBiasClamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineRasterizationState) SetDepthBiasSlopeFactor(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetDepthBiasSlopeFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetDepthBiasSlopeFactor() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetDepthBiasSlopeFactor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineRasterizationState) SetLineWidth(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetLineWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetLineWidth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetLineWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineRasterizationState) SetPatchControlPoints(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnSetPatchControlPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineRasterizationState) GetPatchControlPoints() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineRasterizationState.fnGetPatchControlPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
