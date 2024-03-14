// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDPipelineRasterizationState struct {
  RefCounted
}

func (me *RDPipelineRasterizationState) BaseClass() string {
  return "RDPipelineRasterizationState"
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

func  (me *RDPipelineRasterizationState) SetEnableDepthClamp(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_depth_clamp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetEnableDepthClamp() bool {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_depth_clamp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetDiscardPrimitives(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_discard_primitives")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetDiscardPrimitives() bool {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_discard_primitives")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetWireframe(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wireframe")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetWireframe() bool {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wireframe")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetCullMode(p_member RenderingDevicePolygonCullMode, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2662586502) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetCullMode() RenderingDevicePolygonCullMode {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2192484313) // FIXME: should cache?
  var ret RenderingDevicePolygonCullMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetFrontFace(p_member RenderingDevicePolygonFrontFace, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_front_face")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2637251213) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetFrontFace() RenderingDevicePolygonFrontFace {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_front_face")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 708793786) // FIXME: should cache?
  var ret RenderingDevicePolygonFrontFace
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetDepthBiasEnabled(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_bias_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetDepthBiasEnabled() bool {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_bias_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetDepthBiasConstantFactor(p_member float32, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_bias_constant_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetDepthBiasConstantFactor() float32 {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_bias_constant_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetDepthBiasClamp(p_member float32, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_bias_clamp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetDepthBiasClamp() float32 {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_bias_clamp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetDepthBiasSlopeFactor(p_member float32, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_bias_slope_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetDepthBiasSlopeFactor() float32 {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_bias_slope_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetLineWidth(p_member float32, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetLineWidth() float32 {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineRasterizationState) SetPatchControlPoints(p_member int, )  {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_patch_control_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineRasterizationState) GetPatchControlPoints() int {
  classNameV := StringNameFromStr("RDPipelineRasterizationState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_patch_control_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
